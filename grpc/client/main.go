package main

import (
	"fmt"
	"grpc/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Server struct {
}

func main() {
	clientConnection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewMathServiceClient(clientConnection)
	g := gin.Default()

	g.GET("add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})

			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})

			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
	})
	g.GET("subtract/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})

			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})

			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Subtract(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
	})

	if err := g.Run(":3000"); err != nil {
		log.Fatal("Faild to run client %v", err)
	}
}
