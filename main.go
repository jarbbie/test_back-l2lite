package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jarbbie/test_back-l2lite/parser"
)

func main() {
	r := gin.Default()

	// POST /uploads
	r.POST("/uploads", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		// Save the uploaded file
		filePath := "./tmp/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
			return
		}

		// Parse the Excel file
		parsedRows, err := parser.ParseExcel(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Excel"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"rows": parsedRows})
	})

	r.Run(":8080")
}
