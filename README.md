[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/theboshy/ClientConsum) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE) [![GitHub Stars](https://img.shields.io/github/stars/theboshy/ClientConsum.svg?style=social)](https://github.com/theboshy/ClientConsum/stargazers)

[![DockerHub](https://img.shields.io/badge/DockerHub-ClientConsum-blue)](https://hub.docker.com/r/devile/clientconsum/)

---

# Client Consum API 🚀

Client Consum API es un servicio REST implementado con el framework **Gin-Gonic** para el proyecto **KuberProject**. 
Este servicio actúa como un cliente que consume datos de [KuberProject](https://github.com/theboshy/KuberProject) a través de una conexión **RPC**.

Tras recibir una solicitud a través de **API REST**, el servicio se conecta mediante el protocolo **Buf** al servidor **TCP** alojado en *KuberProject*.

---

## 📌 Requisitos

Asegúrate de tener instaladas las siguientes herramientas antes de comenzar:

- ✅ [Minikube](https://github.com/kubernetes/minikube) - Entorno local para Kubernetes
- ✅ [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - CLI para Kubernetes
- ✅ [ProtoBuf Compiler](https://github.com/google/protobuf) - Compilador de **Protocol Buffers**
- ✅ [GoProtoBuf Compiler](https://github.com/golang/protobuf) - Compilador de **Protocol Buffers** para Golang
- ✅ [VirtualBox](https://www.virtualbox.org/) - Virtualización de máquinas

---

## ⚙️ Construcción y Despliegue

Ejecuta los siguientes comandos para construir y desplegar el servicio en Kubernetes:

```sh
# Navega al directorio del proyecto
$ cd ./[<project_path>]

# Construcción de archivos protobuf (pb)
$ protoc -I ./mcs --go_out=plugins=grpc:./pb ./mcs/*.proto

# Configurar el entorno Docker en Minikube
$ eval $(minikube docker-env)

# Construcción de la imagen Docker
$ docker build -t [<docker_image_name>] -f Dockerfile.api .

# Desplegar en Kubernetes
$ kubectl apply -f api-deployment.yaml
```

---

## 📂 Archivos de Configuración

- 📄 [Deployment YAML](https://github.com/theboshy/ClientConsum/blob/master/api-deployment.yaml)
- 📄 [Dockerfile API](https://github.com/theboshy/ClientConsum/blob/master/Dockerfile.api)

---

## 🔥 Endpoints de la API

Los endpoints disponibles en la API son los siguientes:

### 📌 Obtener el MCD de dos números
**Endpoint:**
```http
GET /gcd/{a}/{b}
```
**Ejemplo de uso:**
```sh
$ curl http://<API_URL>/gcd/6/2
```
**Respuesta:**
```json
{
  "result": 2
}
```

### 📌 Obtener la URL del servicio en Minikube
**Comando:**
```sh
$ minikube service api-service --url
```
**Salida esperada:**
```sh
http://xxx.xxx.xx.xx:xxxx
```

---

## 🔍 Soporte para **net/http/pprof**

### 📥 Instalación de pprof

[net/http/pprof](https://golang.org/pkg/net/http/pprof/)

```sh
$ go get github.com/DeanThompson/ginpprof
```

### 📌 Rutas de Perfilado (Profiler)

| Método | Ruta |
|--------|--------------------------------|
| GET | `/debug/pprof/` |
| GET | `/debug/pprof/heap` |
| GET | `/debug/pprof/goroutine` |
| GET | `/debug/pprof/block` |
| GET | `/debug/pprof/threadcreate` |
| GET | `/debug/pprof/cmdline` |
| GET | `/debug/pprof/profile` |
| GET | `/debug/pprof/symbol` |
| POST | `/debug/pprof/symbol` |
| GET | `/debug/pprof/trace` |
| GET | `/debug/pprof/mutex` |

### 📌 Uso del profiler

Ejemplo de uso fuera del clúster de **Minikube**:

```sh
$ go tool pprof goprofex http://localhost:3000/profiler/debug/pprof/profile/
$ go tool pprof goprofex http://localhost:3000/profiler/debug/pprof/heap/
```

---

## 📊 Generación de Gráficos con Graphviz

### 📥 Instalación de Graphviz 2.38

#### Usando **pip**

```sh
$ pip install graphviz
```

#### Usando **Chocolatey**

```sh
$ choco install graphviz
```

Configura la variable de entorno para **Graphviz** en el *PATH* del sistema.

![Configuración del PATH](https://user-images.githubusercontent.com/14255055/38958417-cf0b53e6-4322-11e8-993b-df7850a63518.PNG)

### 📌 Uso del Profiler con Graphviz

```sh
$ go tool pprof goprofex http://xxx.xxx.xx.xx:xxxx/profiler/debug/pprof/profile/
$ (pprof) web
```

![Ejemplo de gráfica](https://user-images.githubusercontent.com/14255055/38959396-26b2e0ac-4326-11e8-9ac0-d1827aed1357.PNG)

---

## 📝 Licencia

Este proyecto está licenciado bajo la [Licencia Apache 2.0](https://github.com/gojp/goreportcard/blob/master/LICENSE).
