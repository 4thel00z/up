// Package gzip provides gzip compression support.
package gzip

import (
	"net/http"

	"github.com/NYTimes/gziphandler"

	"github.com/4thel00z/up"
)

// New gzip handler.
func New(c *up.Config, next http.Handler) http.Handler {
	return gziphandler.GzipHandler(next)
}
