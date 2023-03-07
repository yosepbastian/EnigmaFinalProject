package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UploadAndDownload struct {
	rg *gin.RouterGroup
}

func (h *UploadAndDownload) handleUpload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while receiving file: " + err.Error()})
		return
	}

	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while writing file to disk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %s uploaded successfully", file.Filename)})

}

func (h *UploadAndDownload) handleDownload(c *gin.Context) {
	filename := c.Param("filename")
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "File not found",
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.FileAttachment(filename, file.Name())

}

func NewUploadAndDownload(router *gin.RouterGroup) *UploadAndDownload {
	newUploadAndDownload := UploadAndDownload{
		rg: router,
	}
	newUploadAndDownload.rg.POST("/upload", newUploadAndDownload.handleUpload)
	newUploadAndDownload.rg.GET("/download/:filename", newUploadAndDownload.handleDownload)
	return &newUploadAndDownload
}
