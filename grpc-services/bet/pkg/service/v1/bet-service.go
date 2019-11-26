package v1

import (
	"context"

	db "github.com/omaressameldin/bet-collector-services/grpc-services/bet/internal/db/v1"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/api/v1"
	"github.com/omaressameldin/go-database-connector/app/pkg/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

// BetServiceServer is implementation of v1.betServiceServer proto interface
type BetServiceServer struct {
	connector database.Connector
}

// NewBetServiceServer creates Bet service
func NewBetServiceServer(connector database.Connector) *BetServiceServer {
	return &BetServiceServer{
		connector: connector,
	}
}

// CloseConnection closes connection to DB
func (s *BetServiceServer) CloseConnection() error {
	return s.connector.CloseConnection()
}

// checkAPI checks if the API version requested by client is supported by server
func (s *BetServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new bet
func (s *BetServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	if err := db.CreateBet(s.connector, req.Bet); err != nil {
		return nil, err
	}
	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  req.Bet.Id,
	}, nil
}

// Read bet data
func (s *BetServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	bet, err := db.ReadBet(s.connector, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ReadResponse{
		Api: apiVersion,
		Bet: bet,
	}, nil
}

func (s *BetServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	bets, err := db.ReadAllBets(s.connector, req.Limit, req.Page, req.UserId)
	if err != nil {
		return nil, err
	}

	return &v1.ReadAllResponse{
		Api:  apiVersion,
		Bets: bets,
	}, nil
}

// Update bet
func (s *BetServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if err := db.UpdateBet(s.connector, req.Id, req.BetterId, req.BetUpdate); err != nil {
		return nil, err
	}

	return &v1.UpdateResponse{
		Api: apiVersion,
	}, nil
}

// Delete Bet
func (s *BetServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if err := db.DeleteBet(s.connector, req.Id); err != nil {
		return nil, err
	}

	return &v1.DeleteResponse{
		Api: apiVersion,
	}, nil
}
