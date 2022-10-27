package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func RelationRoutes(r *gin.RouterGroup) {
	relation := r.Group("relation")
	{
		relation.POST("/action/", common.AuthMiddleware(), controller.RelationAction)
		relation.GET("/follow/list/", common.AuthWithOutMiddleware(), controller.GetFollowList)
		relation.GET("/follower/list/", common.AuthWithOutMiddleware(), controller.GetFollowerList)
	}
}
