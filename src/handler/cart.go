package handler

// import (
// 	"my-gin-gorm/src/business/entity"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func (h *handler) createCart(ctx *gin.Context) {
// 	var hourParam entity.HourParam
// 	if err := h.BindParam(ctx, &hourParam); err != nil {
// 		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
// 		return
// 	}

// 	var cartRequest entity.CartRequest
// 	if err := h.BindBody(ctx, &cartRequest); err != nil {
// 		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
// 		return
// 	}

// 	var cart entity.Cart
// 	cart.NameContact = cartRequest.NameContact
// 	cart.PhoneNumber = cartRequest.PhoneNumber
// 	cart.Email = cartRequest.Email
// 	// cart.StudioCheckIn = 
// 	// cart.StudioCheckOut = 

// 	if err := h.db.Create(&cart).Error; err != nil {
// 		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// }
