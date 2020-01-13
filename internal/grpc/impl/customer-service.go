package impl

import (
	"context"
	"fmt"
	"github.com/marcsanmi/crm-integrator/proto"
)

type CustomerServiceGrpcImpl struct {}

func NewCustomerServiceGrpcImpl() *CustomerServiceGrpcImpl {
	return &CustomerServiceGrpcImpl{}
}

// Add function implementation of gRPC service
func (serviceImpl *CustomerServiceGrpcImpl) Add(ctx context.Context, request *proto.AddCustomersRequest) (*proto.AddCustomersResponse, error) {
	fmt.Println("Received request for adding Customer with id ")
	return nil, nil
}