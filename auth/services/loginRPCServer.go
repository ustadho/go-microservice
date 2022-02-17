package services

import (
	rpc_auth "github.com/ustadho/go-microservice/rpc/rp_auth"
)

type LoginRPCServer struct {
	rpc_auth.UnimplementedLoginServiceServer
}
