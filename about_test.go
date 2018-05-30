package go_rest

import (
	"testing"

	"net/http"

	"github.com/matryer/is"
)

func TestHandleAbout(t *testing.T) {
	is := is.New(t)
	_, err := NewServer("localhost", 12345)
	is.NoErr(err)
	_, err = http.NewRequest("GET", "/about", nil)
	is.NoErr(err)
	/*w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.StatusCode, http.StatusOK)*/
}
