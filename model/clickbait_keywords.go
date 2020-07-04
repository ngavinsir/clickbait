package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// ClickbaitKeywordRepository handles clickbait keywords data management.
type ClickbaitKeywordRepository interface {
	AddClickbaitKeywords(ctx context.Context, labelID string, keywords []string) error
}

// CLickbaitKeywordDatastore holds db information.
type CLickbaitKeywordDatastore struct {
	*DB
}

// AddClickbaitKeywords adds clickbait keywords to given label
func (db *CLickbaitKeywordDatastore) AddClickbaitKeywords(ctx context.Context, labelID string, keywords []string) error {
	labelExists, err := models.LabelExists(ctx, db, labelID)
	if err != nil {
		return err
	}

	if !labelExists {
		return errors.New("can't find label")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, keyword := range keywords {
		clickbaitKeyword := &models.ClickbaitKeyword{
			ID:      ksuid.New().String(),
			LabelID: labelID,
			Keyword: keyword,
		}

		err = clickbaitKeyword.Insert(ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
