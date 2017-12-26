package main

import (
    "log"
    "net/http"
)


// curl --head http://localhost:18888
func main() {
    resp, err := http.Head("http://localhost:18888")
    if err != nil {
        panic(err)
    }
    log.Println("Status:", resp.Status)
    log.Println("Header:", resp.Header)
}

