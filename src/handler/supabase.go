 package handler

// import (
// 	"fmt"

// 	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
// 	"github.com/gin-gonic/gin"
// )

// func (h *handler) uploadFile(ctx *gin.Context) {
// 	supClient := supabasestorageuploader.NewSupabaseClient(
// 		"https://prtlclzsqdfetjatgvbw.supabase.co",
// 		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InBydGxjbHpzcWRmZXRqYXRndmJ3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgxMDA4NjUsImV4cCI6MTk5MzY3Njg2NX0.MPgb1dbjRlkX0tK5tUXQFSDLNvOuB_Pxx4tkNQMEMqE",
// 		"api-service",
// 		"studioPicture",
// 	)
// 	file, err := ctx.FormFile("avatar")
// 	if err != nil {
// 		ctx.JSON(400, gin.H{"data": err.Error()})
// 		return
// 	}
// 	link, err := supClient.Upload(file)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"data": err.Error()})
// 		return
// 	}
// 	ctx.JSON(200, gin.H{"data": link})
// }

// func (h *handler) deleteFile(ctx *gin.Context) {
// 	supClient := supabasestorageuploader.NewSupabaseClient(
// 		"https://prtlclzsqdfetjatgvbw.supabase.co",
// 		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InBydGxjbHpzcWRmZXRqYXRndmJ3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgxMDA4NjUsImV4cCI6MTk5MzY3Njg2NX0.MPgb1dbjRlkX0tK5tUXQFSDLNvOuB_Pxx4tkNQMEMqE",
// 		"api-service",
// 		"studioPicture",
// 	)
// 	linkFile := ctx.Request.FormValue("linkfile")

// 	fmt.Println(linkFile)

// 	data, err := supClient.DeleteFile(linkFile)

// 	if err != nil {
// 		ctx.JSON(500, gin.H{"data": err.Error()})
// 		return
// 	}
// 	ctx.JSON(200, gin.H{"data": data})
// }


