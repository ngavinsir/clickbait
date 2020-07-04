package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
		render.Render(w, r, ErrRender(ErrMissingReqFields))
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

	labels, err := env.labelRepository.GetLabel(r.Context(), user.ID, labelType)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	var labelResponse []*LabelResponse
	for _, label := range labels {
		var keywords []string
		for _, keyword := range label.R.ClickbaitKeywords {
			keywords = append(keywords, keyword.Keyword)
		}
		labelResponse = append(labelResponse, &LabelResponse{
			Label:    label,
			Article:  label.R.Article,
			Keywords: keywords,
		})
	}

	render.JSON(w, r, labelResponse)
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
	if labelType == "clickbait" && data.Value == "Clickbait" && data.Keywords != nil {
		if err = env.clickbaitKeywordRepository.AddClickbaitKeywords(
			context.Background(),
			label.ID,
			data.Keywords,
		); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
	response.LabelID = label.ID
	response.Keywords = data.Keywords

	article, _ := env.articleRepository.GetRandomArticle(r.Context(), user.ID, labelType)
	response.Article = article

	render.JSON(w, r, response)
}

// GetLabelLeaderboard returns labeler score leaderboard.
func (env *Env) GetLabelLeaderboard(w http.ResponseWriter, r *http.Request) {
	labelType := chi.URLParam(r, "labelType")
	limit := chi.URLParam(r, "limit")
	limitU64, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	if limitU64 < 0 {
		limitU64 = 0
	}

	clickbaitLeaderboard, err := env.labelRepository.GetLabelLeaderboard(r.Context(), labelType, uint8(limitU64))
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, clickbaitLeaderboard)
}

// LabelRequest for add label handler request
type LabelRequest struct {
	*models.Label
	Keywords []string `json:"keywords"`
}

// Bind label request if value and headline_id are present
func (req *LabelRequest) Bind(r *http.Request) error {
	if req.Label == nil || req.Value == "" || req.ArticleID == "" {
		return ErrMissingReqFields
	}

	return nil
}

// ClickbaitResponse contains label_id and new_headline
type ClickbaitResponse struct {
	LabelID         string   `boil:"label_id" json:"label_id"`
	Keywords        []string `json:"keywords,omitempty"`
	*models.Article `boil:"new_article" json:"new_article"`
}

// LabelResponse contains label, keywords and article
type LabelResponse struct {
	*models.Label   `json:"label"`
	*models.Article `json:"article"`
	Keywords        []string `json:"keywords,omitempty"`
}
