package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/model"
	"github.com/ngavinsir/clickbait/models"
)

// RandomArticle handler
func RandomArticle(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		labelType := chi.URLParam(r, "labelType")

		article, err := model.GetRandomArticle(r.Context(), db, userID, labelType)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, article)
	})
}

// AddArticle handler
func AddArticle(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &AddArticleRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		article, err := model.InsertArticle(r.Context(), db, data.Headline, data.Content)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, article)
	})
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
