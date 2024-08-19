package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 默认443  目录/文件名
// curl -X POST -H "Content-Type: application/json" -d '{"username": "user", "password": "pass"}'  https://proxy-vercel-blue.vercel.app/api/webhook
// vercel.json building 指定程序入口   "src": "/.*" 无效

// https://proxy-vercel-blue.vercel.app/api/redis

func HandleWebhookTwitter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proxy HandleWebhookTwitter")

	targetURL := "http://ec2-44-214-134-96.compute-1.amazonaws.com:8888"

	url, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing target URL: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("r.Host: %+v\n", r.Host)

	proxy := httputil.NewSingleHostReverseProxy(url)
	r.Host = url.Host

	fmt.Printf("url.Host: %+v\n", url.Host)
	fmt.Printf("r.URL.Path: %+v\n", r.URL.Path)

	// 执行反向代理
	proxy.ServeHTTP(w, r)

}
