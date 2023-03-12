package handler

import (
	"fmt"
	"my-gin-gorm/sdk/config"
	"net/http"

	"my-gin-gorm/middleware"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
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
	supClient := supabasestorageuploader.NewSupabaseClient(
		"https://prtlclzsqdfetjatgvbw.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InBydGxjbHpzcWRmZXRqYXRndmJ3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgxMDA4NjUsImV4cCI6MTk5MzY3Njg2NX0.MPgb1dbjRlkX0tK5tUXQFSDLNvOuB_Pxx4tkNQMEMqE",
		"api-service",
		"",
	)

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
	v1.GET("/user/:user_id", h.getUser)

	//for comment
	v1.POST("/comment", h.createComment)
	v1.GET("/comment", h.getListComment)
	v1.PUT("/comment/:comment_id", h.updateComment)
	v1.DELETE("/comment/:comment_id", h.deleteComment)

	//for facility
	v1.POST("/facility", h.createFacility)
	v1.PUT("/facility/:facility_id", h.updateFacility)

	//supabase
	//r := h.http



	v1.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("avatar")
		if err != nil {
			ctx.JSON(400, gin.H{"data": err.Error()})
			return
		}
		link, err := supClient.Upload(file)
		if err != nil {
			ctx.JSON(500, gin.H{"data": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": link})
	})

	v1.DELETE("file", func(ctx *gin.Context) {
		linkFile := ctx.Request.FormValue("linkfile")
		fmt.Println(linkFile)

		data, err := supClient.DeleteFile(linkFile)

		if err != nil {
			ctx.JSON(500, gin.H{"data": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": data})
	})

}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
