package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
)

func RandomHeadline(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		if userID == "" {
			render.Render(w, r, ErrUnauthorized(errors.New(ErrInvalidUserID)))
		}
	})
}

func AddHeadline(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &AddHeadlineRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		headline := &models.Headline{
			ID: ksuid.New().String(),
			Value: data.Value,
		}

		if err := headline.Insert(r.Context(), db, boil.Infer()); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, headline)
	})
}

type Headline struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type AddHeadlineRequest struct {
	*Headline
}

func (req *AddHeadlineRequest) Bind(r *http.Request) error {
	if req.Headline == nil || req.Value == "" {
		return errors.New(ErrMissingFields)
	}

	return nil
}