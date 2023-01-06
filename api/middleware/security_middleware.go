package middleware

import (
	"github.com/Unbel1evab7e/bank-interface/db/entity"
	"github.com/Unbel1evab7e/bank-interface/domain/dto"
	"github.com/Unbel1evab7e/bank-interface/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func NewSecurityMiddleware(personService *service.PersonService) *jwt.GinJWTMiddleware {
	identityKey := "phone"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.Person); ok {
				return jwt.MapClaims{
					identityKey: v.Phone,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &entity.Person{
				Phone: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals dto.LoginDto
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			phone := loginVals.Phone
			password := loginVals.Password

			person, err := personService.GetPersonByPhoneAndPassword(phone, password)

			if err != nil {
				return nil, err
			}

			return person, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*entity.Person); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		SendCookie:  true,
		TokenLookup: "cookie:Authorization",
		CookieName:  "Authorization",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		logrus.Fatal("Failed to create security middleware")
	}

	return authMiddleware
}
