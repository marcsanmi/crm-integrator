package main

import (
	"fmt"
	"log"
	"net"
	"github.com/marcsanmi/crm-integrator/internal/grpc/impl"
	"github.com/marcsanmi/crm-integrator/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Grpc in Action")

	netListener := getNetListener(7000)
	gRPCServer := grpc.NewServer()

	customerServiceImpl := impl.NewCustomerServiceGrpcImpl()
	proto.RegisterCustomerServiceServer(gRPCServer, customerServiceImpl)

	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
