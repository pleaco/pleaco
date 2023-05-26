package pleaco

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = SetupRouter()

func TestGetRunningContainersRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/containers", nil)
	router.ServeHTTP(w, req)

	// Test if we receive 200 statuscode
	assert.Equal(t, 200, w.Code)
}

func TestRunContainersRouteOK(t *testing.T) {

	w := httptest.NewRecorder()
	jsonBody := []byte(`{"image": "hello, server!", "tag": "latest", "status": "running", "hasNode": false}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/run", bodyReader)
	router.ServeHTTP(w, req)

	// Test if we receive 201 statuscode
	assert.Equal(t, 201, w.Code)
}

func TestRunContainersRouteMethodNotAllowed(t *testing.T) {

	w := httptest.NewRecorder()
	jsonBody := []byte(`{"image": "hello, server!", "tag": "latest", "status": "somethingOtherThanRunning", "hasNode": false}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/run", bodyReader)
	router.ServeHTTP(w, req)

	// Test if we receive 201 statuscode
	assert.Equal(t, 405, w.Code)
}

func TestRunContainersRouteFail(t *testing.T) {

	w := httptest.NewRecorder()
	jsonBody := []byte(``)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/run", bodyReader)
	router.ServeHTTP(w, req)

	// Test if we receive 201 statuscode
	assert.Equal(t, 400, w.Code)
}
