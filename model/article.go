package model

import (
	"context"
	"errors"

	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// ArticleRepository handles article data management.
type ArticleRepository interface {
	InsertArticle(ctx context.Context, headline string, content string) 	(*models.Article, error)
	GetRandomArticle(ctx context.Context, userID string, labelType string)	(*models.Article, error)
}

// ArticleDatastore holds db information.
type ArticleDatastore struct {
	*DB
}

// InsertArticle with headline and content
func (db *ArticleDatastore) InsertArticle(ctx context.Context, headline string, content string) (*models.Article, error) {
	return db.insertArticleWithID(ctx, ksuid.New().String(), headline, content)
}

func (db *ArticleDatastore) insertArticleWithID(ctx context.Context, id string, headline string, content string) (*models.Article, error) {
	if id == "" {
		id = ksuid.New().String()
	}

	article := &models.Article{
		ID:       id,
		Headline: headline,
		Content:  content,
	}

	if err := article.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return article, nil
}

// GetRandomArticle which labeled less than 3 times and haven't labeled by the user before
func (db *ArticleDatastore) GetRandomArticle(ctx context.Context, userID string, labelType string) (*models.Article, error) {
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
	`, labelType, userID, 3).Bind(ctx, db, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}
