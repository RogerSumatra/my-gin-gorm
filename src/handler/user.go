package handler

import (
	"net/http"
	"os"
	"time"

	"my-gin-gorm/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) signUp(ctx *gin.Context) {
	//get username email & pass
	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := h.BindBody(ctx, &body); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to read body", nil)
	}
	//hasp pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to hash password", nil)
	}
	//create the user
	user := entity.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	if err := h.db.Create(&user).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to create user", nil)
		return
	}

	//respond
	h.SuccessResponse(ctx, http.StatusOK, "user created successfully", nil, nil)
}

func (h *handler) login(ctx *gin.Context) {
	//get email and pass
	var body struct {
		Email    string
		Password string
	}

	if err := h.BindBody(ctx, &body); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to read body", nil)
	}
	//look up req user
	var user entity.User
	h.db.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid email or password", nil)
		return
	}
	//compare sent in pass with saved pass hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invaild email or password", nil)
		return
	}
	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to create token", nil)
		return
	}

	//set cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	//sent back
	h.SuccessResponse(ctx, http.StatusOK, "", nil, nil)
}

func (h *handler) validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
	//h.SuccessResponse(ctx, http.StatusOK, user, nil, nil)
}