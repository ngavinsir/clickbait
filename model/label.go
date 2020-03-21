package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// LabelRepository handles label data management.
type LabelRepository interface {
	InsertLabel(ctx context.Context, userID string, articleID string, value string, labelType string)	(*models.Label, error)
	IsLabeledByUser(ctx context.Context, articleID string, userID string, labelType string)				(bool, error)
	DeleteLabel(ctx context.Context, labelID string)													error
	GetArticleLabel(ctx context.Context, userID string, labelType string)								([]*ArticleLabel, error)
	GetArticleLabelCount(ctx context.Context, articleID string, labelType string)						(int64, error)
}

// LabelDatastore holds db information.
type LabelDatastore struct {
	*DB
}

// ArticleLabel contains label and article
type ArticleLabel struct {
	models.Label   `boil:",bind" json:"label"`
	models.Article `boil:",bind" json:"article"`
}

// InsertLabel with maximum of 3 labels per headline
func (db *LabelDatastore) InsertLabel(ctx context.Context, userID string, articleID string, value string, labelType string) (*models.Label, error) {
	isLabeled, err := db.IsLabeledByUser(ctx, articleID, userID, labelType)
	if err != nil {
		return nil, err
	}
	if isLabeled {
		return nil, errors.New("can't label the same article")
	}

	articleLabelCount, err := db.GetArticleLabelCount(ctx, articleID, labelType)
	if err != nil {
		return nil, err
	}

	if articleLabelCount >= 3 {
		return nil, errors.New("maximum label reached")
	}

	label := &models.Label{
		ID:        ksuid.New().String(),
		UserID:    userID,
		ArticleID: articleID,
		Value:     value,
		Type: 	   labelType,
	}
	if err := label.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return label, nil
}

// IsLabeledByUser return true if given article has been labeled by a user.
func (db *LabelDatastore) IsLabeledByUser(ctx context.Context, articleID string, userID string, labelType string) (bool, error) {
	isLabeled, err := models.Labels(
		models.LabelWhere.ArticleID.EQ(articleID),
		models.LabelWhere.UserID.EQ(userID),
		models.LabelWhere.Type.EQ(labelType),
	).Exists(ctx, db)
	if err != nil {
		return false, err
	}

	return isLabeled, nil
}

// DeleteLabel deletes label by label id.
func (db *LabelDatastore) DeleteLabel(ctx context.Context, labelID string) error {
	_, err := models.Labels(
		models.LabelWhere.ID.EQ(labelID),
	).DeleteAll(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

// GetArticleLabel return all label with the same type by user_id with the article value
func (db *LabelDatastore) GetArticleLabel(ctx context.Context, userID string, labelType string) ([]*ArticleLabel, error) {
	data := []*ArticleLabel{}
	err := queries.Raw(`
		select	l.id as "label.id", articles.id as "article.id", articles.headline as "article.headline",
				articles.content as "article.content", l.value as "label.value",
				l.updated_at as "label.updated_at"
		from labels as l
		inner join articles on l.article_id = articles.id
		where l.user_id = $1 and l.type = $2	
	`, userID, labelType).Bind(ctx, db, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetArticleLabelCount returns label count by headline id.
func (db *LabelDatastore) GetArticleLabelCount(ctx context.Context, articleID string, labelType string) (int64, error) {
	labelCount, err := models.Labels(
			models.LabelWhere.ArticleID.EQ(articleID),
			models.LabelWhere.Type.EQ(labelType),
	).Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return labelCount, nil
}
