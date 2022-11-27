package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-redis/utils"
	"net/http"
	"strings"
	"time"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("jwt"),
	}
}

type MyClaims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := utils.InitResBody()
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			res.Message = "未获取到token"
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			res.Message = "token格式不正确"
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			res.Message = "token错误"
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				res.Code = http.StatusInternalServerError
				res.Message = "token授权已过期,请重新登录"
				c.JSON(http.StatusOK, res)
				c.Abort()
				return
			}
			// 其他错误
			res.Code = http.StatusInternalServerError
			res.Message = err.Error()
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}
		c.Set("phone", claims.Phone)
		c.Next()
	}
}

// 生成token
func SetToken(phone string) (string,error) {
	j := NewJWT()
	token, err := j.CreateToken(MyClaims{
		phone,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "go-redis",
		},
	})
	return token, err
}


