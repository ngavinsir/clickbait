package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/ngavinsir/clickbait/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
)

var jwtAuth = jwtauth.New("HS256", []byte("clickbait^secret"), nil)

// UserIDCtxKey to extract user id from context
var UserIDCtxKey = &contextKey{"User_id"}

// Register new user handler
func Register(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &RegisterRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		
		hash, _ := hashPassword(data.Password)
		user := &models.User{
			ID: ksuid.New().String(),
			Username: data.Username,
			Password: hash,
		}
		
		if err := user.Insert(r.Context(), db, boil.Infer()); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		tokenString, err := loginLogic(r.Context(), db, data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		render.JSON(w, r, tokenString)
	})
}

// Login handler
func Login(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := &LoginRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		tokenString, err := loginLogic(r.Context(), db, data.User)
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		
		render.JSON(w, r, tokenString)
	})
}

func loginLogic(ctx context.Context, exec boil.ContextExecutor, data *User) (string, error) {
	user, err := models.Users(models.UserWhere.Username.EQ(data.Username)).One(ctx, exec)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(data.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	_, tokenString, _ := jwtAuth.Encode(jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
	})

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
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

// User general struct
type User struct {
	ID     		string `json:"id,omitempty"`
	Username  	string `json:"username"`
	Password	string `json:"password,omitempty"`
}

// RegisterRequest struct
type RegisterRequest struct {
	*User
}

// Bind RegisterRequest (Username, Password) [Required]
func (req *RegisterRequest) Bind(r *http.Request) error {
	if req.Username == "" || req.Password == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

// LoginRequest struct
type LoginRequest struct {
	*User
}

// Bind LoginRequest (Username, Password) [Required]
func (req *LoginRequest) Bind(r *http.Request) error {
	if req.Username == "" || req.Password == "" {
		return errors.New(ErrMissingReqFields)
	}

	return nil
}

type contextKey struct {
	name string
}