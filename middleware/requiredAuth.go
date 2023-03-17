package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func JwtMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// req token
		tokenString := ctx.Request.Header.Get("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusBadRequest, "token not fouund")
			ctx.Abort()
			return
		}
		//decode/validate it
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		token, err := jwt.Parse(tokenString, ekstractToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := uint(claims["sub"].(float64))

			//attach to req
			ctx.Set("sub", user)
			//continue
			ctx.Next()

			return

		}
		ctx.JSON(http.StatusForbidden, "invalid token")
		ctx.Abort()
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	return []byte(os.Getenv("SECRET")), nil

}
