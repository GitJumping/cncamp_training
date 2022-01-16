# Go语言特性
## 统一思想-12 factors

## Go语言的诞生

## Go语言环境搭建

## Go语言基本命令
```go
go build
```
```go
go test
```
```go
go vet
```
## 代码版本控制

## Golang playground

## Go的程序
### 控制结构
### switch
### For
### for-range

## 常用数据结构
### 变量和常量
### 类型转换与推导
### 数组
### 切片
### Make New
### Map
### 结构体和指针
### 结构体标签

## 函数
### Main函数
### 参数解析
### Init 函数
### 返回值
### 传递变长参数
### 内置函数
### 回调函数(Callback)
### 闭包
### 方法
### 接口
### 反射机制
#### 基于 struct 的反射
### Go 语言中的面向对象编程
### Json 编解码

## Other
### 错误处理
### defer
### Panic 和 recover

## 多线程
### 并发和并行
### 协程
#### 线程和协程的差异
### Communicating Sequential Process
### channel - 多线程通信
### 通道缓冲
#### 遍历通道缓冲区
#### 单向通道
#### select
### 定时器 Timer
### 上下文 Context
### 如何停止一个子协程
#### 基于 Context 停止子协程

### 线程加锁
sync.Mutex
sync.RWMutex
sync.WaitGroup

### 线程调度
#### 深入理解 Go 语言线程调度
#### Linux 进程的内存使用
#### CPU 对内存的访问

### 用户线程
#### Goroutine
##### MPG的对应关系
##### GMP 模型细节
##### G的状态转换图
#### Goroutine 创建过程
##### 将 Goroutine 放到运行队列上
##### 调度器行为

### 内存管理
#### 堆内存管理
#### TCMalloc 概览

### Go 语言内存分配
#### 内存回收
#### mspan
#### GC 工作流程
#### 三色标记
#### 垃圾回收触发机制

## 包引用与依赖管理
### vendor管理工具
#### Go mod 使用
### GOPROXY 和 GOPRIVATE

## Makefile
### Go 语言项目多采用 Makefile 组织项目编译

## 动手编写一个 HTTP Server
### 理解 Socket
### 理解 net.http 包
### 阻塞 IO 模型
### 非阻塞 IO 模型
### IO 多路复用
### 异步 IO
### Linux epoll
### Go 语言高性能 httpserver 的实现细节

## 调试
### debug
### dlv 的配置
### 更多 debug 方法

## 性能分析（Performance Profiling）
### 分析 CPU 瓶颈
### 其他可用 profiling 工具分析的问题
### 针对 http 服务的 pprof
### 分析 go profiling 结果

## Kubernetes 中常用代码
### Rate Limit Queue


# Docker 核心技术
## 系统架构
### 传统分层架构 vs 微服务
#### 微服务改造
#### 分解原则：基于 size, scope and capabilities
#### 微服务间通讯
##### 点对点
##### API 网关

## 理解 Docker
### 为什么要用 Docker
### 虚拟机和容器运行态的对比
### 安装 Docker
### 容器操作
#### 初识容器
### 容器标准
#### 容器主要特性
##### Namespace
###### Linux 内核代码中 Namespace 的实现
###### Linux 对 Namespace 操作方法
###### 隔离性 – Linux Namespace
###### 隔离性 – Linux Namespace
###### 关于 namespace 的常用操作
##### Cgroups
###### Linux 内核代码中 Cgroups 的实现
###### 可配额/可度量 - Control Groups (cgroups)
##### CPU 子系统
##### Linux 调度器
##### CFS 调度器
##### vruntime 红黑树
##### CFS进程调度
##### cpuacct 子系统
##### Memory 子系统
##### Cgroup driver
#### 文件系统
#### 容器镜像
##### Docker 的文件系统
### Docker 启动
### 写操作
### 容器存储驱动
#### 以 OverlayFS 为例
### OCI 容器标准
### Docker 引擎架构
#### 网络
##### Null 模式
##### 默认模式– 网桥和 NAT
##### Underlay
##### Docker Libnetwork Overlay
##### VXLAN
##### Overlay network sample – Flannel
##### Flannel packet sample
### 创建 docker 镜像

## Dockerfile 的最佳实践
### Factor 之进程
### 理解构建上下文（Build Context）
### 镜像构建日志
#### Build Cache
#### 多段构建（Multi-stage build）
#### Dockerfile 常用指令
#### Dockerfile 最佳实践
##### 目标：易管理、少漏洞、镜像小、层级少、利用缓存。
#### 多进程的容器镜像
##### Docker 镜像管理
##### 基于 Docker 镜像的版本管理
##### Docker tag 与 github 的版本管理合力
##### 镜像仓库
#### Docker 优势







