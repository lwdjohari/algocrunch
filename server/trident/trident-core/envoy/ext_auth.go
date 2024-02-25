package envoy

import (
	"context"
	"fmt"
	"net"

	auth_pb "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
)

type AuthServer struct{}

func (server *AuthServer) Check(
	ctx context.Context,
	request *auth_pb.CheckRequest,
) (*auth_pb.CheckResponse, error) {

	// block if path is /private
	path := request.Attributes.Request.Http.Path[1:]
	if path != "private" {
		return nil, fmt.Errorf("private request not allowed")
	}

	// allow all other requests
	return &auth_pb.CheckResponse{}, nil
}

func Auth() {
	// struct with check method
	type AuthServer struct{}

	endPoint := fmt.Sprintf("localhost:%d", 3001)
	listen, err := net.Listen("tcp", endPoint)

	grpcServer := grpc.NewServer()

	// register envoy proto server
	server := &AuthServer{}
	auth_pb.RegisterAuthorizationServer(grpcServer, server)

	grpcServer.Serve(listen)
}
