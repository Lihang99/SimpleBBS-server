package controller

import (
	"SimpleBBS-server/dao/mysql"
	"SimpleBBS-server/models"
	"SimpleBBS-server/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.SignUpParam)
	//if checkRequest(c,p){
	//	return
	//}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	if err := service.SignUp(p); err != nil {
		zap.L().Error("sign up failed", zap.Error(err))
		if errors.Is(err, mysql.UserExistError) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	user, err := service.Login(p)
	if err != nil {
		zap.L().Error("login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.UserNotExistError) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID),
		"username": user.Username,
		"token":    user.Token,
	})
}
