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
		Picture  string
	}

	body.Picture = "https://prtlclzsqdfetjatgvbw.supabase.co/storage/v1/object/public/api-service/blank-profile-picture-gd89c1085b_1280.png"

	if err := h.BindBody(ctx, &body); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to read body", nil)
	}
	//hasp pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to hash password", nil)
	}
	//create the user
	user := entity.User{Username: body.Username, Email: body.Email, Password: string(hash), Picture: body.Picture}

	if err := h.db.Create(&user).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to create user", nil)
		return
	}

	//respond
	h.SuccessResponse(ctx, http.StatusOK, "user created successfully", nil)
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
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid email or password", nil)
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
	ctx.JSON(http.StatusOK, gin.H{
		"ID":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"balance":  user.Balance,
		"token":    tokenString,
	})
}

func (h *handler) validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (h *handler) getUser(ctx *gin.Context) {
	var userParam entity.UserParam

	if err := h.BindParam(ctx, &userParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var user entity.User
	if err := h.db.Model(&user).Where(&userParam).First(&user).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "couldn't get user data", nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user data found",
		"data":    user,
	})

}

func (h *handler) updateUser(ctx *gin.Context) {
	var userParam entity.UserParam
	if err := h.BindParam(ctx, &userParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var userBody entity.UserBody
	if err := h.BindBody(ctx, &userBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
		return
	}

	var user entity.User
	user.ID = uint(userParam.UserID)
	user.Name = userBody.Name
	user.ShortName = userBody.ShortName
	user.PhoneNumber = userBody.PhoneNumber

	if err := h.db.Model(user).Where(userParam).Updates(&user).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "studio data update failed", nil)
	}

	h.SuccessResponse(ctx, http.StatusOK, "update user data success", user)
}
