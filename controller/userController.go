package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	userName := ctx.Query("username")
	password := ctx.Query("password")
	if len(userName) > 32 || len(password) > 32 { //最长32位字符
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	loginResponse, err := service.UserLogin(userName, password)
	if err != nil {
		log.Infof("login error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", loginResponse)
}

func UserRegister(ctx *gin.Context) {
	var err error
	userName := ctx.Query("username")
	password := ctx.Query("password")
	if len(userName) > 32 || len(password) > 32 { //最长32位字符
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	registResponse, err := service.UserRegister(userName, password)
	if err != nil {
		log.Infof("registe error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", registResponse)

}

//获取用户信息
func GetUserInfo(ctx *gin.Context) {

	var err error
	userId := ctx.Query("user_id")
	uids, _ := ctx.Get("UserId")

	uid := uids.(int64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	if strconv.FormatInt(uid, 10) != userId {
		response.Fail(ctx, "token error", nil)
		return
	}
	userinfo, err := service.UserInfo(uid)
	if err != nil {
		log.Infof("get userinfo  error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", userinfo)

}
