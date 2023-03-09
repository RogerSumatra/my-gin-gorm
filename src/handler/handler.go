package handler

import (
	"fmt"
	"my-gin-gorm/sdk/config"
	"net/http"

	"my-gin-gorm/middleware"

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

	//for studio
	v1.POST("/studio", h.createStudio)
	v1.GET("/studio", h.getListStudio)
	v1.GET("/studio/:studio_id", h.getStudio)
	v1.PUT("/studio/:studio_id", h.updateStudio)
	v1.DELETE("/studio/:studio_id", h.deleteStudio)

	//for user
	v1.POST("/user/signup", h.signUp)
	v1.POST("/user/login", h.login)
	v1.GET("/validate", middleware.JwtMiddleware(h.db), h.validate)

	//for comment
	// v1.POST("/studio", h.createStudio)
	// v1.GET("/studio", h.getListStudio)
	// v1.GET("/studio/:studio_id", h.getStudio)
	// v1.PUT("/studio/:studio_id", h.updateStudio)
	// v1.DELETE("/studio/:studio_id", h.deleteStudio)

	//for facility
	// v1.POST("/studio", h.createStudio)
	// v1.GET("/studio", h.getListStudio)
	// v1.GET("/studio/:studio_id", h.getStudio)
	// v1.PUT("/studio/:studio_id", h.updateStudio)
	// v1.DELETE("/studio/:studio_id", h.deleteStudio)

	//supabase
	//v1.POST("/upload", h.uploadFile)
	// v1.DELETE("/delete", h.deleteFile)

}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
