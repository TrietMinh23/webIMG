package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"webimg/config"
	"webimg/initializers"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Upload(c *gin.Context) {
	cfg := config.GetConfig()

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No image found",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot open file",
		})
		return
	}
	defer src.Close()

	objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)

	uploadInfo, err := initializers.MinioClient.PutObject(context.Background(), cfg.MinioConfig.BucketName, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		log.Printf("Upload failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed"})
		return
	}

	fileURL := fmt.Sprintf("http://localhost:%s/image/%s", cfg.Port, uploadInfo.Key)
	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"url":     fileURL,
	})
}

func GetImg(c *gin.Context) {
	cfg := config.GetConfig()
	filename := c.Param("filename")

	obj, err := initializers.MinioClient.GetObject(context.Background(), cfg.MinioConfig.BucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get object failed"})
		return
	}

	stat, err := obj.Stat()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.Header("Content-Type", stat.ContentType)
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size))
	c.Header("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))
	c.Status(http.StatusOK)
	_, err = io.Copy(c.Writer, obj)
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
