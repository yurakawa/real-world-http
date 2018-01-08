package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
)

func main() {
    // 証明書(Certification)を読み込む
    cert, err := ioutil.ReadFile("ca.crt")
    if err != nil {
        panic(err)
    }
    certPool := x509.NewCertPool() // X509 は ISOで定められた証明書の形式
    certPool.AppendCertsFromPEM(cert) // PEM は BASE64エンコードしたバイナリにヘッダとフッダをつけたデータ構造
    tlsConfig := &tls.Config{
        RootCAs: certPool,
        // InsecureSkipVerify: true, // 証明書を確認しない設定
    }
    tlsConfig.BuildNameToCertificate()

    // クライアントを作成
    client := &http.Client {
        Transport: &http.Transport {
            TLSClientConfig: tlsConfig,
        },
    }

    // 通信を行う
    resp, err := client.Get("https://localhost:18443")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    dump, err :=httputil.DumpResponse(resp, true)
    if err != nil {
        panic(err)
    }
    log.Println(string(dump))
}

