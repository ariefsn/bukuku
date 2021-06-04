package helper

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
)

var tokenAuth *jwtauth.JWTAuth

func InitJwt() {
	secret := "This is default secret jwt"

	if os.Getenv("JWT_SECRET") != "" {
		secret = os.Getenv("JWT_SECRET")
	}

	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func EncodeJwt(claims map[string]interface{}) (token jwt.Token, tokenString string, err error) {
	return tokenAuth.Encode(claims)
}

func DecodeJwt(r *http.Request) (token jwt.Token, claims map[string]interface{}, err error) {
	return jwtauth.FromContext(r.Context())
}

func TokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			render.Render(w, r, ResponseError(http.StatusUnauthorized, err))

			return
		}

		if token == nil || jwt.Validate(token) != nil {
			render.Render(w, r, ResponseError(http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized))))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
