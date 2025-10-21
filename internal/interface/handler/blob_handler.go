package handler

import (
    "net/http"
    "phillip-cms-jobs/internal/app/usecase"

    "github.com/gin-gonic/gin"
)

type BlobHandler struct {
    blobUsecase *usecase.BlobUsecase
}

func NewBlobHandler() *BlobHandler {
    return &BlobHandler{
        blobUsecase: usecase.NewBlobUsecase(), // สร้าง usecase
    }
}

// UploadFile รับ POST /upload
func (h *BlobHandler) UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
        return
    }

    if err := h.blobUsecase.Upload(file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
