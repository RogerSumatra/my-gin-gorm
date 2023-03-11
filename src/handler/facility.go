package handler

import (
	"my-gin-gorm/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) createFacility(ctx *gin.Context) {
	var facilityBody entity.FacilityBody

	if err := h.BindBody(ctx, &facilityBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	var facility entity.Facility

	facility.StudioID = facilityBody.StudioID
	facility.Content = facilityBody.Content
	// facility.Picture = facilityBody.Picture
	//tambah

	if err := h.db.Create(&facility).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Successfully created facility", nil)

}

func (h *handler) updateFacility(ctx *gin.Context) {
	var facilityParam entity.FacilityParam
	if err := h.BindParam(ctx, &facilityParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var facilityBody entity.FacilityBody
	if err := h.BindBody(ctx, &facilityBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
		return
	}

	var facility entity.Facility
	facility.ID = uint(facilityParam.FacilityID)
	facility.Content = facilityBody.Content
	// facility.Picture = facilityBody.Picture

	if err := h.db.Model(facility).Where(facilityParam).Updates(&facility).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "facility update failed", nil)
	}

	h.SuccessResponse(ctx, http.StatusOK, "update facility success", facility)
}
