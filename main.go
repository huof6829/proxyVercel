package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/webhook", handleAPI)

	port := ":8443"
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proxy handleAPI")

	targetURL := "http://ec2-52-77-241-219.ap-southeast-1.compute.amazonaws.com:8443/webhook"
	// targetURL := "http://localhost:8888"

	url, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing target URL: %s", err), http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	// 修改请求主机头，确保目标服务器能正确处理
	r.Host = url.Host

	// 执行反向代理
	proxy.ServeHTTP(w, r)

}
