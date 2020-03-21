package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/clickbait/models"
)

const (
	testUsername = "TEST_USERNAME"
	testPassword = "TEST_PASSWORD"
)

func TestUser(t *testing.T) {
	db, err := ConnectTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		ResetTestDB(db)
		db.Close()
	}()

	t.Run("Create", testCreateUser(&UserDatastore{DB: db}))
}

func testCreateUser(userRepository UserRepository) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := userRepository.CreateNewUser(context.Background(), &models.User{
			Username: testUsername,
			Password: testPassword,
		})
		if err != nil {
			t.Error(err)
		}

		if user.ID == "" {
			t.Errorf("Want user id assigned, got %s", user.ID)
		}
		if user.Password == testPassword {
			t.Errorf("Want user password hashed, got %s", user.Password)
		}
		if got, want := user.Username, testUsername; got != want {
			t.Errorf("Want user username %s, got %s", want, got)
		}

		t.Run("Get", testGetUser(userRepository))
	}
}

func testGetUser(userRepository UserRepository) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := userRepository.GetUser(context.Background(), testUsername)
		if err != nil {
			t.Error(err)
		}

		if got, want := user.Username, testUsername; got != want {
			t.Errorf("Want user username %s, got %s", want, got)
		}
	}
}
