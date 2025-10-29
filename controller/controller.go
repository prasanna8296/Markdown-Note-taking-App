package controller

import (
	"io"
	"markdown-notes/config"
	"markdown-notes/models"
	"markdown-notes/services"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"gorm.io/gorm"
)

// SaveNoteHandler saves a new note to the database
func SaveNoteHandler(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.CreatedAt = time.Now()

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "note saved",
		"note":    note,
	})
}

// UploadMarkdownHandler handles file upload and saves markdown content
func UploadMarkdownHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open uploaded file"})
		return
	}
	defer src.Close()

	// Read file content
	contentBytes, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot read file"})
		return
	}

	content := strings.TrimSpace(string(contentBytes))

	note := models.Note{
		Title:     strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)),
		Filename:  file.Filename,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Markdown uploaded and saved to database",
		"note":    note,
	})
}

// RenderNoteHandler renders a note as HTML
func RenderNoteHandler(c *gin.Context) {
	id := c.Param("id")
	var note models.Note

	if err := config.DB.First(&note, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	html := markdown.ToHTML([]byte(note.Content), nil, nil)
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

// CheckGrammarHandler checks grammar using external service
func CheckGrammarHandler(c *gin.Context) {
	var payload struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grammarResult, err := services.CheckGrammar(payload.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grammarResult)
}

// ListNotesHandler lists all notes from the database
func ListNotesHandler(c *gin.Context) {
	var notes []models.Note

	if err := config.DB.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}
