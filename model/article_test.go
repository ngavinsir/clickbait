package model

import (
	"context"
	"testing"

	"github.com/ngavinsir/clickbait/models"
)

const (
	testHeadline = "TEST_HEADLINE"
	testContent = "TEST_CONTENT"
)

func TestArticle(t *testing.T) {
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

	t.Run("Insert", testInsertArticle(testRepository))
}

func testInsertArticle(testRepository *testRepository) func(t *testing.T) {
	return func(t *testing.T) {
		article, err := testRepository.InsertArticle(context.Background(), testHeadline, testContent)
		if err != nil {
			t.Error(err)
		}

		if article.ID == "" {
			t.Errorf("Want article id assigned, got %s", article.ID)
		}
		if article.Headline != testHeadline {
			t.Errorf("Want article headline %s, got %s", testHeadline, article.Headline)
		}
		if article.Content != testContent {
			t.Errorf("Want article content %s, got %s", testContent, article.Content)
		}

		t.Run("Get Random", testGetRandomArticle(testRepository, article.ID))
	}
}

func testGetRandomArticle(testRepository *testRepository, articleID string) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := testRepository.CreateNewUser(context.Background(), &models.User{
			Username: "a",
			Password: "b",
		})
		if err != nil {
			t.Error(err)
		}

		// Should return article because user hasn't labeled any articles.
		t.Run("1", func(t *testing.T) {
			article, err := testRepository.GetRandomArticle(context.Background(), user.ID, "clickbait")
			if err != nil {
				t.Error(err)
			}
			
			if article.Headline != testHeadline {
				t.Errorf("Want article headline %s, got %s", testHeadline, article.Headline)
			}
			if article.Content != testContent {
				t.Errorf("Want article content %s, got %s", testContent, article.Content)
			}
		})

		// Shouldn't get article because the user has labeled the article.
		t.Run("2", func(t *testing.T) {
			_, err = testRepository.InsertLabel(context.Background(), user.ID, articleID, "Clickbait", "clickbait")
			if err != nil {
				t.Error(err)
			}

			article, err := testRepository.GetRandomArticle(context.Background(), user.ID, "clickbait")
			if err == nil || err.Error() != "sql: no rows in result set" {
				t.Errorf("Want error, got %v", err)
			}

			if article != nil {
				t.Errorf("Want no article, got %#v", article)
			}
		})

		t.Run("3", func(t *testing.T) {
			article, err := testRepository.GetRandomArticle(context.Background(), user.ID, "summary")
			if err != nil {
				t.Error(err)
			}
			
			if article.Headline != testHeadline {
				t.Errorf("Want article headline %s, got %s", testHeadline, article.Headline)
			}
			if article.Content != testContent {
				t.Errorf("Want article content %s, got %s", testContent, article.Content)
			}
		})
	}
}
