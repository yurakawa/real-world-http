package main

import (
    "strings"
    "net/url"
    "log"
    "net/http"
    "net/http/httputil"
)

func main() {
    values := url.Values{"test": {"value"}}
    reader := strings.NewReader(values.Encode())
    resp, err := http.Post("http://localhost:18888", "text/plain", reader)
    if err != nil {
        panic(err)
    }
    dump, err := httputil.DumpResponse(resp, true)
    if err != nil {
        panic(err)
    }
    log.Println(string(dump))
}

