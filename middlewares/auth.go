package middlewares

import (
	"net/http"

	"github.com/dg943/MyProject/backend/helpers"
)

var (
	AllowAnonymousPaths = map[string]helpers.Empty{
		"/login":  helpers.Empty{},
		"/signup": helpers.Empty{},
	}
)

func AuthenticateMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// bypassing the authentication for signup and login paths
		path := r.URL.Path
		if _, ok := AllowAnonymousPaths[path]; ok {
			next.ServeHTTP(w, r)
			return
		}

		_, err := AuthenticateJWT(r)

		if err != nil {
			http.Error(w, "Unauthorized access", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func AuthenticateJWT(r *http.Request) (*helpers.Claims, error) {
	tokenString := r.Header.Get("Authorization")
	return helpers.ParseJWT(tokenString)
}
