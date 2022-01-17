# cncamp-a05

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

### 通过Istio Ingress Gateway发布http-server
![Screenshot from 2021-12-27 02-43-56](https://user-images.githubusercontent.com/10457633/147448062-ae8926ed-99b7-4663-8036-ff7fc353411f.png)

### 通过VirtualService对路由的配置转发特定路径
![Screenshot from 2021-12-27 02-48-11](https://user-images.githubusercontent.com/10457633/147448426-01b6151c-fc2b-4105-a83c-4b7dbd85db16.png)
![Screenshot from 2021-12-27 02-48-44](https://user-images.githubusercontent.com/10457633/147448430-710cf353-fcfe-4b91-8da0-3de004208f68.png)

### 在http-server代码中按照教程加入请求转发逻辑，并将镜像分别推送到docker hub, ```tracing```测试结果如下：
![Screenshot from 2021-12-27 02-56-36](https://user-images.githubusercontent.com/10457633/147449109-d8d0d448-26e2-436c-bda9-403d04c9d9e6.png)

### https配置遇到些问题，暂时未能实现


[极客时间云原生第0期毕业总结.pdf](https://github.com/brian880120/cncamp-a05/files/7878362/0.pdf)


