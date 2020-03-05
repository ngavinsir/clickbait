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
	"github.com/volatiletech/sqlboiler/queries"
)

// RandomHeadline handler
func RandomHeadline(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(UserIDCtxKey).(string)

		headline, err := GetRandomHeadline(r.Context(), db, userID)
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

// GetRandomHeadline which labeled less than 3 times and haven't labeled by the user before
func GetRandomHeadline(ctx context.Context, exec boil.ContextExecutor, userID string) (*Headline, error) {
	headline := &Headline{}
	err := queries.Raw(`
		select h.id, h.value
		from headlines h left join labels l on h.id = l.headline_id
		where h.id not in (
			select l.headline_id
			from labels l
			where l.user_id = $1
			group by l.headline_id
		)
		group by h.id
		having count(l.id) < $2
		order by random()
		limit 1
	`, userID, 3).Bind(ctx, exec, headline)

	if err != nil {
		return &Headline{}, err
	}

	return headline, nil
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
		return errors.New(ErrMissingReqFields)
	}

	return nil
}