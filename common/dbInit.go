//package repository
package common

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DataBase *gorm.DB

func InitDatabase() {
	var err error
	conf := config.GetConfig()
	host := conf.Mysql.Host
	port := conf.Mysql.Port
	database := conf.Mysql.Database
	username := conf.Mysql.Username
	password := conf.Mysql.Password
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	log.Info(args)
	DataBase, err = gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	log.Infof("connect database success,user:%s,database:%s", username, database)
}

func GetDB() *gorm.DB {
	return DataBase
}
func CloseDataBase() {
	DataBase.Close()
}
