# cncamp-training
go-base-to-cloud

# Week1
实现了一小点功能。。
见  [代码main.go](https://github.com/GitJumping/cncamp_training/week_1/main/main.go)

# Module2
实现功能见 [代码](https://github.com/GitJumping/cncamp_training/module_2/main)

# Module_8
代码见 [代码main.go](https://github.com/GitJumping/cncamp_training/module_8/main/main.go)
kubernetes部署yaml
[deploy-configmap.yaml](https://github.com/GitJumping/cncamp_training/module_8/main/deploy-configmap.yaml)
[deploy-deployment.yaml](https://github.com/GitJumping/cncamp_training/module_8/main/deploy-deployment.yaml)
[deploy-ingress.yaml](https://github.com/GitJumping/cncamp_training/module_8/main/deploy-ingress.yaml)
[deploy-service.yaml](https://github.com/GitJumping/cncamp_training/module_8/main/deploy-service.yaml)

# Module_10
实现功能见 [代码main.go](https://github.com/GitJumping/cncamp_training/module_10/main/main.go)
- 增加随机访问延迟
- 增加 /metrics 接口

# Module_12
支持透传header，支持环境变量设置调用后端服务
实现功能见 [代码main.go](https://github.com/GitJumping/cncamp_training/module_12/main/main.go)

istio自动注入的命令
```shell
kubectl create ns istio-demo
kubectl label ns istio-demo istio-injection=enabled --overwrite
kubectl get ns -L istio-injection
```

安装jaeger
```shell
kubectl apply -f jaeger.yaml
```

配置istio gateway
```shell
kubectl apply -f istio-specs.yaml
```

查询ip，请求服务，查询tracing
```shell
k get svc -nistio-system
curl -H "Host: demo.cncv.vip" 10.108.82.159/hello -v
```
