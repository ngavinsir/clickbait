package model

import (
	"context"
	"testing"
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
	
	t.Run("Insert", testInsertArticle(&ArticleDatastore{DB: db}))
}

func testInsertArticle(articleRepository ArticleRepository) func (t *testing.T) {
	return func(t *testing.T) {
		article, err := articleRepository.InsertArticle(context.Background(), testHeadline, testContent)
		if err != nil {
			t.Error(err)
		}

		if article.Headline != testHeadline {
			t.Errorf("Want article headline %s, got %s", testHeadline, article.Headline)
		}
		if article.Content != testContent {
			t.Errorf("Want article content %s, got %s", testContent, article.Content)
		}

		t.Run("Get Random", testGetRandomArticle(articleRepository))
	}
}

func testGetRandomArticle(articleRepository ArticleRepository) func(t *testing.T) {
	return func(t *testing.T) {
		article, err := articleRepository.GetRandomArticle(context.Background(), "a", "a")
		if err != nil {
			t.Error(err)
		}
		
		if article.Headline != testHeadline {
			t.Errorf("Want article headline %s, got %s", testHeadline, article.Headline)
		}
		if article.Content != testContent {
			t.Errorf("Want article content %s, got %s", testContent, article.Content)
		}
	}
}
