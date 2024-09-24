package modlels

type ParamRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required,eqfield=Password"`

}