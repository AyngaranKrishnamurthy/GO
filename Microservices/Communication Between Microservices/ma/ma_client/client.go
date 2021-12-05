package main

import (
	"fmt"
	"go_workspace/cleanc/ma/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client of Microservice A")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/id/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		req := &proto.Request{Name: string(name)}
		if response, err := client.Name(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Employee Details": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	if err := g.Run(":9090"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
