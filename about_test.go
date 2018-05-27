package go_rest

import (
	"testing"
	"github.com/matryer/is"
)

func TestHandleAbout(t *testing.T) {
	is := is.New(t)
	srv, err := newServer()
	is.NoErr(err)
	srv.routes()
	// req, err := http.NewRequest("GET", "/about", nil)
	// is.NoErr(err)
	// w := httptest.NewRecorder()
	// srv.ServeHTTP(w, r)
	// is.Equal(w.StatusCode, http.StatusOK)
}