package main

// HTTP 代理
// 支持CONNECT 方法

import (
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
)


func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Println("proxy start at 0.0.0.0:8888")
	log.Fatal(http.ListenAndServe(":8888", proxy))
}

