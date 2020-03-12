package model

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/ksuid"
)

func TestInsertArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testID := ksuid.New().String()
	testHeadline := "test_article_value"
	testContent := "test_article_content"
	mock.ExpectExec(`INSERT INTO "articles"`).WithArgs(testID, testHeadline, testContent).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = insertArticleWithID(context.Background(), db, testID, testHeadline, testContent)
	if err != nil {
		t.Fatalf("error was not expected while inserting new article: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRandomarticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUserID := "test_user_id"
	ret := sqlmock.NewRows([]string{"id", "value"})
	ret.AddRow(driver.Value(testUserID), driver.Value("test_article_value_1"))
	ret.AddRow(driver.Value(testUserID), driver.Value("test_article_value_2"))

	mock.ExpectQuery(`select h.id, h.value`).WithArgs(testUserID, 3).WillReturnRows(ret)

	_, err = GetRandomArticle(context.Background(), db, testUserID, "")
	if err != nil {
		t.Fatalf("error was not expected while getting random article: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
