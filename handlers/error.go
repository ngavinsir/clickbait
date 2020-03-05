package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

const ErrInvalidUserID = "ERR_INVALID_USER_ID"
const ErrMissingFields = "ERR_MISSING_FIELDS"

// HandleErr to handle common error
func HandleErr(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`          
	AppCode    int64  `json:"code,omitempty"`  
	ErrorText  string `json:"error,omitempty"` 
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		StatusText:     "Unauthorized",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}