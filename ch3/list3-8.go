package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    file, err := os.Open("list3-8.go")
    if err != nil {
        // 送信失敗
        panic(err)
    }
    resp, err := http.Post("http://localhost:18888", "text/plain", file)
    if err != nil {
        // 送信失敗
        panic(err)
    }
    log.Println("Status:", resp.Status)
}

