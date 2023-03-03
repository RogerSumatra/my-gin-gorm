package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"my-gin-gorm/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *handler) RequiredAuth(ctx *gin.Context) {
	// get cookie off req
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token subctx
		var user entity.User
		h.db.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		ctx.Set("user", user)
		//continue
		ctx.Next()

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
