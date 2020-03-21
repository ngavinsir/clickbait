package model

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/models"
)

type testRepository struct {
	UserRepository
	LabelRepository
	ArticleRepository
}

// ConnectTestDB connects to test db.
func ConnectTestDB() (*DB, error) {
	conn := "dbname=clickbait_test host=localhost user=postgres password=postgres"
	return Connect(conn)
}

// ResetTestDB clears test db data.
func ResetTestDB(db *DB) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	models.Users().DeleteAll(context.Background(), tx)
	models.Articles().DeleteAll(context.Background(), tx)
	models.Labels().DeleteAll(context.Background(), tx)

	err = tx.Commit()
	return err
}

func initTestRepository(db *DB) *testRepository {
	return &testRepository{
		UserRepository: &UserDatastore{DB: db},
		ArticleRepository: &ArticleDatastore{DB: db},
		LabelRepository: &LabelDatastore{DB: db},
	}
}