package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/model"
)

// AddLabel handler
func AddLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)

		data := &LabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		label, err := model.InsertLabel(r.Context(), db, userID, data.HeadlineID, data.Value)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, label)
	})
}

// Clickbait handler return new headline after labeled previous headline
func Clickbait(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)

		data := &LabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		model.InsertLabel(r.Context(), tx, userID, data.HeadlineID, data.Value)
		headline, _ := GetRandomHeadline(r.Context(), tx, userID)

		err = tx.Commit()
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, headline)
	})
}

// Label generat struct
type Label struct {
	ID			string `json:"id"`
	UserID		string `json:"user_id"`
	HeadlineID	string `json:"headline_id"`
	Value		string `json:"value"`
}

// LabelRequest for add label handler request
type LabelRequest struct {
	*Label
}

// Bind label request if value and headline_id are present
func (req *LabelRequest) Bind(r *http.Request) error {
	if req.Label == nil || req.Value == "" || req.HeadlineID == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}