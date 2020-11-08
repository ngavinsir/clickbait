package model

import (
	"context"
	"errors"
	"time"

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
	GetLabelScore(ctx context.Context, labelType, userID string) (int, error)
	GetLabelProgress(ctx context.Context, labelType string, startWeek time.Time, duration int) ([][]*LabelProgress, error)
	GetLabelCount(ctx context.Context, userID string) (int64, error)
}

// LabelProgress tracks each user labeling progress
type LabelProgress struct {
	Name     string `json:"name"`
	Progress int64	`json:"progress"`
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

	thisWeekLabelCount, err := db.GetLabelCount(ctx, userID)
	if err != nil {
		return nil, err
	}
	if thisWeekLabelCount >= 200 {
		return nil, errors.New("you have reached your weekly label count limit")
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

// GetLabelScore returns user's label score.
func (db *LabelDatastore) GetLabelScore(ctx context.Context, labelType, userID string) (int, error) {
	score, err := models.Labels(
		models.LabelWhere.Type.EQ(labelType),
		models.LabelWhere.UserID.EQ(userID),
	).Count(ctx, db)
	if err != nil {
		return 0, err
	}

	return int(score), nil
}

// GetLabelProgress returns user's labeling progress
func (db *LabelDatastore) GetLabelProgress(
	ctx context.Context, labelType string, startWeek time.Time, duration int,
) ([][]*LabelProgress, error) {

	type LabelProgressWeek struct {
		Name     string    `json:"name"`
		Progress int64     `json:"progress"`
		Week     time.Time `json:"week"`
	}

	startWeek = startWeek.Truncate(7 * 24 * time.Hour).UTC()
	endWeek := startWeek.Add(time.Duration(duration-1) * 7 * 24 * time.Hour)

	labelProgressWeeks := &[]LabelProgressWeek{}
	if err := queries.Raw(`
		with weeks as (
			select generate_series($1::timestamp, $2::timestamp, '7 days') week
		)
		select w.week week, u.name, count(l.id) progress
		from weeks w
			cross join users u
			left join labels l on date_trunc('week', l.created_at) = w.week and l.user_id = u.id
		group by week, u.name
		order by week;
	`, startWeek.Format(time.RFC3339), endWeek.Format(time.RFC3339)).Bind(ctx, db, labelProgressWeeks); err != nil {
		return nil, err
	}

	result := make([][]*LabelProgress, duration)
	currentWeek := (*labelProgressWeeks)[0].Week
	currentIndex := 0
	for _, labelProgressWeek := range *labelProgressWeeks {
		if labelProgressWeek.Week != currentWeek {
			currentIndex++
			currentWeek = labelProgressWeek.Week
		}
		result[currentIndex] = append(result[currentIndex], &LabelProgress{
			Name:     labelProgressWeek.Name,
			Progress: labelProgressWeek.Progress,
		})
	}

	return result, nil
}

// GetLabelCount returns user's label count this week
func (db *LabelDatastore) GetLabelCount(ctx context.Context, userID string) (int64, error) {

	type Count struct {
		Count int64 `boil:"count"`
	}

	var data Count;
	if err := queries.Raw(`
		select count(*) as count
		from labels
		where user_id = $1
	  	and date_trunc('week', created_at) = date_trunc('week', now());
	`, userID).Bind(ctx, db, &data); err != nil {
		return 0, err
	}

	return data.Count, nil
}

// LabelScore holds label leaderboard information.
type LabelScore struct {
	Name  string `json:"name"`
	Score uint32 `json:"score"`
}
