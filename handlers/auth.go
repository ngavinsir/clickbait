package handlers

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/model"
	"github.com/ngavinsir/clickbait/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtAuth = jwtauth.New("HS256", []byte("clickbait^secret"), nil)
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// UserIDCtxKey to extract user id from context
var UserIDCtxKey = &contextKey{"User_id"}

// Register new user handler
func (env *Env) Register(w http.ResponseWriter, r *http.Request) {
	data := &RegisterRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if !emailRegexp.MatchString(data.Email) {
		render.Render(w, r, ErrRender(errors.New("invalid email")))
		return
	}

	_, err := env.userRepository.CreateNewUser(r.Context(), data.User)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	tokenString, err := loginLogic(r.Context(), env.userRepository, data.User)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, tokenString)
}

// Login handler
func (env *Env) Login(w http.ResponseWriter, r *http.Request) {
	data := &LoginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	tokenString, err := loginLogic(r.Context(), env.userRepository, data.User)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.JSON(w, r, tokenString)
}

func loginLogic(ctx context.Context, userRepository model.UserRepository, data *models.User) (string, error) {
	user, err := userRepository.GetUser(ctx, data.Email)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(data.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	_, tokenString, _ := jwtAuth.Encode(jwt.MapClaims{
		"user_id":  user.ID,
		"email": user.Email,
	})

	return tokenString, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthMiddleware to handle request jwt token
func AuthMiddleware(next http.Handler) http.Handler {
	return jwtauth.Verifier(jwtAuth)(extractUserID(next))
}

func extractUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDCtxKey, claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RegisterRequest struct
type RegisterRequest struct {
	*models.User
}

// Bind RegisterRequest (Email, Password, Age, Name) [Required]
func (req *RegisterRequest) Bind(r *http.Request) error {
	if req.Email == "" || req.Password == "" || req.Age <= 0 || req.Name == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

// LoginRequest struct
type LoginRequest struct {
	*models.User
}

// Bind LoginRequest (Email, Password) [Required]
func (req *LoginRequest) Bind(r *http.Request) error {
	if req.Email == "" || req.Password == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

type contextKey struct {
	name string
}
