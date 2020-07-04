package handlers

import "github.com/ngavinsir/clickbait/model"

// Env holds handler's datasource information.
type Env struct {
	labelRepository            model.LabelRepository
	userRepository             model.UserRepository
	articleRepository          model.ArticleRepository
	clickbaitKeywordRepository model.ClickbaitKeywordRepository
}

// CreateEnv creates new env with the same database.
func CreateEnv(db *model.DB) *Env {
	return &Env{
		labelRepository:            &model.LabelDatastore{DB: db},
		userRepository:             &model.UserDatastore{DB: db},
		articleRepository:          &model.ArticleDatastore{DB: db},
		clickbaitKeywordRepository: &model.CLickbaitKeywordDatastore{DB: db},
	}
}
