package handler

import (
	"my-gin-gorm/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) createStudio(ctx *gin.Context) {
	var studioBody entity.StudioBody

	if err := h.BindBody(ctx, &studioBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	var studio entity.Studio

	studio.Name = studioBody.Name
	studio.Description = studioBody.Description

	if err := h.db.Create(&studio).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Successfully created new studio", nil)

}

func (h *handler) getListStudio(ctx *gin.Context) {
	var studioParam entity.StudioParam

	if err := h.BindParam(ctx, &studioParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	studioParam.FormatPagination()

	var studios []entity.Studio

	if err := h.db.
		Model(entity.Studio{}).
		Limit(int(studioParam.Limit)).
		Offset(int(studioParam.Offset)).
		Find(&studios).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return

	}

	var totalElements int64

	if err := h.db.
		Model(entity.Studio{}).
		Limit(int(studioParam.Limit)).
		Offset(int(studioParam.Offset)).
		Count(&totalElements).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	studioParam.ProcessPagination(totalElements)

	h.SuccessResponse(ctx, http.StatusOK, "Successfully get list studios", studios)
}

func (h *handler) getStudio(ctx *gin.Context) {
	var studioParam entity.StudioParam

	if err := h.BindParam(ctx, &studioParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	// var PostPayLoad entity.PostBody
	// if err := h.BindBody(ctx, &PostPayLoad); err != nil {
	// 	h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
	// 	return
	// }

	var studio entity.Studio
	if err := h.db.Model(&studio).Where(&studioParam).First(&studio).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "couldn't get studio data", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "studio data found", studio)

}

func (h *handler) updateStudio(ctx *gin.Context) {
	var studioParam entity.StudioParam
	if err := h.BindParam(ctx, &studioParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var studioBody entity.StudioBody
	if err := h.BindBody(ctx, &studioBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
		return
	}

	var studio entity.Studio
	studio.ID = uint(studioParam.StudioID)
	studio.Name = studioBody.Name
	studio.Description = studioBody.Description

	if err := h.db.Model(studio).Where(studioParam).Updates(&studio).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "studio data update failed", nil)
	}

	h.SuccessResponse(ctx, http.StatusOK, "update studio data success", studio)
}

func (h *handler) deleteStudio(ctx *gin.Context) {
	var studioParam entity.StudioParam
	if err := h.BindParam(ctx, &studioParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var studio entity.Studio
	studio.ID = uint(studioParam.StudioID)
	if err := h.db.Delete(&studio).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "delete studio data failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "delete studio data success", nil)
}
