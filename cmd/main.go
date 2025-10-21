package main

import (
    "log"
    "phillip-cms-jobs/internal/interface/router"
)

func main() {
    r := router.NewRouter() // เรียกสร้าง router จาก router.go

    // รัน server บน port 8080
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
