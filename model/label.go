package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// InsertLabel with maximum of 3 labels per headline
func InsertLabel(ctx context.Context, exec boil.ContextExecutor, 
	userID string, headlineID string, value string) (*models.Label, error) {
		headlineLabelCount, err := GetHeadlineLabelCount(ctx, exec, headlineID)
		if err != nil {
			return nil, err
		}

		if headlineLabelCount >= 3 {
			return nil, errors.New("maximum label reached")
		}

		label := &models.Label{
			ID: ksuid.New().String(),
			UserID: userID,
			HeadlineID: headlineID,
			Value: value,
		}

		if err := label.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		return label, nil
}

// DeleteLabel helper
func DeleteLabel(ctx context.Context, exec boil.ContextExecutor, labelID string) error {
	_, err := models.Labels(models.LabelWhere.ID.EQ(labelID)).DeleteAll(ctx, exec)
	if err != nil {
		return err
	}

	return nil
}

func GetHeadlineLabel(ctx context.Context, exec boil.ContextExecutor, userID string) (*[]HeadlineLabel, error) {
	var data []HeadlineLabel
	err := models.NewQuery(
		Select("labels.id as id", "headlines.id as headline_id",
		"headlines.value as headline_value", "labels.value as label_value"),
		From("labels"),
		InnerJoin("headlines on labels.headline_id = headlines.id"),
		models.LabelWhere.UserID.EQ(userID),
	).Bind(ctx, exec, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// GetHeadlineLabelCount return label count by headline id
func GetHeadlineLabelCount(ctx context.Context, exec boil.ContextExecutor, headlineID string) (int64, error) {
	labelCount, err := models.Labels(models.LabelWhere.HeadlineID.EQ(headlineID)).Count(ctx, exec)
	if err != nil {
		return 0, err
	}

	return labelCount, nil
}

type HeadlineLabel struct {
	ID				string	`boil:"id" json:"id"`
	HeadlineID		string	`boil:"headline_id" json:"headline_id"`
	HeadlineValue	string	`boil:"headline_value" json:"headline_value"`
	LabelValue		string	`boil:"label_value" json:"label_value"`
}