package db

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/api/v1"
	userService "github.com/omaressameldin/bet-collector-services/grpc-services/user/pkg/api/v1"
	"github.com/omaressameldin/go-database-connector/app/pkg/database"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

const (
	MIN_DESCRIPTION_LENGTH int   = 10
	MIN_PAYMENT_LENGTH     int   = 2
	DEFAULT_ALL_LIMIT      int32 = 15
	DEFAULT_PAGE           int32 = 1
	USER_SERVICE                 = "user-service"
)

func validateDescription(description string) error {
	if len(description) <= MIN_DESCRIPTION_LENGTH {
		return fmt.Errorf("length should be at least %d", MIN_DESCRIPTION_LENGTH)
	}
	return nil
}

func validateExpiryDate(expiry *timestamp.Timestamp) error {
	t, err := ptypes.Timestamp(expiry)
	if err != nil {
		return err
	}
	if t.Before(time.Now()) {
		return fmt.Errorf("Expiry date has to be in the future")
	}
	return nil
}

func validatePayment(payment string) error {
	if len(payment) < MIN_PAYMENT_LENGTH {
		return fmt.Errorf("length should be at least %d", MIN_PAYMENT_LENGTH)
	}
	return nil
}

func validateUserId(userServiceUrl, userId string) error {
	client, err := grpc.Dial(userServiceUrl, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c := userService.NewUserServiceClient(client)
	req := userService.DoesUserExistRequest{
		Api: "v1",
		Id:  userId,
	}
	res, err := c.DoesUserExist(context.Background(), &req)
	if err != nil {
		return err
	} else if !res.Exists {
		return fmt.Errorf("User with id %s does not exist!", userId)
	} else {
		return nil
	}
}

func validateBetterIsNotAccepter(betterId, accepterId string) error {
	if betterId == accepterId {
		return fmt.Errorf("You can't bet yourself!")
	}

	return nil
}

func validateBet(
	userServiceUrl string,
	betterId string,
	description,
	payment,
	accepterId *string,
	winnerId *string,
	expiryDate *timestamp.Timestamp,
	completionDate *timestamp.Timestamp,
) []database.Validator {
	var descriptionError error
	if description != nil {
		descriptionError = validateDescription(*description)
	}

	var expiryError error
	if expiryDate != nil {
		expiryError = validateExpiryDate(expiryDate)
	}

	var paymentError error
	if payment != nil {
		paymentError = validatePayment(*payment)
	}

	var winnerError error
	if winnerId != nil {
		winnerError = validateUserId(userServiceUrl, *winnerId)

		if completionDate == nil {
			winnerError = fmt.Errorf("has to set completionDate with winner")
		}
	}

	var accepterError error
	var accepterBetterError error
	if accepterId != nil {
		accepterError = validateUserId(userServiceUrl, *accepterId)
		accepterBetterError = validateBetterIsNotAccepter(betterId, *accepterId)
	}

	betterError := validateUserId(userServiceUrl, betterId)

	return []database.Validator{
		database.CreateValidator("Description", descriptionError),
		database.CreateValidator("Payment", paymentError),
		database.CreateValidator("AccepterId", accepterError),
		database.CreateValidator("BetterId", betterError),
		database.CreateValidator("AccepterId", accepterBetterError),
		database.CreateValidator("BetterId", accepterBetterError),
		database.CreateValidator("ExpiryDate", expiryError),
		database.CreateValidator("winnerId", winnerError),
	}
}

func CreateBet(
	connector database.Connector,
	dependencies map[string]string,
	betterID string,
	bet *v1.Bet,
) error {
	bet.CreatedAt, _ = ptypes.TimestampProto(time.Now())
	bet.UpdatedAt = bet.CreatedAt
	key := xid.New().String()
	bet.Id = key
	bet.BetterId = betterID
	var winnerId *string

	if bet.ExpiryDate == nil {
		bet.ExpiryDate = &timestamp.Timestamp{}
	}

	if bet.WinnerId != nil {
		winnerId = &bet.WinnerId.Value
	}
	return connector.Create(
		validateBet(
			dependencies[USER_SERVICE],
			bet.BetterId,
			&bet.Description,
			&bet.Payment,
			&bet.AccepterId,
			winnerId,
			bet.ExpiryDate,
			bet.CompletionDate,
		),
		key,
		bet,
	)
}

func ReadBet(
	connector database.Connector,
	dependencies map[string]string,
	key string,
) (*v1.Bet, error) {
	var bet v1.Bet

	if err := connector.Read(key, &bet); err != nil {
		return nil, err
	}

	return &bet, nil
}

func ReadAllBets(
	connector database.Connector,
	dependencies map[string]string,
	userID string,
	limit int32,
	page int32,
) ([]*v1.Bet, error) {
	bets := make([]*v1.Bet, 0)
	getRefFn := func() interface{} { return &v1.Bet{} }
	appendFn := func(bet interface{}) {
		if bet, ok := bet.(*v1.Bet); ok {
			if bet.AccepterId == userID || bet.BetterId == userID {
				bets = append(bets, bet)
			}
		}
	}
	if err := connector.ReadAll(getRefFn, appendFn, []database.Filter{}); err != nil {
		return nil, err
	}
	pageLimit := DEFAULT_ALL_LIMIT
	if limit > 0 {
		pageLimit = limit
	}
	pageNumber := DEFAULT_PAGE
	if page > 0 {
		pageNumber = page
	}
	if int32(len(bets)) > (pageNumber-1)*pageLimit {
		return bets[(pageNumber-1)*pageLimit : int32(
			math.Min(float64(len(bets)),
				float64(pageNumber*pageLimit)),
		)], nil
	}
	return bets, nil
}

func getUpdated(bet *v1.BetUpdate) []database.Updated {
	updated := []database.Updated{}
	if bet.Description != nil {
		updated = append(updated, database.Updated{Key: "Description", Val: bet.Description.Value})
	}

	if bet.ExpiryDate != nil {
		updated = append(updated, database.Updated{Key: "ExpiryDate", Val: bet.ExpiryDate})
	}

	if bet.CompletionDate != nil {
		updated = append(updated, database.Updated{Key: "CompletionDate", Val: bet.CompletionDate})
	}

	if bet.Payment != nil {
		updated = append(updated, database.Updated{Key: "Payment", Val: bet.Payment.Value})
	}
	if bet.AccepterId != nil {
		updated = append(updated, database.Updated{Key: "AccepterId", Val: bet.AccepterId.Value})
	}
	if bet.WinnerId != nil {
		updated = append(updated, database.Updated{Key: "WinnerId", Val: bet.WinnerId})
	}
	updatedAt, _ := ptypes.TimestampProto(time.Now())
	updated = append(updated, database.Updated{Key: "UpdatedAt", Val: updatedAt})

	return updated
}

func UpdateBet(
	connector database.Connector,
	dependencies map[string]string,
	key string,
	userID string,
	bet *v1.BetUpdate,
) error {
	var description *string
	var payment *string
	var accepterId *string
	var winnerId *string

	if bet.Description != nil {
		description = &bet.Description.Value
	}

	if bet.Payment != nil {
		payment = &bet.Payment.Value
	}

	if bet.AccepterId != nil {
		accepterId = &bet.AccepterId.Value
	}

	if bet.WinnerId != nil {
		winnerId = &bet.WinnerId.Value
	}

	return connector.Update(
		validateBet(
			dependencies[USER_SERVICE],
			userID,
			description,
			payment,
			accepterId,
			winnerId,
			bet.ExpiryDate,
			bet.CompletionDate,
		),
		key,
		getUpdated(bet),
	)
}

func DeleteBet(
	connector database.Connector,
	dependencies map[string]string,
	key string,
) error {
	return connector.Delete(key)
}
