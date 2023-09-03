package logic

import (
	"basic_service/dao/mysql"
	"basic_service/kitex_gen/basic"
	"basic_service/pkg/jwt"
	"basic_service/pkg/snowflake"
	"log"
)

func SignUp(p *basic.DouyinUserRegisterRequest) (string, int64, error) {
	// 1.判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return "", 0, err
	}

	// 2.生成UID
	userID := snowflake.GenID()
	// 构建一个User实例
	u := &basic.User{
		Id:       userID,
		Name:     p.Username,
		Password: p.Password,
	}
	// 生成JWT
	//token, err := jwt.GenToken(userID, p.GetUsername())
	//if err != nil {
	//	log.Printf("gen token failed: %v\n", err)
	//	return "", 0, err
	//}

	// 3.保存进数据库
	// redis.xxx ...
	if err := mysql.InsertUser(u); err != nil {
		return "", 0, err
	}

	return "", userID, nil
}

func Login(p *basic.DouyinUserLoginRequest) (int64, string, error) {

	// 1.判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != mysql.ErrorUserExist && err != nil {
		return 0, "", err
	}
	userId, err := mysql.Login(p)
	if err != nil {
		return 0, "", err
	}

	// 生成JWT
	token, err := jwt.GenToken(userId, p.GetUsername())
	if err != nil {
		log.Printf("gen token failed: %v\n", err)
		return 0, "", err
	}

	return userId, token, nil
}

func GetUserInfo(userId int64) (*basic.User, error) {

	return mysql.GetUserInfo(userId)
}
