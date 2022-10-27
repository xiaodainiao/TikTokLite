package response

import (
	"TikTokLite/log"
	"github.com/gin-gonic/gin"
	"reflect"
)

const (
	successCode = 0
	errorCode   = 1
)

type response struct {
	StatusCode int32
	StatusMsg  string
}

func Response(ctx *gin.Context, httpStatus int, v interface{}) {
	ctx.JSON(httpStatus, v)
}

func Success(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 200, response{successCode, msg})
	} else {
		setResponse(ctx, successCode, msg, v)
		Response(ctx, 200, v)
	}
}

func Fail(ctx *gin.Context, msg string, v interface{}) {
	if v == nil {
		Response(ctx, 200, response{errorCode, msg})
	} else {
		setResponse(ctx, errorCode, msg, v)
		Response(ctx, 200, v)
		ctx.Abort()
	}

}

func setResponse(ctx *gin.Context, StatusCode int64, StatusMsg string, v interface{}) {
	getValue := reflect.ValueOf(v)
	field := getValue.Elem().FieldByName("StatusMsg")
	if field.CanSet() {
		field.SetString(StatusMsg)
	} else {
		log.Debug("cant set msg")
	}
	fieldCode := getValue.Elem().FieldByName("StatusCode")
	if fieldCode.CanSet() {
		fieldCode.SetInt(StatusCode)
	} else {
		log.Debug("cant set StatusCode")
	}
}
