package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/model"
	"github.com/ngavinsir/clickbait/models"
)

// RandomHeadline handler
func RandomHeadline(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)

		headline, err := model.GetRandomHeadline(r.Context(), db, userID)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, headline)
	})
}

// AddHeadline handler
func AddHeadline(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &AddHeadlineRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		headline, err := model.InsertHeadline(r.Context(), db, data.Value)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, headline)
	})
}

// AddHeadlineRequest struct
type AddHeadlineRequest struct {
	*models.Headline
}

// Bind add headline request if headline_value is not missing
func (req *AddHeadlineRequest) Bind(r *http.Request) error {
	if req.Headline == nil || req.Value == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}
