package v1

import (
	"context"
	"fmt"

	db "github.com/omaressameldin/bet-collector-services/grpc-services/user/internal/db/v1"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/user/pkg/api/v1"
	"github.com/omaressameldin/go-database-connector/app/pkg/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// UserServiceServer is implementation of v1.userServiceServer proto interface
type UserServiceServer struct {
	connector database.Connector
}

// NewUserServiceServer creates User service
func NewUserServiceServer(connector database.Connector) *UserServiceServer {
	return &UserServiceServer{
		connector: connector,
	}
}

// CloseConnection closes connection to DB
func (s *UserServiceServer) CloseConnection() error {
	return s.connector.CloseConnection()
}

// checkAPI checks if the API version requested by client is supported by server
func (s *UserServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *UserServiceServer) isAuthorizedtoEditUser(token string, id string) bool {
	authUser, err := s.connector.Authenticate(token)
	if err != nil {
		return false
	}
	return authUser.ID == id
}

// Login user login through google sign in
func (s *UserServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	user, err := db.CreateUserFromAuthId(s.connector, req.AuthId)
	if err != nil {
		return nil, err
	}

	users, err := db.FindUsersBy(s.connector, db.Filters{Email: &user.Email})
	if len(users) > 0 {
		user = users[0]
	} else {
		err = db.CreateUser(s.connector, user)
		if err != nil {
			return nil, err
		}
	}

	return &v1.LoginResponse{
		Api:  req.Api,
		User: user,
	}, nil
}

// Read user data
func (s *UserServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	user, err := db.ReadUser(s.connector, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ReadResponse{
		Api:  apiVersion,
		User: user,
	}, nil
}

func (s *UserServiceServer) DoesUserExist(
	ctx context.Context,
	req *v1.DoesUserExistRequest,
) (*v1.DoesUserExistResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	users, err := db.FindUsersBy(s.connector, db.Filters{ID: &req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.DoesUserExistResponse{
		Api:    apiVersion,
		Exists: len(users) > 0,
	}, nil
}

// Update User
func (s *UserServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if !s.isAuthorizedtoEditUser(req.Token, req.Id) {
		return nil, fmt.Errorf("UnAuthorized to change user")
	}

	if err := db.UpdateUser(s.connector, req.Id, req.User); err != nil {
		return nil, err
	}
	return &v1.UpdateResponse{
		Api: apiVersion,
	}, nil
}

// Delete User
func (s *UserServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if !s.isAuthorizedtoEditUser(req.Token, req.Id) {
		return nil, fmt.Errorf("UnAuthorized to change user")
	}

	if err := db.DeleteUser(s.connector, req.Id); err != nil {
		return nil, err
	}
	return &v1.DeleteResponse{
		Api: apiVersion,
	}, nil
}
