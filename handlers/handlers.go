package handlers

import "github.com/ngavinsir/clickbait/model"

// Env holds handler's datasource information.
type Env struct {
	labelRepository		model.LabelRepository
	userRepository		*model.DB
	articleRepository	model.ArticleRepository
}

// CreateEnv creates new env with the same database.
func CreateEnv(db *model.DB) *Env {
	return &Env{
		labelRepository: &model.LabelDatastore{DB: db},
		userRepository: db,
		articleRepository: &model.ArticleDatastore{DB: db},
	}
}