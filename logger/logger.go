package logger

import (
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"

	// "gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//使用自定义的zap logger
func Init() (err error){
	
	//编码器
	encoder := getEncoder()
	//输出位置
	writeSyncer:= getwriteSyncer(viper.GetString("log.filename"))

	//定义core
	core1 := zapcore.NewCore(
		encoder,
		writeSyncer,
		zapcore.DebugLevel,//全记录
	)
	//错误日志
	core2 :=  zapcore.NewCore(
		encoder,
		getwriteSyncer("log.err"),
		zapcore.ErrorLevel,
	)

	//创建单个logger
	// logger:= zap.New(core1,zap.AddCaller(), zap.AddCallerSkip(1)) //AddCaller详细记录调用的代码行，AddCallerSkip(1)调用链很多时直接跳过
	// return logger.Sugar()

	//创建双日志，全日志和错误日志
	c:=zapcore.NewTee(core1,core2)
	logger:= zap.New(c,zap.AddCaller())
	//替换zap的全局logger
	zap.ReplaceGlobals(logger)
	return
}

//设置日志编译器，什么类型的日志
func getEncoder() zapcore.Encoder{
	//encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//设置时间格式为2024-9-1-12.32
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//json格式
	// jsonencoder := zapcore.NewJSONEncoder(encoderConfig)

	//终端形式
	ConsoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	return ConsoleEncoder

}

//设置输出位置
func getwriteSyncer(logfilename string) zapcore.WriteSyncer {
	//日志文件
	// logfile, _ :=os.OpenFile(logfilename,os.O_APPEND | os.O_CREATE|os.O_RDWR,0666)

	//分割日志
	
	l, _ := rotatelogs.New(
		logfilename+".%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)

	//也输出到终端
	wc := io.MultiWriter(l,os.Stdout)
	return zapcore.AddSync(wc)
}


//Ginlogger
func GinLogger() gin.HandlerFunc {

	return func (c *gin.Context)  {

		start := time.Now()
		path  := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()
		cost :=time.Since(start)
		zap.L().Info(
			path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost), // 运行时间
		)
		
	}

}
// GinRecovery
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