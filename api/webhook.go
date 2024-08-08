package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proxy HandleWebhook")

	targetURL := "http://ec2-52-77-241-219.ap-southeast-1.compute.amazonaws.com:8443"
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
