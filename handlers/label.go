package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/model"
	"github.com/ngavinsir/clickbait/models"
)

// AddLabel handler
func (env *Env) AddLabel(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCtxKey).(*models.User)
	labelType := chi.URLParam(r, "labelType")

	data := &LabelRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	label, err := env.labelRepository.InsertLabel(r.Context(), user.ID, data.ArticleID, data.Value, labelType)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, label)
}

// DeleteLabel handler
func (env *Env) DeleteLabel(w http.ResponseWriter, r *http.Request) {
	labelID := chi.URLParam(r, "labelID")
	if labelID == "" {
		render.Render(w, r, ErrRender(errors.New(ErrMissingReqFields)))
	}

	err := env.labelRepository.DeleteLabel(r.Context(), labelID)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// GetAllLabel handler
func (env *Env) GetAllLabel(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCtxKey).(*models.User)
	labelType := chi.URLParam(r, "labelType")

	labels := []*model.ArticleLabel{}
	labels, err := env.labelRepository.GetArticleLabel(r.Context(), user.ID, labelType)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, labels)
}

// Labeling handler return new headline after labeled previous headline
func (env *Env) Labeling(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(UserCtxKey).(*models.User)
	labelType := chi.URLParam(r, "labelType")

	data := &LabelRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	response := &ClickbaitResponse{}
	label, err := env.labelRepository.InsertLabel(r.Context(), user.ID, data.ArticleID, data.Value, labelType)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	response.LabelID = label.ID

	article, _ := env.articleRepository.GetRandomArticle(r.Context(), user.ID, labelType)
	response.Article = article

	render.JSON(w, r, response)
}

// LabelRequest for add label handler request
type LabelRequest struct {
	*models.Label
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
