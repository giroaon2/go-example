package main

import (
    "fmt"
    "net/http"
    "sync"
)

func callAPI(wg *sync.WaitGroup, url string, id int) {
    defer wg.Done() // แจ้งว่า goroutine เสร็จ
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Request %d failed: %v\n", id, err)
        return
    }
    defer resp.Body.Close()
    fmt.Printf("Request %d completed with status: %s\n", id, resp.Status)
}

func main() {
    url := "https://httpbin.org/get"
    numRequests := 5
    var wg sync.WaitGroup

    for i := 1; i <= numRequests; i++ {
        wg.Add(1)
        go callAPI(&wg, url, i)
    }

    wg.Wait() // รอทุก goroutine เสร็จ
    fmt.Println("All requests finished")
}
