package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// LabelRepository handles label data management.
type LabelRepository interface {
	InsertLabel(ctx context.Context, userID string, articleID string, value string, labelType string) (*models.Label, error)
	IsLabeledByUser(ctx context.Context, articleID string, userID string, labelType string) (bool, error)
	DeleteLabel(ctx context.Context, labelID string) error
	GetLabel(ctx context.Context, userID string, labelType string) ([]*models.Label, error)
	GetArticleLabelCount(ctx context.Context, articleID string, labelType string) (int64, error)
	GetLabelLeaderboard(ctx context.Context, labelType string, limit uint8) (*[]LabelScore, error)
}

// LabelDatastore holds db information.
type LabelDatastore struct {
	*DB
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
		Type:      labelType,
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

// GetLabel return all label with the same type by user_id with the article value
func (db *LabelDatastore) GetLabel(ctx context.Context, userID string, labelType string) ([]*models.Label, error) {
	labels, err := models.Labels(
		Load(models.LabelRels.Article),
		Load(models.LabelRels.ClickbaitKeywords),
		models.LabelWhere.UserID.EQ(userID),
		models.LabelWhere.Type.EQ(labelType),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	return labels, nil
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

// GetLabelLeaderboard returns top labeling score.
func (db *LabelDatastore) GetLabelLeaderboard(ctx context.Context, labelType string, limit uint8) (*[]LabelScore, error) {
	data := &[]LabelScore{}
	if err := queries.Raw(`
		select u.name, count(l.id) as score
		from labels as l
		inner join users u on l.user_id = u.id
		where l.type = $1
		group by u.name
		order by count(l.id) desc
		limit $2;
	`, labelType, limit).Bind(ctx, db, data); err != nil {
		return nil, err
	}

	return data, nil
}

// LabelScore holds label leaderboard information.
type LabelScore struct {
	Name  string `json:"name"`
	Score uint32 `json:"score"`
}
