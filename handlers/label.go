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

func AddLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)

		data := &AddLabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		label := &models.Label{
			ID: ksuid.New().String(),
			UserID: userID,
			HeadlineID: data.HeadlineID,
			Value: data.Value,
		}

		if err := label.Insert(r.Context(), db, boil.Infer()); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, label)
	})
}

type Label struct {
	ID			string `json:"id"`
	UserID		string `json:"user_id"`
	HeadlineID	string `json:"headline_id"`
	Value		string `json:"value"`
}

type AddLabelRequest struct {
	*Label
}

func (req *AddLabelRequest) Bind(r *http.Request) error {
	if req.Label == nil || req.Value == "" || req.HeadlineID == "" {
		return errors.New(ErrMissingFields)
	}

	return nil
}