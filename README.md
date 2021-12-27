## 作业要求
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来
- 如何实现安全保证
- 七层路由规则
- 考虑 open tracing 的接入

## 环境
- minikube
- HOST_IP: 通过 ```minikube ip```得到本地HOST_IP为```192.168.49.2```
- HOST_PORT: 通过 ```kubectl get svc -n istio-system```得到本地HOST_PORT为```32208```

## 启动命令
可通过make命令部署:
```sh
make deploy
```
主要的YAML文件包括三个展示```tracing```的```deployment```文件以及一个```istio gateway和virtualservice```文件
![Screenshot from 2021-12-27 02-36-07](https://user-images.githubusercontent.com/10457633/147447693-f01bf2ab-d5ec-45b9-b8ec-ae3356362d73.png)
