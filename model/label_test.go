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
	ResetTestDB(db)

	testRepository := initTestRepository(db)

	user, err := testRepository.CreateNewUser(context.Background(), &models.User{
		Email: "a",
		Password: "b",
	})
	if err != nil {
		t.Error(err)
	}

	article, err := testRepository.InsertArticle(context.Background(), "TEST_HEADLINE", "TEST_CONTENT")
	if err != nil {
		t.Error(err)
	}

	t.Run("Insert", testInsertLabel(testRepository, user, article))
}

func testInsertLabel(testRepository *testRepository, user *models.User, article *models.Article) func(t *testing.T) {
	return func(t *testing.T) {
		label, err := testRepository.InsertLabel(context.Background(), user.ID, article.ID, "Clickbait", "clickbait")
		if err != nil {
			t.Error(err)
		}

		if label.ID == "" {
			t.Errorf("Want label id assigned, got %s", label.ID)
		}
		if got, want := label.UserID, user.ID; got != want {
			t.Errorf("Want label user id %s, got %s", want, got)
		}
		if got, want := label.ArticleID, article.ID; got != want {
			t.Errorf("Want label article id %s, got %s", want, got)
		}
		if got, want := label.Value, "Clickbait"; got != want {
			t.Errorf("Want label value %s, got %s", want, got)
		}
		if got, want := label.Type, "clickbait"; got != want {
			t.Errorf("Want label type %s, got %s", want, got)
		}

		t.Run("Is Labeled", testIsLabeled(testRepository, user.ID, article.ID))
		t.Run("Get", testGetLabel(testRepository, user.ID, article.ID, label.ID))
		t.Run("Get Count", testGetLabelCount(testRepository, article.ID))
		t.Run("Delete", testDeleteLabel(testRepository, label.ID, article.ID))
	}
}

func testIsLabeled(testRepository *testRepository, userID string, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		isLabeled, err := testRepository.IsLabeledByUser(context.Background(), articleID, userID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if !isLabeled {
			t.Errorf("Want is labeled to be true, got %v", isLabeled)
		}
	}
}

func testGetLabel(testRepository *testRepository, userID string, articleID string, labelID string) func(t *testing.T) {
	return func(t *testing.T) {
		articleLabel, err := testRepository.GetArticleLabel(context.Background(), userID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := len(articleLabel), 1; got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}
		if got, want := articleLabel[0].Label.ID, labelID; got != want {
			t.Errorf("Want label id %s, got %s", want, got)
		}
		if got, want := articleLabel[0].Article.ID, articleID; got != want {
			t.Errorf("Want article id %s, got %s", want, got)
		}

		articleLabel, err = testRepository.GetArticleLabel(context.Background(), userID, "summary")
		if err != nil {
			t.Error(err)
		}

		if got, want := len(articleLabel), 0; got != want {
			t.Errorf("Want get label return nothing, got %+v", articleLabel)
		}
	}
}

func testGetLabelCount(testRepository *testRepository, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		labelCount, err := testRepository.GetArticleLabelCount(context.Background(), articleID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(1); got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}

		labelCount, err = testRepository.GetArticleLabelCount(context.Background(), articleID, "summary")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(0); got != want {
			t.Errorf("Want get label count %d, got %d", want, labelCount)
		}
	}
}

func testDeleteLabel(testRepository *testRepository, labelID string, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		err := testRepository.DeleteLabel(context.Background(), labelID)
		if err != nil {
			t.Error(err)
		}

		labelCount, err := testRepository.GetArticleLabelCount(context.Background(), articleID, "clickbait")
		if err != nil {
			t.Error(err)
		}

		if got, want := labelCount, int64(0); got != want {
			t.Errorf("Want get label count %d, got %d", want, got)
		}
	}
}
