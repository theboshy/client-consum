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
)

func main() {
	/*targetport (yaml config) : gcd-service:3001*/
	conn, err := grpc.Dial("gcd-service:3001", grpc.WithInsecure())
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
	r.GET("/gcd/:firstNumber/:secondNumber", func(c *gin.Context) {
		a, err := strconv.ParseUint(c.Param("firstNumber"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(c.Param("secondNumber"), 10, 64)
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

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof
func HeapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof
func GoroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// BlockHandler will pass the call from /debug/pprof/block to pprof
func BlockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof
func ThreadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof
func CmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof
func ProfileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof
func SymbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof
func TraceHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
	}
}

// MutexHandler will pass the call from /debug/pprof/mutex to pprof
func MutexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("mutex").ServeHTTP(ctx.Writer, ctx.Request)
	}
}
