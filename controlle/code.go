package controlle

// 定义返回状态码类型
type ResCode int64

// 定义返回状态码及其对应消息
const (
	CodeSuccess ResCode=iota + 1000
	CodeInvalidParam
	CodeUserExists
	CodeuserNotExists
	CodePasswordNotMatch
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

)


var Msgmap = map[ResCode]string {
	CodeSuccess: "success",
    CodeInvalidParam: "invalid params",
    CodeUserExists: "user already exists",
    CodeuserNotExists: "user not exists",
    CodePasswordNotMatch: "password not match",
    CodeServerBusy: "server is busy",

	CodeNeedLogin: "need login",
    CodeInvalidToken: "invalid token",
}
func (this ResCode) Msg() string {
	return Msgmap[this]
}
