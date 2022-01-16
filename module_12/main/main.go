package main

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"httpserver/metrics"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var version = os.Getenv("VERSION")
var callName = os.Getenv("CALL_NAME")

// 健康检查
func healthz(w http.ResponseWriter, r *http.Request) {
	for k, vv := range r.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.Header().Set("VERSION", version)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// 配置
func config(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/app/app.json")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// 随机耗时测试
func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(4).Info("entering root handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if callName != "" {
		req, err := http.NewRequest("GET", "http://"+callName+"/hello", nil)
		if err != nil {
			fmt.Printf("%s", err)
		}
		lowerCaseHeader := make(http.Header)
		for key, value := range r.Header {
			lowerCaseHeader[strings.ToLower(key)] = value
		}
		glog.Info("headers:", lowerCaseHeader)
		req.Header = lowerCaseHeader
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			glog.Info("HTTP get failed with error: ", "error", err)
		} else {
			glog.Info("HTTP get succeeded")
		}
		resp.Write(w)
	} else {
		io.WriteString(w, "===================Details of the http request header:============\n")
		for k, v := range r.Header {
			io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		}
		glog.V(4).Infof("Respond in %d ms", delay)
	}
}

func main() {

	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/config", config)
	mux.HandleFunc("/hello", rootHandler)
	mux.Handle("/metrics", promhttp.Handler())
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
