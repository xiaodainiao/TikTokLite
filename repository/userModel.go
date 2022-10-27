package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"encoding/json"
	"strconv"

	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// gorm.Model
	Id              int64  `gorm:"column:user_id; primary_key;"`
	Name            string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	Follow          int64  `gorm:"column:follow_count"`
	Follower        int64  `gorm:"column:follower_count"`
	Avatar          string `gorm:"column:avatar"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
	TotalFav        int64  `gorm:"column:total_favorited"`
	FavCount        int64  `gorm:"column:favorite_count"`
}

func (User) TableName() string {
	return "users"
}

//检查该用户名是否已经存在
func UserNameIsExist(userName string) error {
	db := common.GetDB()
	user := User{}
	err := db.Where("user_name = ?", userName).Find(&user).Error
	if err == nil {
		return errors.New("username exist")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

//创建用户
func InsertUser(userName, password string) (*User, error) {
	db := common.GetDB()
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		Name:            userName,
		Password:        string(hasedPassword),
		Follow:          0,
		Follower:        0,
		TotalFav:        0,
		FavCount:        0,
		Avatar:          "https://tse1-mm.cn.bing.net/th/id/R-C.d83ded12079fa9e407e9928b8f300802?rik=Gzu6EnSylX9f1Q&riu=http%3a%2f%2fwww.webcarpenter.com%2fpictures%2fGo-gopher-programming-language.jpg&ehk=giVQvdvQiENrabreHFM8x%2fyOU70l%2fy6FOa6RS3viJ24%3d&risl=&pid=ImgRaw&r=0",
		BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpvO33-wQEPgHaEq?pid=ImgDet&rs=1",
		Signature:       "test sign",
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("regist user:%+v", user)
	go CacheSetUser(user)
	return &user, nil
}

//获取用户信息
func GetUserInfo(u interface{}) (User, error) {
	db := common.GetDB()
	user := User{}
	var err error
	switch u := u.(type) {
	case int64:
		user, err = CacheGetUser(u)
		if err == nil {
			return user, nil
		}
		err = db.Where("user_id = ?", u).Find(&user).Error

	case string:
		err = db.Where("user_name = ?", u).Find(&user).Error
	default:
		err = errors.New("")
	}
	if err != nil {
		return user, errors.New("user error")
	}
	go CacheSetUser(user)
	log.Infof("%+v", user)
	return user, nil
}

func CacheSetUser(u User) {
	uid := strconv.FormatInt(u.Id, 10)
	err := common.CacheSet("user_"+uid, u)
	if err != nil {
		log.Errorf("set cache error:%+v", err)
	}
}

func CacheGetUser(uid int64) (User, error) {
	key := strconv.FormatInt(uid, 10)
	data, err := common.CacheGet("user_" + key)
	user := User{}
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
