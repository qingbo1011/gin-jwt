package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	// 声明签名信息
	Secret []byte
}

// NewJWT 初始化JWT对象
func NewJWT() *JWT {
	return &JWT{
		Secret: []byte("qingbo1011.top"),
	}
}

// CustomClaims 自定义有效载荷
type CustomClaims struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	jwt.StandardClaims        // StandardClaims结构体实现了Claims接口(Valid()函数)
}

// CreateToken 调用jwt-go库生成token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// 指定编码算法为jwt.SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 返回一个token结构体指针
	return token.SignedString(j.Secret)
}

// ParserToken token解码
func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	//Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	//func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok { // jwt.ValidationError：是一个无效token的错误结构
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { // ValidationErrorMalformed是一个uint常量，表示token不可用
				return nil, fmt.Errorf("token不可用")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { // ValidationErrorExpired表示Token过期
				return nil, fmt.Errorf("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { // ValidationErrorNotValidYet表示无效token
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}
		}
		return nil, err
	}
	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token不可用")
}

// JWTAuth 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权访问！",
				"data":   nil,
			})
			c.Abort()	// Abort 函数在被调用的函数中阻止后续中间件的执行(这里都没有携带token，后续就不用执行了)
			return
		}

		log.Println("fet token：" + token)

		j := NewJWT()                       // 初始化一个JWT对象实例，并根据结构体方法来解析token
		claims, err := j.ParserToken(token) // 解析token中包含的相关信息（有效载荷）
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中（gin的上下文）
		c.Set("claims",claims)
	}
}
