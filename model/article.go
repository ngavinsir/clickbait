package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// InsertArticle with headline and content
func InsertArticle(ctx context.Context, exec boil.ContextExecutor, headline string, content string) (*models.Article, error) {
	return insertArticleWithID(ctx, exec, ksuid.New().String(), headline, content)
}

func insertArticleWithID(
	ctx context.Context, exec boil.ContextExecutor, id string, headline string, content string,
) (*models.Article, error) {
	if id == "" {
		id = ksuid.New().String()
	}

	article := &models.Article{
		ID:       id,
		Headline: headline,
		Content:  content,
	}

	if err := article.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}

	return article, nil
}

// GetRandomArticle which labeled less than 3 times and haven't labeled by the user before
func GetRandomArticle(ctx context.Context, exec boil.ContextExecutor, userID string, labelType string) (*models.Article, error) {
	if labelType == "" {
		return nil, errors.New("invalid label type")
	}

	article := &models.Article{}
	err := queries.Raw(`
		select a.id, a.headline, a.content
		from articles a left join labels l on a.id = l.article_id and l.type = $1
		where a.id not in (
			select l.article_id
			from labels l
			where l.user_id = $2 and l.type = $1
		)
		group by a.id
		having count(l.id) < $3
		order by random()
		limit 1
	`, labelType, userID, 3).Bind(ctx, exec, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}
