package handlers

import (
	"errors"
	"go-pdf2jpeg/service"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// body represents the request body
type body struct {
	FileName string `json:"filename"`
}

// ConvertPdf converts a PDF file to a JPEG image
func POSTConvertPdf(ctx *gin.Context) {
	var body body

	// Create a mutex to lock the goroutine
	mutex := &sync.Mutex{}
	mutex.Lock()
	
	// Create a channel to notify when the goroutine is finished
	finished := make(chan bool)
	
	// Convert the PDF to JPEG in a goroutine
	go func() {
		defer mutex.Unlock()
		err := ctx.BindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		err = service.PdfToJpeg(body.FileName)
		if err == nil {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Conversion successful", "status": "success"})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("conversion failed").Error()})
		}
		finished <- true
	}()
	<-finished
}
