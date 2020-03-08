package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
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

// DeleteLabel handler
func DeleteLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if labelID := chi.URLParam(r, "labelID"); labelID != "" {
			err := model.DeleteLabel(r.Context(), db, labelID)
			if err != nil {
				render.Render(w, r, ErrRender(err))
				return
			}
		}
	})
}

func GetAllLabel(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)
		
		labels := []*model.HeadlineLabel{}
		labels, err := model.GetHeadlineLabel(r.Context(), db, userID)
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

		data := &LabelRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		_, err := model.InsertLabel(r.Context(), db, userID, data.HeadlineID, data.Value)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		headline, err := model.GetRandomHeadline(r.Context(), db, userID)
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