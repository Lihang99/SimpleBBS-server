package mysql

import (
	"SimpleBBS-server/models"
	"SimpleBBS-server/utils"
	"database/sql"
)

func InsertUser(user models.User) (err error) {
	user.Password = utils.EncryptPassword(user.Password)
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return UserExistError
	}
	return
}

func Login(user *models.User) (err error) {
	originPwd := user.Password
	sqlStr := `select user_id,password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return UserNotExistError
	}
	password := utils.EncryptPassword(originPwd)
	if password != user.Password {
		return PasswordInvalidError
	}
	return
}
