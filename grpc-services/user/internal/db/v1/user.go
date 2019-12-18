package db

import (
	"fmt"
	"net/url"
	"time"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/user/pkg/api/v1"
	"github.com/omaressameldin/go-database-connector/app/pkg/database"
	"github.com/rs/xid"
)

const MIN_NAME_LENGTH int = 3

func validateName(name string) error {
	if len(name) <= MIN_NAME_LENGTH {
		return fmt.Errorf("length should be at least %d", MIN_NAME_LENGTH)
	}
	return nil
}

func validateAvatar(avatar string) error {
	_, err := url.ParseRequestURI(avatar)
	return err
}

func validateUser(name, avatar *string) []database.Validator {
	var nameError error
	nameError = validateName(*name)

	var avatarError error
	if avatar != nil {
		avatarError = validateAvatar(*avatar)
	}

	return []database.Validator{
		database.CreateValidator("Name", nameError),
		database.CreateValidator("Avatar", avatarError),
	}
}

func CreateUser(connector database.Connector, user *v1.User) error {
	user.CreatedAt, _ = ptypes.TimestampProto(time.Now())
	user.UpdatedAt = user.CreatedAt
	key := xid.New().String()
	user.Id = key

	return connector.Create(
		validateUser(&user.Name, nil),
		key,
		user,
	)
}

func CreateUserFromAuthId(connector database.Connector, authId string) (*v1.User, error) {
	user, err := connector.Authenticate(authId)
	if err != nil {
		return nil, err
	}

	return &v1.User{
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}
func ReadUser(connector database.Connector, key string) (*v1.User, error) {
	var user v1.User
	if err := connector.Read(key, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func getUpdated(user *v1.UserUpdate) []database.Updated {
	updated := []database.Updated{}
	if user.Name != nil {
		updated = append(updated, database.Updated{Key: "Name", Val: user.Name.Value})
	}

	if user.Avatar != nil {
		updated = append(updated, database.Updated{Key: "Avatar", Val: user.Avatar.Value})
	}
	updatedAt, _ := ptypes.TimestampProto(time.Now())
	updated = append(updated, database.Updated{Key: "UpdatedAt", Val: updatedAt})

	return updated
}

func UpdateUser(connector database.Connector, key string, user *v1.UserUpdate) error {
	var name *string
	var avatar *string

	if user.Name != nil {
		name = &user.Name.Value
	}

	if user.Avatar != nil {
		avatar = &user.Avatar.Value
	}

	return connector.Update(
		validateUser(name, avatar),
		key,
		getUpdated(user),
	)
}

func DeleteUser(connector database.Connector, key string) error {
	if err := connector.Delete(key); err != nil {
		return err
	}

	return nil
}

func DoesUserExist(connector database.Connector, id string) bool {
	_, err := ReadUser(connector, id)

	return err == nil
}
