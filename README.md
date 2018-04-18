# Client Consum API
Api client para KuberProject con servicios rest implementando el framework gin-gonic

Este servicio se encarga de consumir a <a href="https://github.com/theboshy/KuberProject"> **KuberProject** </a> <img style="display:inline-block" width="40" heigth="40" src="https://png.icons8.com/ios/50/000000/developer.png">, por medio de conexion **rpc**

Despues de resolver la solicitud por **API REST** , se conectara mediante el protocolo **buf** , a el servidor **tcp** 
en *kuberproject*

### Requerimientos 
* [Minikube](https://github.com/kubernetes/minikube) - (mini) servicio local de kubernetes 
* [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - herramienta de línea de comandos de Kubernetes 
* [ProtoBufCompiler](https://github.com/google/protobuf) - compilador de proto buf
* [GoProtoBufCompiler](https://github.com/golang/protobuf) - compilador de proto buf para **golang**

### Build

```sh
$ cd ./[<project_path>]

# cosntruir protobuf *pb*
$ protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto

#minikube mantiene un servicio docker el cual podemos usar para generar nuestro contenedor e imagen
$ eval $(minikube docker-env)
#generar la imagen con el servicio ClientConsum
$ docker build -t [<docker_image_name>] -f Dockerfile.api .

#generar el nodo contenedor del servicio cconsum
$ kubectl apply -f api-deployment.yaml

```

> support **[net/http/pprof]**

### Instalar pprof
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
### Uso de profiler
```sh
$  go tool pprof goprofex http://localhost:3000/profiler/debug/pprof/profile/
```
