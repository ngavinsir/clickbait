package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
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

		label, err := insertLabel(r.Context(), db, data, userID)
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

		insertLabel(r.Context(), tx, data, userID)
		headline, _ := GetRandomHeadline(r.Context(), tx, userID)

		err = tx.Commit()
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, headline)
	})
}

func insertLabel(ctx context.Context, exec boil.ContextExecutor, data *LabelRequest, userID string) (*models.Label, error) {
	label := &models.Label{
		ID: ksuid.New().String(),
		UserID: userID,
		HeadlineID: data.HeadlineID,
		Value: data.Value,
	}

	if err := label.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}

	return label, nil
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