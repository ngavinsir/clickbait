package model

import (
	"context"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles user data management.
type UserRepository interface {
	CreateNewUser(ctx context.Context, data *models.User) (*models.User, error)
	GetUser(ctx context.Context, username string) (*models.User, error)
}

// UserDatastore holds db information.
type UserDatastore struct {
	*DB
}

// CreateNewUser creates a new user with given username and password.
func (db *UserDatastore) CreateNewUser(ctx context.Context, data *models.User) (*models.User, error) {
	hash, _ := hashPassword(data.Password)
	user := &models.User{
		ID:       ksuid.New().String(),
		Username: data.Username,
		Password: hash,
	}

	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// GetUser returns user by given username.
func (db *UserDatastore) GetUser(ctx context.Context, username string) (*models.User, error) {
	return models.Users(models.UserWhere.Username.EQ(username)).One(ctx, db)
}
