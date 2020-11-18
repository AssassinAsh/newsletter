package server

import (
	"context"

	"github.com/AssassinAsh/newsletter/pkg/proto"
	"github.com/AssassinAsh/newsletter/pkg/services"

	"google.golang.org/grpc"
)

//Server represents all server handlers.
type server struct {
	authenticationService services.IAuthenticationService
}

//RegisterServer new api server.
func RegisterServer(srv *grpc.Server) {
	proto.RegisterGUMIServer(srv, &server{
		authenticationService: services.NewUserService(),
	})
}

//GetByID Authenticates the UserId.
func (sv *server) GetByID(ctx context.Context, req *proto.GetByIDRequest) (*proto.User, error) {
	res, err := sv.authenticationService.GetByID(req)
	if err != nil {
		return nil, err
	}
	return res, err
}
