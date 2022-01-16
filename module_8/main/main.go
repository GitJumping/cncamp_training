package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"http01/metrics"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	Host string
	Port string
	Addr string
)

func init() {
	host := os.Getenv("SERVER_ADDR")
	if host == "" {
		Host = "0.0.0.0"
	} else {
		Host = host
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		Port = "80"
	} else {
		Port = port
	}
	Addr = fmt.Sprintf("%s:%s", Host, Port)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	metrics.Register()
	mux := http.NewServeMux() //初始化Handler
	mux.HandleFunc("/", GetMyRequest)
	mux.HandleFunc("/healthz", Healthz)
	mux.Handle("/metrics", promhttp.Handler())

	svc := http.Server{ //初始化server
		Addr:    Addr,
		Handler: mux,
	}
	//err := svc.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 监听系统信号：即将系统信号抽象成os.Signal通道
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Fatalf("httpserver start fail：%s\n", err)
		}
	}()
	//log.Printf("server starting at %s", Addr)
	glog.V(4).Info("server starting at ", Addr)
	<-signalChan // 没有信号时在这里阻塞，保证程序持续运行
	glog.V(4).Info("get signal server stop")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := svc.Shutdown(ctx); err != nil { //开始shutdown，并设置5s超时时间
		glog.Fatalf("shutdown fail %s\n", err)
	}
	glog.V(4).Info("server quit")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// GetMyRequest “/”路径函数
func GetMyRequest(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	//获取全部请求头
	h := r.Header
	//作业二：从请求携带的参数中获取想要获取的环境变量，并返回到reqponse header
	query := r.URL.Query()
	env := query.Get("env")
	version := os.Getenv(env)
	if version != "" {
		w.Header().Set("version", version)
	}

	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	glog.V(4).Info("delay：", delay, "millisecond")
	//作业一：将获取的请求头循环返回到reqponse header
	for k, v := range h {
		//fmt.Println(k, v)
		value := strings.Join(v, ",")
		w.Header().Set(k, value) //写入到ResponseHeader
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v)) //返回到页面
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello World!"))
	glog.V(4).Info(logFmt(r))
}

// Healthz 健康检查函数
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	glog.V(4).Info(logFmt(r))
}

// 日志格式化 并获取真实IP
func logFmt(r *http.Request) string {
	// 获取客户端IP
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	AccessLog := fmt.Sprintf("clientip:%s,ststuscode:%d,URL:%s", ip, 200, r.URL)
	return AccessLog
}
