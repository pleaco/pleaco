package pleaco

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRunningContainersRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/containers", nil)
	router.ServeHTTP(w, req)

	// Test if we receive 200 statuscode
	assert.Equal(t, 200, w.Code)
}

func TestRunContainersRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	jsonBody := []byte(`{"image": "hello, server!", "tag": "latest", "status": "running", "hasNode": "false"}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/run", bodyReader)
	router.ServeHTTP(w, req)

	// Test if we receive 200 statuscode
	assert.Equal(t, 201, w.Code)
}
