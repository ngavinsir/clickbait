package model

import (
	"context"
	"errors"
	"fmt"

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
	query := fmt.Sprintf(`
		select a.id, a.headline, a.content
		from articles a left join %s_labels l on a.id = l.article_id
		where a.id not in (
			select l.article_id
			from %s_labels l
			where l.user_id = $1
			group by l.article_id
		)
		group by a.id
		having count(l.id) < $2
		order by random()
		limit 1
	`, labelType, labelType)
	err := queries.Raw(query, userID, 3).Bind(ctx, exec, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}
