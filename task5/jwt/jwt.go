package JWT

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var signkey string = "Default key"

type Jwt struct {
	SignKey []byte
}

type MyClaims struct {
	jwt.StandardClaims
}

func SetKey(key string) {
	signkey = key
}

func GetKey() string {
	return signkey
}

func NewJwt() *Jwt {
	return &Jwt{
		[]byte(GetKey()),
	}
}

func (j *Jwt) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

func (j *Jwt) ParseToken(token string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		if er, ok := err.(*jwt.ValidationError); ok {
			if er.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("错误token结构")
				// ValidationErrorExpired表示Token过期
			} else if er.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("过期token")
				// ValidationErrorNotValidYet表示无效token
			} else if er.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("无效token")
			} else {
				return nil, errors.New("无法辨认的token")
			}

		}
	}
	if claim, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claim, nil
	}
	return nil, errors.New("Invalid Token")
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "无token",
			})
			c.Abort()
			return
		}
		j := NewJwt()
		claims, err := j.ParseToken(token)
		//fmt.Println(claims, err)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func GenerateToken(c *gin.Context) {
	j := NewJwt()
	claims := MyClaims{
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1),
			ExpiresAt: int64(time.Now().Unix() + 3600*24),
			Issuer:    "Threes",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
	return
}
