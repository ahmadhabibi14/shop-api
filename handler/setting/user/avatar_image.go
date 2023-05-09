package user

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"shop-api/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func EditAvatar(c *gin.Context) {
	uid, errs := config.ExtractTokenID(c)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	user_id := fmt.Sprint(uid)
	file, _ := c.FormFile("file")
	if isImage(file) != "image" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": isImage(file)})
		return
	}
	avatar_file_path := fmt.Sprintf("assets/img/avatar/%s.png", user_id)
	err := c.SaveUploadedFile(file, avatar_file_path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File upload successfully !!"})
}

func isImage(file *multipart.FileHeader) string {
	f, err := file.Open()
	if err != nil {
		return "Could not open uploaded file"
	}
	defer f.Close()

	// Check that file is an image
	img, format, err := image.DecodeConfig(f)
	if err != nil {
		return "Uploaded file is not an image"
	}

	// Check that the file extension is allowed
	if !allowedExtension(format) {
		return "Uploaded file extension is not allowed"
	}

	// Check that the image has the same dimension
	if img.Width != img.Height {
		return "Uploaded image does not have the same dimension, recommend to 500x500"
	}

	// Check that the image meets the minimum and maximum pixel dimensions
	if (img.Width < minPixelWidth) && (img.Height < minPixelHeight) {
		return fmt.Sprintf("Uploaded image does not meet minimum pixel dimensions of %dx%d", minPixelWidth, minPixelHeight)
	}
	if (img.Width > maxPixelWidth) && (img.Height > maxPixelHeight) {
		return fmt.Sprintf("Uploaded image exceeds maximum pixel dimensions of %dx%d", maxPixelWidth, maxPixelHeight)
	}

	// Check if the file size is not too large
	max := float64(maxFileSize) / 1000000
	if file.Size > maxFileSize {
		return fmt.Sprintf("Uploaded file exceeds maximum size of %.2f MB", max)
	}

	return "image"
}

const (
	minPixelWidth  = 100
	minPixelHeight = 100
	maxPixelWidth  = 1000
	maxPixelHeight = 1000
	maxFileSize    = 1024 * 1024 * 5 // 5 MB
)

func allowedExtension(format string) bool {
	for _, ext := range allowedExtensions {
		if strings.EqualFold(ext, format) {
			return true
		}
	}
	return false
}

var allowedExtensions = []string{
	"jpeg", "jpg", "png",
}
