package src

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// getScheme returns the request scheme
func getScheme(req *http.Request) string {
	if req.URL.Scheme != "" {
		return req.URL.Scheme
	}
	if req.TLS == nil {
		return "http"
	}
	return "https"
}

// ReverseProxy allows for forwarding of routes without redirects
func ReverseProxy(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		localURL, err := url.Parse(fmt.Sprintf("%s://%s%s", getScheme(c.Request), c.Request.Host, path))
		if err != nil {
			log.Println("Failed to parse URL:", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(localURL)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = localURL.Host
			req.URL.Scheme = localURL.Scheme
			req.URL.Host = localURL.Host
			req.URL.Path = localURL.Path
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
