package main

import (
    "fmt"
    "sync"
    "time"
)

// Job function ตัวอย่าง
func job(id int) {
    fmt.Printf("Job %d started\n", id)
    time.Sleep(2 * time.Second) // simulate work
    fmt.Printf("Job %d finished\n", id)
}

// StartDailyJob จะ trigger ทุกวันเวลา 6:00AM
func StartDailyJob() {
    go func() {
        for {
            now := time.Now()

            // ตรวจสอบเวลา 6:00AM
            if now.Hour() == 6 && now.Minute() == 0 {
                fmt.Println("Trigger jobs at 6:00AM")

                // Fire-and-forget jobs
                var wg sync.WaitGroup
                jobCount := 3
                wg.Add(jobCount)

                for i := 1; i <= jobCount; i++ {
                    go func(id int) {
                        defer wg.Done()
                        job(id)
                    }(i)
                }

                wg.Wait() // รอให้ทุก job เสร็จ ก่อน loop ต่อไป
                fmt.Println("All daily jobs finished")
            }

            time.Sleep(1 * time.Minute) // check time ทุก 1 นาที
        }
    }()
}

// ตัวอย่าง main
func main() {
    fmt.Println("Server started")
    StartDailyJob() // เริ่ม daily job background

    // สมมติรัน server แบบ infinite loop
    select {}
}
