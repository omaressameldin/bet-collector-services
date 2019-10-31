package v1

import (
	"context"

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

// Create new user
func (s *UserServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if err := db.CreateUser(s.connector, req.User.AuthId, req.User); err != nil {
		return nil, status.Error(codes.Unknown, "failed to inssert into User->"+err.Error())
	}
	return &v1.CreateResponse{
		Api: apiVersion,
	}, nil
}

// Read user data
func (s *UserServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	user, err := db.ReadUser(s.connector, req.AuthId)
	if err != nil {
		return nil, err
	}
	return &v1.ReadResponse{
		User: user,
	}, nil
}

// Update User
func (s *UserServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if err := db.UpdateUser(s.connector, req.AuthId, req.User); err != nil {
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
	if err := db.DeleteUser(s.connector, req.AuthId); err != nil {
		return nil, err
	}
	return &v1.DeleteResponse{
		Api: apiVersion,
	}, nil
}
