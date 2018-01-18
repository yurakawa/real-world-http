package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
    "io"
    "net/http"
    "os"
    "strings"
)

// GitHubで取得したものを貼り付けてください
var clientID = "xxxxxxxxxxxxxxxxxxxx"
var clientSecret = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

var redirecdtURL = "https://localhost:18888"
var state = "your state"

func main() {
    // OAuth2の接続先などの情報
    conf := &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        Scopes:       []string{"user:email", "gist"},
        Endpoint:     github.Endpoint,
    }
    // これをこれから初期化する
    var token *oauth2.Token

    // ローカルにすでに保存済み?
    file, err := os.Open("access_token.json")
    if os.IsNotExist(err) {
        // 初回アクセス
        // まず認可画面のURLを取得
        url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

        // コールバックを受け取るウェブサーバをセットアップ
        code := make(chan string)
        var server *http.Server
        server = &http.Server{
            Addr: ":18888",
            Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                // クエリパラメータからcodeを取得し、ブラウザを閉じる
                w.Header().Set("Content-Type", "text/html")
                io.WriteString(w, "<html><script>window.open('about:blank', '_self').close()</script></html>")
                w.(http.Flusher).Flush()
                code <- r.URL.Query().Get("code")
                // サーバも閉じる
                server.Shutdown(context.Background())
            }),
        }
        go server.ListenAndServe()

        // ブラウザで認可画面を開く
        // GitHubの認可が完了すれば上記のサーバにリダイレクトされて、Handlerが実行される
        open.Start(url)

        // 取得したコードをアクセストークンに交換
        token, err := conf.Exchange(oauth2.NoContext, <-code)
        if err != nil {
            panic(err)
        }

        // アクセストークンをファイルに保存
        file, err = os.Create("access_token.json")
        if err != nil {
            panic(err)
        }
        json.NewEncoder(file).Encode(token)
    } else if err == nil {
        // 一度認可をしてローカルに保存済み
        token = &oauth2.Token{}
        json.NewDecoder(file).Decode(token)
    } else {
        panic(err)
    }
    client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))
    // ここで様々なことを行う
}

