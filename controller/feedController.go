package controller

import (
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"strconv"

	// "encoding/json"
	"github.com/gin-gonic/gin"
)

//视频流
func Feed(ctx *gin.Context) {
	var userId int64
	currentTime, err := strconv.ParseInt(ctx.Query("latest_time"), 10, 64)
	if err != nil || currentTime == int64(0) {
		currentTime = util.GetCurrentTime()
	}
	//token := ctx.Query("token")
	//userId, err = common.VerifyToken(token)
	userIds, _ := ctx.Get("UserId")
	userId = userIds.(int64)

	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	feedList, err := service.GetFeedList(currentTime, userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	response.Success(ctx, "success", feedList)
}
