# ClientConsum
Api client para KuberProject con servicios rest implementando el framework gin-gonic

Es servicio se encarga de consumir a <a href="https://github.com/theboshy/KuberProject"> **KuberProject** </a> <img style="display:inline-block" width="40" heigth="40" src="https://png.icons8.com/ios/50/000000/developer.png">

Despues de resolver la solicitud por **API REST** , se conectara mediante el protocolo **buf** , por *RPC* a el servidor rpc/**tcp** 
en *kuberproject*

```sh
$ cd ./[<project_path>]

$ protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto

$ eval $(minikube docker-env)
$ docker build -t [<docker_image_name>] -f Dockerfile.api .

$ kubectl apply -f api-deployment.yaml

```

> support **[net/http/pprof]**

### Instalar 
```sh
$ go get github.com/DeanThompson/ginpprof
```

### Profiler End-routers
``` go
GET("/debug/pprof/")
GET("/debug/pprof/heap")
GET("/debug/pprof/goroutine")
GET("/debug/pprof/block")
GET("/debug/pprof/threadcreate")
GET("/debug/pprof/cmdline")
GET("/debug/pprof/profile")
GET("/debug/pprof/symbol")
POST("/debug/pprof/symbol")
GET("/debug/pprof/trace")
GET("/debug/pprof/mutex")
```
### Uso
```sh
$  go tool pprof goprofex http://localhost:3000/profiler/debug/pprof/profile/
```
