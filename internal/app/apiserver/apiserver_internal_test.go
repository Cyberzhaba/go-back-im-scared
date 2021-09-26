package apiserver

import (
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAPIserver_handleMain(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	s.Ping()
	assert.Equal(t, rec.Body.String(), "{'message': 'pong'}")
}
