package handler

import (
	"net/http"

	"my-gin-gorm/src/business/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) signUp(ctx *gin.Context) {
	//get email & pass
	var body struct {
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
	user := entity.User{Email: body.Email, Password: string(hash)}

	if err := h.db.Create(&user).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to create user", nil)
		return
	}

	//respond
	h.SuccessResponse(ctx, http.StatusOK, "user created successfully", nil, nil)
}
