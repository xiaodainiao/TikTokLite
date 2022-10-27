package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.POST("/action/", common.AuthMiddleware(), controller.CommentAction)
		comment.GET("/list/", common.AuthWithOutMiddleware(), controller.GetCommentList)
	}

}
