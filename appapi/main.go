package main

import (
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"fmt"
	"log"
	"ClientConsum/mcs"
	"net/http/pprof"
	"io/ioutil"
)

func main() {
	/*targetport (node/yaml config) : gcd-service:3001*/
	/*targetport (local config) : :3001*/
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcdClient := mcs.NewGCDServiceClient(conn)

	r := gin.Default()
	profilerGroup := r.Group("/profiler")
	{
		profilerGroup.GET("/debug/pprof/", IndexHandler())
		profilerGroup.GET("/debug/pprof/heap", HeapHandler())
		profilerGroup.GET("/debug/pprof/goroutine", GoroutineHandler())
		profilerGroup.GET("/debug/pprof/block", BlockHandler())
		profilerGroup.GET("/debug/pprof/threadcreate", ThreadCreateHandler())
		profilerGroup.GET("/debug/pprof/cmdline", CmdlineHandler())
		profilerGroup.GET("/debug/pprof/profile", ProfileHandler())
		profilerGroup.GET("/debug/pprof/symbol", SymbolHandler())
		profilerGroup.POST("/debug/pprof/symbol", SymbolHandler())
		profilerGroup.GET("/debug/pprof/trace", TraceHandler())
		profilerGroup.GET("/debug/pprof/mutex", MutexHandler())
	}
	r.GET("/gcd/compute", func(c *gin.Context) {
		a, err := strconv.ParseUint(c.Query("firstNumber"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(c.Query("secondNumber"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		}
		req := &mcs.GCDRequest{A: a, B: b}
		if res, err := gcdClient.Compute(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.POST("/gcd/file", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter file"})
			return
		}
		reda,_ :=file.Open()
		byteContainer, err := ioutil.ReadAll(reda)
		//fmt.Print(len(byteContainer))
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		req := &mcs.FileRequest{BinaryFile:byteContainer,FileSize:file.Size,FileName:file.Filename}
		if res, err := gcdClient.SaveFile(c, req); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"outputpath": fmt.Sprint(res.OutPath),
				"message": fmt.Sprint(res.Message),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})





	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

func HeapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func GoroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func BlockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func ThreadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func CmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

func SymbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}

func TraceHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
	}
}

func MutexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("mutex").ServeHTTP(ctx.Writer, ctx.Request)
	}
}
