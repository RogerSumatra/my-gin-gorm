package handler

import (
	"my-gin-gorm/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) createComment(ctx *gin.Context) {
	var commentBody entity.CommentBody

	if err := h.BindBody(ctx, &commentBody); err != nil {
		h.ErrorResponse(ctx, http.StatusUnprocessableEntity, "create comment failed", nil)
		return
	}

	comment := entity.Comment{
		UserID:   commentBody.UserID,
		StudioID: commentBody.StudioID,
		Content:  commentBody.Content,
		Rating:   commentBody.Rating,
	}

	if err := h.db.Create(&comment).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	h.SuccessResponse(ctx, http.StatusOK, "comment created successfully", nil)
}

func (h *handler) getListComment(ctx *gin.Context) {
	var commentParam entity.CommentParam

	if err := h.BindParam(ctx, &commentParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	commentParam.FormatPagination()

	var comments []entity.Comment

	if err := h.db.
		Model(entity.Comment{}).
		Limit(int(commentParam.Limit)).
		Offset(int(commentParam.Offset)).
		Find(&comments).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var totalElements int64

	if err := h.db.
		Model(entity.Comment{}).
		Limit(int(commentParam.Limit)).
		Offset(int(commentParam.Offset)).
		Count(&totalElements).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	commentParam.ProcessPagination(totalElements)

	h.SuccessResponse(ctx, http.StatusOK, "successfully get comment list", comments)
}

func (h *handler) updateComment(ctx *gin.Context) {
	var commentParam entity.CommentParam
	if err := h.BindBody(ctx, &commentParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var commentBody entity.CommentBody
	if err := h.BindBody(ctx, &commentBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
		return
	}

	var comment entity.Comment
	comment.ID = uint(commentParam.CommentID)
	comment.Content = commentBody.Content

	if err := h.db.Model(comment).Where(commentParam).Updates(&comment).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "update comment failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "update comment success", comment)
}

func (h *handler) deleteComment(ctx *gin.Context) {
	var commentParam entity.CommentParam
	if err := h.BindParam(ctx, &commentParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var comment entity.Comment
	comment.ID = uint(commentParam.CommentID)
	if err := h.db.Delete(&comment).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "delete comment failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "delete comment success", nil)
}
