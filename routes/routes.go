package routes

import (
	"markdown-notes/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.POST("/save-note", controller.SaveNoteHandler)
	r.POST("/upload-markdown", controller.UploadMarkdownHandler)
	r.POST("/check-grammar", controller.CheckGrammarHandler)
	r.GET("/notes", controller.ListNotesHandler)
	r.GET("/notes/:id", controller.RenderNoteHandler)
}
