package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
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

// GetHeadlineLabelCount return label count by headline id
func GetHeadlineLabelCount(ctx context.Context, exec boil.ContextExecutor, headlineID string) (int64, error) {
	labelCount, err := models.Labels(models.LabelWhere.HeadlineID.EQ(headlineID)).Count(ctx, exec)
	if err != nil {
		return 0, err
	}

	return labelCount, nil
}