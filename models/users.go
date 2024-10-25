package models

type User struct {
	USER_ID       int64    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token string `json:"token"`
}