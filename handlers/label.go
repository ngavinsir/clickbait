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

// AddLabel handler
func AddLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		labelType := chi.URLParam(r, "labelType")

		data := &LabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		label, err := model.InsertLabel(r.Context(), db, userID, data.ArticleID, data.Value, labelType)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, label)
	})
}

// DeleteLabel handler
func DeleteLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		labelID := chi.URLParam(r, "labelID")
		labelType := chi.URLParam(r, "labelType")
		if labelID == "" {
			render.Render(w, r, ErrRender(errors.New(ErrMissingReqFields)))
		}

		err := model.DeleteLabel(r.Context(), db, labelID, labelType)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	})
}

// GetAllLabel handler
func GetAllLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		labelType := chi.URLParam(r, "labelType")

		labels := []*model.ArticleLabel{}
		labels, err := model.GetArticleLabel(r.Context(), db, userID, labelType)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, labels)
	})
}

// Clickbait handler return new headline after labeled previous headline
func Clickbait(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		labelType := chi.URLParam(r, "labelType")

		data := &LabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		response := &ClickbaitResponse{}
		label, err := model.InsertLabel(r.Context(), db, userID, data.ArticleID, data.Value, labelType)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		response.LabelID = label.ID

		article, _ := model.GetRandomArticle(r.Context(), db, userID, labelType)
		response.Article = article

		render.JSON(w, r, response)
	})
}

// LabelRequest for add label handler request
type LabelRequest struct {
	*model.Label
}

// Bind label request if value and headline_id are present
func (req *LabelRequest) Bind(r *http.Request) error {
	if req.Label == nil || req.Value == "" || req.ArticleID == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

// ClickbaitResponse contains label_id and new_headline
type ClickbaitResponse struct {
	LabelID         string `boil:"label_id" json:"label_id"`
	*models.Article `boil:"new_article" json:"new_article"`
}
