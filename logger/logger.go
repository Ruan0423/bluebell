package logger

import (

	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
	"web_framework/settings"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init 初始化Logger
func Init(cfg *settings.APPConfig) (err error) {

	encoder := getEncoder()

	// //1.把日志只记录到一个文件里面，日志level来源于config配置文件。
	// var l = new(zapcore.Level)
	// err = l.UnmarshalText([]byte(cfg.Level))
	// if err != nil {
	// 	return
	// }
	// writeSyncer := getLogWriter(
	// 	cfg.Filename,
	// 	cfg.MaxSize,
	// 	cfg.MaxBackups,
	// 	cfg.MaxAge,
	// )
	// core1 := zapcore.NewCore(encoder, writeSyncer, l)
	// lg := zap.New(core1, zap.AddCaller())   
	// // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	// zap.ReplaceGlobals(lg)


	//2.分等级创建日志，比如info 类和 error类的日志分别存在log_info.log 和log_err.log中
	core_debug := zapcore.NewCore(encoder, getLogWriter(cfg.Model,"log_debug.log",cfg.MaxSize,cfg.MaxBackups,cfg.MaxAge),zap.DebugLevel)
	core_info := zapcore.NewCore(encoder,getLogWriter(cfg.Model,"log_info.log",cfg.MaxSize,cfg.MaxBackups,cfg.MaxAge),zap.InfoLevel)
	core_err := zapcore.NewCore(encoder, getLogWriter(cfg.Model,"log_err.log",cfg.MaxSize,cfg.MaxBackups,cfg.MaxAge), zap.ErrorLevel)
	core := zapcore.NewTee(core_debug,core_info,core_err)
	log_ingo_err := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(log_ingo_err)
	

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(model string,filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	LogWriter :=zapcore.AddSync(lumberJackLogger) //日志记录
	if model == "dev" {
		consoleWriter := zapcore.AddSync(os.Stdout) // 打印在控制台
		return zapcore.NewMultiWriteSyncer(consoleWriter, LogWriter)
	}

	return LogWriter
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
