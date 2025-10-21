package main

import (
    "fmt"
    "sync"
)

func fetchData(url string, wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()
    // จำลองโหลดข้อมูล
    ch <- fmt.Sprintf("Data from %s", url)
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan string, 3)

    urls := []string{"A", "B", "C"}

    for _, url := range urls {
        wg.Add(1)
        go fetchData(url, &wg, ch)
    }

    wg.Wait()
    close(ch)

    for data := range ch {
        fmt.Println(data)
    }
}
