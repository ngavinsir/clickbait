package model

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// InsertLabel with maximum of 3 labels per headline
func InsertLabel(ctx context.Context, exec boil.ContextExecutor,
	userID string, articleID string, value string, labelType string) (*Label, error) {
	isLabeled, err := isLabeledByUser(ctx, exec, articleID, userID, labelType)
	if err != nil {
		return nil, err
	}
	if isLabeled {
		return nil, errors.New("can't label the same article")
	}

	articleLabelCount, err := GetArticleLabelCount(ctx, exec, articleID, labelType)
	if err != nil {
		return nil, err
	}

	if articleLabelCount >= 3 {
		return nil, errors.New("maximum label reached")
	}

	label := &Label{
		ID:        ksuid.New().String(),
		UserID:    userID,
		ArticleID: articleID,
		Value:     value,
	}

	if err := insertRawLabel(ctx, exec, label, labelType); err != nil {
		return nil, err
	}

	return label, nil
}

func insertRawLabel(ctx context.Context, exec boil.ContextExecutor, label *Label, labelType string) error {
	if labelType == "clickbait" {
		rawLabel := &models.ClickbaitLabel{
			ID:        label.ID,
			UserID:    label.UserID,
			ArticleID: label.ArticleID,
			Value:     label.Value,
		}
		if err := rawLabel.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}

		return nil
	} else if labelType == "summary" {
		rawLabel := &models.SummaryLabel{
			ID:        label.ID,
			UserID:    label.UserID,
			ArticleID: label.ArticleID,
			Value:     label.Value,
		}
		if err := rawLabel.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("invalid label type")
	}
}

func isLabeledByUser(ctx context.Context, exec boil.ContextExecutor, articleID string, userID string, labelType string) (bool, error) {
	if labelType == "clickbait" {
		isLabeled, err := models.ClickbaitLabels(
			models.ClickbaitLabelWhere.ArticleID.EQ(articleID),
			models.ClickbaitLabelWhere.UserID.EQ(userID),
		).Exists(ctx, exec)
		if err != nil {
			return false, err
		}

		return isLabeled, nil
	} else if labelType == "summary" {
		isLabeled, err := models.SummaryLabels(
			models.SummaryLabelWhere.ArticleID.EQ(articleID),
			models.SummaryLabelWhere.UserID.EQ(userID),
		).Exists(ctx, exec)
		if err != nil {
			return false, err
		}

		return isLabeled, nil
	} else {
		return false, errors.New("invalid label type")
	}
}

// DeleteLabel helper
func DeleteLabel(ctx context.Context, exec boil.ContextExecutor, labelID string, labelType string) error {
	if labelType == "clickbait" {
		_, err := models.ClickbaitLabels(
			models.ClickbaitLabelWhere.ID.EQ(labelID),
		).DeleteAll(ctx, exec)
		if err != nil {
			return err
		}

		return nil
	} else if labelType == "summary" {
		_, err := models.SummaryLabels(
			models.SummaryLabelWhere.ID.EQ(labelID),
		).DeleteAll(ctx, exec)
		if err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("invalid label type")
	}
}

// GetArticleLabel return all label with the same type by user_id with the article value
func GetArticleLabel(ctx context.Context, exec boil.ContextExecutor, userID string, labelType string) ([]*ArticleLabel, error) {
	data := []*ArticleLabel{}
	err := queries.Raw(
		fmt.Sprintf(`
			select	l.id as "label.id", articles.id as "article.id", articles.headline as "article.headline",
					articles.content as "article.content", l.value as "label.value",
					l.updated_at as "label.updated_at"
			from %s_labels as l
			inner join articles on l.article_id = articles.id
			where l.user_id = $1	
		`, labelType),
		userID).Bind(ctx, exec, &data)
	if err != nil {
		return nil, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Label.UpdatedAt.After(data[j].Label.UpdatedAt)
	})

	return data, nil
}

// GetArticleLabelCount return label count by headline id
func GetArticleLabelCount(ctx context.Context, exec boil.ContextExecutor, articleID string, labelType string) (int64, error) {
	if labelType == "clickbait" {
		labelCount, err := models.ClickbaitLabels(
			models.ClickbaitLabelWhere.ArticleID.EQ(articleID),
		).Count(ctx, exec)
		if err != nil {
			return 0, err
		}

		return labelCount, nil
	} else if labelType == "summary" {
		labelCount, err := models.SummaryLabels(
			models.SummaryLabelWhere.ArticleID.EQ(articleID),
		).Count(ctx, exec)
		if err != nil {
			return 0, err
		}

		return labelCount, nil
	} else {
		return 0, errors.New("invalid label type")
	}
}

// Label general struct
type Label struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    string    `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id"`
	ArticleID string    `boil:"article_id" json:"article_id,omitempty" toml:"article_id" yaml:"article_id"`
	Value     string    `boil:"value" json:"value" toml:"value" yaml:"value"`
	CreatedAt time.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

// ArticleLabel contains label_id, article_id, article_headline, article_content, label_value, label_updated_at
type ArticleLabel struct {
	Label `boil:",bind" json:"label"`
	models.Article `boil:",bind" json:"article"`
}
