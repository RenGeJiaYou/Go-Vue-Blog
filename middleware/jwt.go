package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-vue-blog/utils"
	"go-vue-blog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

//todo 细化错误分类，如超时错误，无 token 错误，验证失败错误

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, error) {
	expiresTime := time.Now().Add(24 * time.Hour)
	claims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			Issuer:    "govueblog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// ParseToken 验证token
func ParseToken(token string) (*MyClaims, int) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
		return claims, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR_TOKEN_WRONG
	}
}

// JwtToken 中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//1.若不存在token
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			//中断数据向内层流动
			c.Abort()
			return
		}

		//2.若存在token,但不是正确的格式
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			//中断数据向内层流动
			c.Abort()
			return
		}

		//3.若格式正确，但内容没通过验证
		claims, tCode := ParseToken(checkToken[1])
		if tCode == errmsg.ERROR_TOKEN_WRONG {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		//4.若token 内容无误，但过期了
		if time.Now().Unix() > claims.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}

}
