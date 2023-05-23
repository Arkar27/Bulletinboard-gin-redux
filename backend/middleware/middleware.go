package middleware

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	constants "github.com/Arkar27/gin-bulletinboard/backend/consts"
	"github.com/Arkar27/gin-bulletinboard/backend/dao/userDao"
	"github.com/Arkar27/gin-bulletinboard/backend/helper"
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			helper.ErrorPanic(constants.TokenNotProvided, ctx)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("INVALID SIGNNIG")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			helper.ErrorPanic(err, ctx)
			return
		}

		// check user role and allow only admin to delete user
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract the user email and name from the claims
			userEmail := claims["email"].(string)
			userName := claims["name"].(string)

			if userEmail != "" && userName != "" {

				userDao := userDao.NewUserDao(initializers.DB)
				user := userDao.FindUser(userName, userEmail, ctx)
				method := ctx.Request.Method
				isSensRoute := isSensRoute(ctx)

				switch method {
				case "DELETE":
					// admin can do delete
					if user.Type == "1" && isSensRoute {
						helper.ErrorPanic(constants.NoPermission, ctx)
						return
					}

				case "PUT":
					userId := ctx.Param("id")

					// member type is not allowed to update other user info
					if user.Type == "1" &&
						isSensRoute &&
						strconv.Itoa(int(user.ID)) != userId {

						helper.ErrorPanic(constants.NoPermission, ctx)
						return
					}
				}

			} else {
				helper.ErrorPanic(constants.NotValidToken, ctx)
				return
			}

		} else {
			helper.ErrorPanic(constants.NotValidToken, ctx)
			return
		}

		if !token.Valid {
			helper.ErrorPanic(err, ctx)
			return
		}

		// Token is valid, continue to the next middleware or handler
		ctx.Next()
	}
}

func isSensRoute(ctx *gin.Context) bool {
	route := strings.TrimSuffix(ctx.FullPath(), "/")
	pattern := "/api/users/:id"

	return route == pattern
}
