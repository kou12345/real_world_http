package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// file, err := os.Open("chapter3/post3.8/hoge.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	reader := strings.NewReader("テキスト")

	// ファイルの内容を送信するリクエストを作成
	// resp, err := http.Post("http://localhost:18888", "text/plain", file)
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)

}
