package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    resp, err := http.Get("http://localhost:18888")
    if err != nil {
        // エラー発生
        panic(err)
    }
    // このスコープが抜けたところで必ずクローズ
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body) // ioutil.ReadAllでサーバレスポンスを最後まで一括で読み切る
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
}

