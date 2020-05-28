package middlewares

import (
	"net/http"

	"github.com/fsunset/twittApp/routers"
)

// CheckingJWT validates correct entering-by-request-JWT structure & data
func CheckingJWT(nextRequest http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// "Authorization" contains user's JWT from Request, which we use to validate it
		_, _, _, err := routers.ValidateJWT(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		// If JWT validation passes; pass "ResponseWriter" & "Request" to routers.Profile endpoint
		nextRequest.ServeHTTP(w, r)
	}
}
