package mysql

import "errors"

var (
	UserExistError       = errors.New("用户已经存在")
	UserNotExistError    = errors.New("用户不存在")
	PasswordInvalidError = errors.New("用户名或密码错误")
)
