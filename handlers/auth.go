package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/ngavinsir/clickbait/model"
	"github.com/ngavinsir/clickbait/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtAuth *jwtauth.JWTAuth
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// UserCtxKey to extract user from context
var UserCtxKey = &contextKey{"User"}

func init() {
	godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(errors.New("env JWT_SECRET not provided"))
	}
	jwtAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
}

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

	password := data.User.Password
	_, err := env.userRepository.CreateNewUser(r.Context(), data.User)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	data.User.Password = password
	tokenString, err := loginLogic(r.Context(), env.userRepository, data.User)
	if err != nil {
		render.Render(w, r, ErrUnauthorized(err))
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
		render.Render(w, r, ErrUnauthorized(err))
		return
	}

	render.JSON(w, r, tokenString)
}

func loginLogic(ctx context.Context, userRepository model.UserRepository, data *models.User) (string, error) {
	user, err := userRepository.GetUser(ctx, data.Email)
	if err != nil {
		return "", ErrInvalidAccount
	}

	if !checkPasswordHash(data.Password, user.Password) {
		return "", ErrInvalidAccount
	}

	_, tokenString, _ := jwtAuth.Encode(jwt.MapClaims{
		"user_id":  user.ID,
		"email": user.Email,
		"name": user.Name,
	})

	return tokenString, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthMiddleware to handle request jwt token
func (env *Env) AuthMiddleware(next http.Handler) http.Handler {
	return jwtauth.Verifier(jwtAuth)(extractUser(env.userRepository)(next))
}

func extractUser(repo model.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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

			userID := claims["user_id"].(string)
			if userID == "" {
				render.Render(w, r, ErrRender(errors.New("invalid user id")))
				return
			}

			user, err := repo.GetUserbyID(r.Context(), userID)
			if err != nil {
				render.Render(w, r, ErrRender(errors.New("invalid user")))
				return
			}

			ctx := context.WithValue(r.Context(), UserCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RegisterRequest struct
type RegisterRequest struct {
	*models.User
}

// Bind RegisterRequest (Email, Password, Age, Name) [Required]
func (req *RegisterRequest) Bind(r *http.Request) error {
	if req.Email == "" || req.Password == "" || req.Age <= 0 || req.Name == "" {
		return ErrMissingReqFields
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
		return ErrMissingReqFields
	}

	return nil
}

type contextKey struct {
	name string
}
