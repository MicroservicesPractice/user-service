package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpcApi "user-service/app/api/grpc/user"
	httpApi "user-service/app/api/http"
	"user-service/app/config/initializers"
	"user-service/app/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")
var GRPC_SERVER_PORT = helpers.GetEnv("GRPC_SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	startGRPCServer()
	startHTTPServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", GRPC_SERVER_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userServer := &grpcApi.UserGRPCServer{}
	grpcApi.RegisterUserServer(grpcServer, userServer)

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Infof("Server has been started on port %v", GRPC_SERVER_PORT)
	}()
}

func startHTTPServer() {
	dataBase := initializers.ConnectDb()

	defer dataBase.Close()

	router := gin.Default()

	httpApi.Controllers(router, dataBase)

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}
}
