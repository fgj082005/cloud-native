package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func getEnv(key string) string {
	return os.Getenv(key)
}

func getClient(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])

	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// 将request header 写入response header
func index(w http.ResponseWriter, r *http.Request) {
	reqHeader := r.Header
	for key, value := range reqHeader {
		w.Header().Set(key, strings.Join(value, ","))
	}
	w.WriteHeader(200)
	w.Write([]byte("This in index page"))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 获取环境变量VERSION值写入响应头
	version := getEnv("VERSION")
	w.Header().Set("VERSION", version)
	// 将客户端IP， 返回码输出到server标准输出,访问healthz，返回200
	statusCode := 200
	w.WriteHeader(statusCode)
	clientIp := getClient(r)
	fmt.Printf("client IP is:%s, response code is %d", clientIp, statusCode)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthzHandler)
	http.ListenAndServe(":8000", nil)
}
