package models

type SignUpParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
