package service

import (
	"SimpleBBS-server/dao/mysql"
	"SimpleBBS-server/models"
	"SimpleBBS-server/utils/jwt"
	"SimpleBBS-server/utils/snowflake"
)

func SignUp(p *models.SignUpParam) (err error) {
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	userid := snowflake.GenID()
	return mysql.InsertUser(models.User{
		UserID:   userid,
		Username: p.Username,
		Password: p.Password,
	})
}

func Login(p *models.LoginParam) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
