package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	_, close, r, err := SetupAPIServer()
	defer close()
	if err != nil {
		log.Fatal().Err(err).Msg("test api server setup fail")
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ready", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "ready", w.Body.String())
}
