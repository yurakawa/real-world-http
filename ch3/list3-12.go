package main

import (
    "log"
    "net/http"
    "net/http/cookiejar"
    "net/http/httputil"
)

func main() {
    jar, err :=cookiejar.New(nil) // クッキーを保存するcookiejar のインスタンスを作成する
    if err != nil {
        panic(err)
    }
    client := http.Client{ // クッキーを保存可能なhttp.Client インスタンスを作成する
        Jar: jar,
    }
    for i := 0; i < 2; i++ { // クッキーは初回アクセスでクッキーを受信し、2回目以降のアクセスでクッキーをサーバに対して送信する仕組みなので、2回アクセスする
        resp, err := client.Get("http://localhost:18888/cookie")  // http.Get()の代わりに、作成したクライアントのGet()メソッドを使ってアクセスする
        if err != nil {
            panic(err)
        }
        dump, err := httputil.DumpResponse(resp, true)
        if err != nil {
            panic(err)
        }
        log.Println(string(dump))
    }
}