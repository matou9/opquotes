package router

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	`github.com/gin-gonic/gin`
	`net/http`
	"opquotes/config"
	"strconv"
	"time"
)

/**
  JWTAuth 中间件，检查token
*/
var (
	TokenExpired     = errors.New("token已经过期")
	TokenNotValidYet = errors.New("token没有激活")
	TokenMalformed   = errors.New("这不是一个token")
	TokenInvalid     = errors.New("token非法")
	SignKey          = config.SecretKey
)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func AuthorizeAgent() gin.HandlerFunc {
	return func(context *gin.Context){
		var (
			token string
			claims *CustomClaims
			err    error
		)
		token =context.Request.Header.Get("token")
		if token ==""{
			context.JSON(http.StatusOK,config.ErrUserNoToken)
			context.Abort()
		} else{
			if claims, err = NewJWT("adminsign",jwt.SigningMethodHS512).ParseToken(token); err != nil {
				if err == TokenExpired {
					context.JSON(http.StatusOK, config.ErrUserTokenExpired)
					context.Abort()
					return
				} else if err == TokenNotValidYet {
					context.JSON(http.StatusOK, config.ErrUserTokenNotValidYet)
					context.Abort()
					return
				} else if err == TokenMalformed {
					context.JSON(http.StatusOK, config.ErrUserTokenMalformed)
					context.Abort()
					return
				} else if err == TokenInvalid {
					context.JSON(http.StatusOK, config.ErrUserTokenInvalid)
					context.Abort()
					return
				}
				context.JSON(http.StatusOK, config.ErrBaseUnknown)
				context.Abort()
				return
			}
			if claims.Uid==19790205{   //claims.Code==account || (len(account)==11 && claims.Mobile==account)||
				context.Set("claims", claims)
				context.Next()

			}else{
				context.JSON(http.StatusOK, config.ErrUserTokenInvalid)
				context.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去

		}
	}

}


func Authorize() gin.HandlerFunc {
	return func(context *gin.Context){
		var (
			token string
			//account string
			claims *CustomClaims
			err    error
		)
		_uid :=context.PostForm("uid")//支持uid，code，mobile验证
		uid,_:=strconv.ParseInt(_uid,10,64)
		token =context.Request.Header.Get("token")
		if token ==""{
			context.JSON(http.StatusOK,config.ErrUserNoToken)
			context.Abort()
		} else{
			if claims, err = NewJWT("fengxiaoyong7925",jwt.SigningMethodHS256).ParseToken(token); err != nil {
				if err == TokenExpired {
					context.JSON(http.StatusOK, config.ErrUserTokenExpired)
					context.Abort()
					return
				} else if err == TokenNotValidYet {
					context.JSON(http.StatusOK, config.ErrUserTokenNotValidYet)
					context.Abort()
					return
				} else if err == TokenMalformed {
					context.JSON(http.StatusOK, config.ErrUserTokenMalformed)
					context.Abort()
					return
				} else if err == TokenInvalid {
					context.JSON(http.StatusOK, config.ErrUserTokenInvalid)
					context.Abort()
					return
				}
				context.JSON(http.StatusOK, config.ErrBaseUnknown)
				context.Abort()
				return
			}
			if claims.Uid==uid{   //claims.Code==account || (len(account)==11 && claims.Mobile==account)||
				context.Set("claims", claims)
				context.Next()

			}else{
				context.JSON(http.StatusOK, config.ErrUserTokenInvalid)
				context.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去

		}
	}

}


/**
  JWT 签名结构
*/
type CustomClaims struct {
	Uid   int64 `json:"uid"`
	Code   string `json:"code"`
	Name  string `json:"name"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
	Encryption  *jwt.SigningMethodHMAC
}
func NewJWT(SignKey string,Encryption *jwt.SigningMethodHMAC) *JWT {
	return &JWT{
		[]byte(SignKey),
		Encryption,
	}
}
func GetSignKey() string {
	return SignKey
}

/**
设置SignKey
*/
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(j.Encryption, claims)
	return token.SignedString(j.SigningKey)
}
/**
  更新token
*/
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
/**
  解析Token
*/
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	if token, err = jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	}); err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
/**
  一些常量
*/


