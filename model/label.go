package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// InsertLabel with maximum of 3 labels per headline
func InsertLabel(ctx context.Context, exec boil.ContextExecutor,
	userID string, articleID string, value string, labelType string) (*models.Label, error) {
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

	label := &models.Label{
		ID:        ksuid.New().String(),
		UserID:    userID,
		ArticleID: articleID,
		Value:     value,
		Type: 	   labelType,
	}
	if err := label.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}

	return label, nil
}

func isLabeledByUser(ctx context.Context, exec boil.ContextExecutor, articleID string, userID string, labelType string) (bool, error) {
	isLabeled, err := models.Labels(
		models.LabelWhere.ArticleID.EQ(articleID),
		models.LabelWhere.UserID.EQ(userID),
		models.LabelWhere.Type.EQ(labelType),
	).Exists(ctx, exec)
	if err != nil {
		return false, err
	}

	return isLabeled, nil
}

// DeleteLabel deletes label by label id.
func DeleteLabel(ctx context.Context, exec boil.ContextExecutor, labelID string) error {
	_, err := models.Labels(
		models.LabelWhere.ID.EQ(labelID),
	).DeleteAll(ctx, exec)
	if err != nil {
		return err
	}

	return nil
}

// GetArticleLabel return all label with the same type by user_id with the article value
func GetArticleLabel(ctx context.Context, exec boil.ContextExecutor, userID string, labelType string) ([]*ArticleLabel, error) {
	data := []*ArticleLabel{}
	err := queries.Raw(`
		select	l.id as "label.id", articles.id as "article.id", articles.headline as "article.headline",
				articles.content as "article.content", l.value as "label.value",
				l.updated_at as "label.updated_at"
		from labels as l
		inner join articles on l.article_id = articles.id
		where l.user_id = $1 and l.type = $2	
	`, userID, labelType).Bind(ctx, exec, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetArticleLabelCount returns label count by headline id.
func GetArticleLabelCount(ctx context.Context, exec boil.ContextExecutor, articleID string, labelType string) (int64, error) {
	labelCount, err := models.Labels(
			models.LabelWhere.ArticleID.EQ(articleID),
			models.LabelWhere.Type.EQ(labelType),
	).Count(ctx, exec)
	if err != nil {
		return 0, err
	}

	return labelCount, nil
}

// ArticleLabel contains label and article
type ArticleLabel struct {
	models.Label   `boil:",bind" json:"label"`
	models.Article `boil:",bind" json:"article"`
}
