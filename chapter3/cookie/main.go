package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	// クッキーを保存するcookiejar（クッキーの瓶）のインスタンスを作成
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	// クッキーを保存可能なhttp.Clientインスタンスを作成
	client := http.Client{
		Jar: jar,
	}
	// クッキーは初回アクセスでクッキーを受信し、
	// 二回目以降のアクセスでクッキーをサーバーに対して送信する仕組みなので、二回アクセスする
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://localhost:18888/cookie")
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
