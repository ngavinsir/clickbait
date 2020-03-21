package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/clickbait/models"
)

func TestLabel(t *testing.T) {
	db, err := ConnectTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		ResetTestDB(db)
		db.Close()
	}()

	testRepository := initTestRepository(db)

	user, err := testRepository.CreateNewUser(context.Background(), &models.User{
		Username: "a",
		Password: "b",
	})
	if err != nil {
		t.Error(err)
	}

	article, err := testRepository.InsertArticle(context.Background(), "TEST_HEADLINE", "TEST_CONTENT")
	if err != nil {
		t.Error(err)
	}

	t.Run("Insert", testInsertLabel(testRepository, user.ID, article.ID))
}

func testInsertLabel(testRepository *testRepository, userID string, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		label, err := testRepository.InsertLabel(context.Background(), userID, articleID, "Clickbait", "clickbait")
		if err != nil {
			t.Error(err)
		}

		if label.ID == "" {
			t.Errorf("Want label id assigned, got %s", label.ID)
		}
		if got, want := label.UserID, userID; got != want {
			t.Errorf("Want label user id %s, got %s", want, got)
		}
		if got, want := label.ArticleID, articleID; got != want {
			t.Errorf("Want label article id %s, got %s", want, got)
		}
		if got, want := label.Value, "Clickbait"; got != want {
			t.Errorf("Want label value %s, got %s", want, got)
		}
		if got, want := label.Type, "clickbait"; got != want {
			t.Errorf("Want label type %s, got %s", want, got)
		}
	}
}
