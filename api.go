package main

import (
	"flag"
	"fmt"
	"mime"
	"net/http"
	"time"
)

type fileServer struct {
}

var (
	port   = flag.Int("p", 8099, "监听端口")
	second = flag.Int("s", 0, "延时秒数")
)

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 模拟网络请求延时
	time.Sleep(time.Duration(*second) * time.Second)

	// 允许跨域请求
	w.Header().Set("Access-Control-Allow-Origin", "*")

	server := http.FileServer(http.Dir("./"))
	server.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	// 增加json文件的支持
	mime.AddExtensionType(".json", "application/json")

	var s fileServer

	fmt.Printf("开始监听端口：%d, 延时时间：%d秒\n", *port, *second)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.Handler(&s))
	if err != nil {
		println(err.Error())
	}
}
