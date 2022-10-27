package common

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"errors"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

var (
	Secret = []byte("TikTok")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

//生成token
func GenToken(userid int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userid,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "server",
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),可用于设定token过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("TikTok"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//解析token
func ParsenToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//验证token
func VerifyToken(tokenString string) (int64, error) {

	log.Debugf("tokenString:%v", tokenString)

	if tokenString == "" {
		return int64(0), nil
	}
	claims, err := ParsenToken(tokenString)
	if err != nil {
		return int64(0), err
	}
	return claims.UserId, nil
}

//=============================gin的中间件，就是一个函数，返回gin 的HandlerFunc======================================================

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.PostForm("token")
		if tokenString == "" {
			tokenString = c.Query("token")
		}

		userId, err := VerifyToken(tokenString)
		if err != nil || userId == int64(0) {
			response.Fail(c, "auth error", nil)
			c.Abort()
		}

		c.Set("UserId", userId)
		c.Next()
	}
}

//部分接口不需要用户登录也可访问，如feed，pushlishlist，favList，follow/follower list
func AuthWithOutMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.Query("token")

		userId, err := VerifyToken(tokenString)
		if err != nil {
			response.Fail(c, "auth error", nil)
			c.Abort()
		}

		c.Set("UserId", userId)
		c.Next()
	}
}
