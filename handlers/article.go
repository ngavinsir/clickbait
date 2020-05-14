package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/models"
)

// RandomArticle handler
func (env *Env) RandomArticle(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCtxKey).(*models.User)
	labelType := chi.URLParam(r, "labelType")

	article, err := env.articleRepository.GetRandomArticle(r.Context(), user.ID, labelType)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, article)
}

// AddArticle handler
func (env *Env) AddArticle(w http.ResponseWriter, r *http.Request) {
	data := &AddArticleRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	article, err := env.articleRepository.InsertArticle(r.Context(), data.Headline, data.Content)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, article)
}

// AddArticleRequest struct
type AddArticleRequest struct {
	*models.Article
}

// Bind add headline request if headline_value is not missing
func (req *AddArticleRequest) Bind(r *http.Request) error {
	if req.Article == nil || req.Headline == "" || req.Content == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}
