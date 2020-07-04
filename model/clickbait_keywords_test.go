package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/clickbait/models"
)

func TestClickbaitKeywords(t *testing.T) {
	db, err := ConnectTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		ResetTestDB(db)
		db.Close()
	}()
	ResetTestDB(db)

	testRepository := initTestRepository(db)

	user, err := testRepository.CreateNewUser(context.Background(), &models.User{
		Email:    "a",
		Password: "b",
	})
	if err != nil {
		t.Error(err)
	}

	article, err := testRepository.InsertArticle(
		context.Background(),
		"",
		"",
		"",
		"",
		"",
	)
	if err != nil {
		t.Error(err)
	}

	label, err := testRepository.InsertLabel(
		context.Background(),
		user.ID,
		article.ID,
		"",
		"",
	)

	t.Run("Add", func(t *testing.T) {
		keywords := []string{"keyword1", "keyword2"}

		err := testRepository.AddClickbaitKeywords(
			context.Background(),
			label.ID,
			keywords,
		)
		if err != nil {
			t.Error(err)
		}

		err = label.L.LoadClickbaitKeywords(
			context.Background(),
			testRepository.LabelDatastore,
			true,
			label,
			nil,
		)
		if err != nil {
			t.Error(err)
		}

		if got, want := len(label.R.ClickbaitKeywords), len(keywords); got != want {
			t.Errorf("got clickbait keywords length %d, want %d", got, want)
		}

		for _, keyword := range keywords {
			for _, clickbaitKeyword := range label.R.ClickbaitKeywords {
				if keyword == clickbaitKeyword.Keyword {
					return
				}
			}
			t.Errorf("want keyword %s to exist, got nil", keyword)
		}
	})
}
