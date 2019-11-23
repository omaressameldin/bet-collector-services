package db

import (
	"fmt"
	"math"
	"time"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/api/v1"
	"github.com/omaressameldin/go-database-connector/app/pkg/database"
	"github.com/rs/xid"
)

const MIN_DESCRIPTION_LENGTH int = 10
const MIN_PAYMENT_LENGTH int = 2
const DEFAULT_ALL_LIMIT int32 = 15
const DEFAULT_PAGE int32 = 1

func validateDescription(description string) error {
	if len(description) <= MIN_DESCRIPTION_LENGTH {
		return fmt.Errorf("length should be at least %d", MIN_DESCRIPTION_LENGTH)
	}
	return nil
}

func validatePayment(payment string) error {
	if len(payment) <= MIN_PAYMENT_LENGTH {
		return fmt.Errorf("length should be at least %d", MIN_PAYMENT_LENGTH)
	}
	return nil
}

func validateBet(description, payment *string) []database.Validator {
	var descriptionError error
	descriptionError = validateDescription(*description)

	var paymentError error
	if payment != nil {
		paymentError = validatePayment(*payment)
	}

	return []database.Validator{
		database.CreateValidator("Description", descriptionError),
		database.CreateValidator("Payment", paymentError),
	}
}

func CreateBet(connector database.Connector, bet *v1.Bet) error {
	bet.CreatedAt, _ = ptypes.TimestampProto(time.Now())
	key := xid.New().String()
	bet.Id = key

	bet.UpdatedAt = bet.CreatedAt

	return connector.Create(
		validateBet(&bet.Description, &bet.Payment),
		key,
		bet,
	)
}

func ReadBet(connector database.Connector, key string) (*v1.Bet, error) {
	var bet v1.Bet
	if err := connector.Read(key, &bet); err != nil {
		return nil, err
	}

	return &bet, nil
}

func ReadAllBets(
	connector database.Connector,
	limit int32,
	page int32,
	userId string,
) ([]*v1.Bet, error) {
	bets := make([]*v1.Bet, 0)
	getRefFn := func() interface{} { return &v1.Bet{} }
	appendFn := func(bet interface{}) {
		if bet, ok := bet.(*v1.Bet); ok {
			if bet.AccepterId == userId || bet.BetterId == userId {
				bets = append(bets, bet)
			}
		}
	}
	if err := connector.ReadAll(getRefFn, appendFn); err != nil {
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

	if bet.Payment != nil {
		updated = append(updated, database.Updated{Key: "Payment", Val: bet.Payment.Value})
	}
	updatedAt, _ := ptypes.TimestampProto(time.Now())
	updated = append(updated, database.Updated{Key: "UpdatedAt", Val: updatedAt})

	return updated
}

func UpdateBet(connector database.Connector, key string, bet *v1.BetUpdate) error {
	var description *string
	var payment *string

	if bet.Description != nil {
		description = &bet.Description.Value
	}

	if bet.Payment != nil {
		payment = &bet.Payment.Value
	}

	return connector.Update(
		validateBet(description, payment),
		key,
		getUpdated(bet),
	)
}

func DeleteBet(connector database.Connector, key string) error {
	if err := connector.Delete(key); err != nil {
		return err
	}

	return nil
}