package model

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/segmentio/ksuid"
)

func TestInsertHeadline(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testID := ksuid.New().String()
	testValue := "test_headline_value"
	mock.ExpectExec(`INSERT INTO "headlines"`).WithArgs(testID, testValue).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = insertHeadlineWithID(context.Background(), db, testID, testValue)
	if err != nil {
		t.Fatalf("error was not expected while inserting new headline: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRandomHeadline(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUserID := "test_user_id"
	ret := sqlmock.NewRows([]string{"id", "value"})
	ret.AddRow(driver.Value(testUserID), driver.Value("test_headline_value_1"))
	ret.AddRow(driver.Value(testUserID), driver.Value("test_headline_value_2"))

	mock.ExpectQuery(`select h.id, h.value`).WithArgs(testUserID, 3).WillReturnRows(ret)

	_, err = GetRandomHeadline(context.Background(), db, testUserID)
	if err != nil {
		t.Fatalf("error was not expected while getting random headline: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}