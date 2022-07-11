package api

import (
	"gin-JWT/middleware"
	"gin-JWT/model"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login 定义登录逻辑
func Login(c *gin.Context) {
	var loginReq model.LoginReq
	err := c.ShouldBind(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		return
	}
	isPass, user, err := model.LoginCheck(loginReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	if isPass {
		generateToken(c, user) // 登录成功，签发token
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "登录失败,用户名或密码错误！",
		})
	}

}

// token生成器
func generateToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()           // 构造SignKey: 签名和解签名需要使用一个值
	claims := middleware.CustomClaims{ // 构造用户claims信息（负载）
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "qingbo1011.top",                // 签名颁发者
		},
	}
	token, err := j.CreateToken(claims)	// 根据claims生成token对象
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":"生成token失败",
			"error":err.Error(),
		})
		return
	}

	log.Println(token)

	data := model.LoginResponse{	// 封装一个响应数据，返回用户名和token
		Name:  user.Name,
		Token: token,
	}
	c.JSON(http.StatusOK,gin.H{
		"status":0,
		"msg":"登录成功！",
		"data":data,
	})
}
