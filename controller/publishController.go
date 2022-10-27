package controller

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

//视频发布
func PublishAction(ctx *gin.Context) {
	// publishResponse := &message.DouyinPublishActionResponse{}
	userId, _ := ctx.Get("UserId")
	//token := ctx.PostForm("token")
	//userId, err := common.VerifyToken(token)
	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%s_%s", util.RandomString(), filename)
	videoPath := config.GetConfig().Path.Videofile
	saveFile := filepath.Join(videoPath, finalName)

	log.Info("saveFile:", saveFile)

	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	publish, err := service.PublishVideo(userId.(int64), saveFile, title)
	//publish, err := service.PublishVideo(userId, saveFile)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Infof("publish:%v", publish)
	response.Success(ctx, "success", publish)

}

//获取视频列表
func GetPublishList(ctx *gin.Context) {
	tokenUserId, _ := ctx.Get("UserId")
	id := ctx.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	list, err := service.PublishList(tokenUserId.(int64), userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", list)
}
