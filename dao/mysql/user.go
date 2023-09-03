package mysql

import (
	"basic_service/kitex_gen/basic"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

const serret = "clmmwjzb"

//func RegisterUser(username string) error {
//	var existingUser *models.User
//	err := DB.Table("users").Where("username = ?", username).First(&existingUser).Error
//	if existingUser != nil {
//		return fmt.Errorf("Username already exists")
//	}
//	if err != nil {
//		log.Printf("search user Isexist failed: %v", err)
//		return err
//	}
//	return nil
//}

func CheckUserExist(username string) (err error) {
	var count int64
	err = DB.Table("users").Select("count(id)").Where("name = ?", &username).Scan(&count).Error
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *basic.User) error {
	password := encryptPassword(user.Password)
	user.Password = password
	return DB.Create(user).Error
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(serret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(p *basic.DouyinUserLoginRequest) (int64, error) {

	var tmp basic.User
	if err := DB.Select("id,name,password").Where("name = ?", p.GetUsername()).First(&tmp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("user not found")
			return 0, ErrorUserNotExist
		}
		log.Printf("login failed：%v", err)
		return 0, err
	}
	// 判断密码是否正确
	password := encryptPassword(p.GetPassword())
	if password != tmp.Password {
		log.Printf("password is wrong")
		return 0, ErrorInvalidPassword
	}
	return tmp.GetId(), nil
}

func GetUserInfo(userId int64) (*basic.User, error) {

	var tmp basic.User

	if err := DB.Omit("password").First(&tmp, "id = ?", userId).Error; err != nil {
		fmt.Printf("get user failed: %v", err)
		return nil, err
	}

	return &tmp, nil
}
