package handler

import (
	"fmt"
	"my-gin-gorm/sdk/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	http   *gin.Engine
	config config.Interface
	db     *gorm.DB
}

func Init(config config.Interface, db *gorm.DB) *handler {
	rest := handler{
		http:   gin.Default(),
		config: config,
		db:     db,
	}

	rest.registerRoutes()
	return &rest
}

func (h *handler) registerRoutes() {
	h.http.GET("/", h.ping)

	v1 := h.http.Group("api/v1")

	v1.POST("/post", h.createPost)
	v1.GET("/post", h.getListPost)
	v1.GET("/post/:post_id", h.getPost)
	v1.PUT("/post/:post_id", h.updatePost)
	v1.DELETE("/post/:post_id", h.deletePost)

	v1.POST("/signup", h.signUp)
	v1.POST("/login", h.login)
	v1.GET("/validate", h.RequiredAuth, h.validate)

}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil, nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
