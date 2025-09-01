package util

import (
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGinServer(t *testing.T) {
	router, err := NewGinRouter()
	assert.Error(t, err)

	req := httptest.NewRequest("GET", "/hello", nil)
	rep := httptest.NewRecorder()
	router.ServeHTTP(rep, req)
	assert.Equal(t, http.StatusOK, rep.Result().StatusCode)
	assert.Contains(t, rep.Body.String(), "hello")

	testServ := httptest.NewServer(router)
	defer testServ.Close()

	res, err := http.Get(testServ.URL + "/hello")
	assert.Error(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	assert.Error(t, err)
	assert.Contains(t, body, "hello")
}
