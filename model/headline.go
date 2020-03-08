package model

import (
	"context"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// InsertHeadline with value
func InsertHeadline(ctx context.Context, exec boil.ContextExecutor, value string) (*models.Headline, error) {
	return insertHeadlineWithID(ctx, exec, ksuid.New().String(), value)
}

func insertHeadlineWithID(ctx context.Context, exec boil.ContextExecutor, id string, value string) (*models.Headline, error) {
	if id == "" {
		id = ksuid.New().String()
	}

	headline := &models.Headline{
		ID:    id,
		Value: value,
	}

	if err := headline.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}

	return headline, nil
}

// GetRandomHeadline which labeled less than 3 times and haven't labeled by the user before
func GetRandomHeadline(ctx context.Context, exec boil.ContextExecutor, userID string) (*models.Headline, error) {
	headline := &models.Headline{}
	err := queries.Raw(`
		select h.id, h.value
		from headlines h left join labels l on h.id = l.headline_id
		where h.id not in (
			select l.headline_id
			from labels l
			where l.user_id = $1
			group by l.headline_id
		)
		group by h.id
		having count(l.id) < $2
		order by random()
		limit 1
	`, userID, 3).Bind(ctx, exec, headline)

	if err != nil {
		return nil, err
	}

	return headline, nil
}
