package middlewares

import (
	"net/http"

	"github.com/fsunset/twittApp/database"
)

// CheckingDB checks DB status
func CheckingDB(nextRequest http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// If existent; return error
		if !database.CheckConnection() {
			http.Error(w, "Bad DB connection", 500)
			return
		}

		// If DB connection is OK; pass "ResponseWriter" & "Request" to routers.Register endpoint
		nextRequest.ServeHTTP(w, r)
	}
}
