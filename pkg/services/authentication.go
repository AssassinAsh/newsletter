package services

import "github.com/AssassinAsh/newsletter/pkg/proto"

//IAuthenticationService wraps all authentication methods.
type IAuthenticationService interface {
	GetByID(req *proto.GetByIDRequest) (*proto.User, error)
}

//AuthenticationService represents common authentication receiver.
type AuthenticationService struct {
}

//NewUserService creates authentication service instance.
func NewUserService() IAuthenticationService {
	return &AuthenticationService{}
}

//GetByID authenticates the user.
func (auth *AuthenticationService) GetByID(req *proto.GetByIDRequest) (*proto.User, error) {
	user := proto.User{
		Name:  "Ashvin",
		Email: "ashvinrokade@gmail.com",
		Id:    req.Id,
	}

	return &user, nil
}
