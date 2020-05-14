package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles user data management.
type UserRepository interface {
	CreateNewUser(ctx context.Context, data *models.User) (*models.User, error)
	GetUser(ctx context.Context, email string) (*models.User, error)
	GetUserbyID(ctx context.Context, ID string) (*models.User, error)
}

// UserDatastore holds db information.
type UserDatastore struct {
	*DB
}

// CreateNewUser creates a new user with given user details.
func (db *UserDatastore) CreateNewUser(ctx context.Context, user *models.User) (*models.User, error) {
	emailUsed, err := models.Users(models.UserWhere.Email.EQ(user.Email)).Exists(ctx, db)
	if err != nil {
		return nil, err
	}
	if emailUsed {
		return nil, errors.New("email has already been used")
	}

	hash, _ := hashPassword(user.Password)
	user.ID = ksuid.New().String()
	user.Password = hash

	if err = user.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// GetUser returns user by given email.
func (db *UserDatastore) GetUser(ctx context.Context, email string) (*models.User, error) {
	return models.Users(models.UserWhere.Email.EQ(email)).One(ctx, db)
}

// GetUserbyID returns user by given user id.
func (db *UserDatastore) GetUserbyID(ctx context.Context, ID string) (*models.User, error) {
	return models.Users(models.UserWhere.ID.EQ(ID)).One(ctx, db)
}
