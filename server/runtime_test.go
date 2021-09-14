package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	handler := runtime()
	resp := response(t, get(t, "/"), handler)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	ok, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "OK", string(ok))
}

func response(t *testing.T, r *http.Request, h http.Handler) *http.Response {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w.Result()
}

func get(t *testing.T, url string) (req *http.Request) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err, "failed to create request:", url)
	return
}
