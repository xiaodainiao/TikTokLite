package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(r *gin.RouterGroup) {
	favorite := r.Group("favorite")
	{
		favorite.POST("/action/", common.AuthMiddleware(), controller.FavoriteAction)
		favorite.GET("/list/", common.AuthWithOutMiddleware(), controller.GetFavoriteList)
	}

}
