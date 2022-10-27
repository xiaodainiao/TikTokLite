package main

import (
	"TikTokLite/common"
	"TikTokLite/config"
	"TikTokLite/log"
	"TikTokLite/minioStore"
	"TikTokLite/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	defer common.CloseDataBase()
	defer common.CloseRedis()
	defer log.Sync()

	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
}

func Init() {
	config.LoadConfig()
	log.InitLog()
	common.InitDatabase()
	minioStore.InitMinio()
	common.RedisInit()
}
