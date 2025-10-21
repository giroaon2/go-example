package router

import (
    "github.com/gin-gonic/gin"
    "phillip-cms-jobs/internal/interface/handler"
)

func NewRouter() *gin.Engine {
    r := gin.Default()

    // สร้าง handler (สามารถ inject usecase ที่ต้องการได้)
    blobHandler := handler.NewBlobHandler()

    // กำหนด route
    r.POST("/upload", blobHandler.UploadFile)

    return r
}
