package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"

	"user-service/app/config/initializers"
	"user-service/app/helpers"
	"user-service/app/proto"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

type myGRPCServer struct {
	// type embedded to comply with Google lib
	proto.UnimplementedInvoicerServer
}

func (m *myGRPCServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Println("Create called")
	log.Println("Create called")
	log.Println("Create called")
	return &proto.CreateResponse{Pdf: []byte("TODO")}, nil
}

func main() {
	// dataBase := initializers.ConnectDb()

	// defer dataBase.Close()

	// router := gin.Default()

	// log.Infof("FUUUUUUCK")

	// api.Controllers(router, dataBase)

	// err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	// if err != nil {
	// 	log.Panicf("Server listen err: %v", err)
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", SERVER_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myInvoicerServer := &myGRPCServer{}
	proto.RegisterInvoicerServer(s, myInvoicerServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}
