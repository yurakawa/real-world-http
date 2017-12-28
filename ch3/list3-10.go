package main

import (
    "bytes"
    "io"
    "log"
    "mime/multipart"
    "net/http"
    "os"
)

func main() {
    var buffer bytes.Buffer // マルチパート部を組み立てた後のバイト列を格納するバッファを宣言する
    writer := multipart.NewWriter(&buffer) // マルチパートを組み立てるライターを作る
    writer.WriteField("name", "Michael Jackson") // ファイル以外のフィールドは、WriteField() メソッドを使って登録する

    // ファイルを読み込む操作
    fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")  // 個別のファイル書き込みのio.Writerを作る
    if err != nil {
        panic(err)
    }
    readFile, err := os.Open("photo.jpg") // ファイルを開く
    if err != nil {
        panic(err)
    }
    defer readFile.Close()
    io.Copy(fileWriter, readFile) // io.Copy() を使って、ファイルの全コンテンツを、ファイル書き込み用のio.Writerにコピーします

    writer.Close() // マルチパートのio.Writerをクローズし、バッファにすべてを書き込む

    resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
    if err != nil {
        panic(err)
    }
    log.Println("Status", resp.Status)
}

