package src

import (
	"net/http"
	"os"
)

const (
	xForwardedProtoHeader = "x-forwarded-proto"
	debugEnv              = "DEBUG"
)

// ForceSSL enforces HTTPS when not in debug mode
func ForceSSL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv(debugEnv) == "false" {
			if r.Header.Get(xForwardedProtoHeader) != "https" {
				sslURL := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslURL, http.StatusTemporaryRedirect)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
