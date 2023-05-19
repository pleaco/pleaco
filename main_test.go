package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRunningContainersRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/containers", nil)
	router.ServeHTTP(w, req)

	// Test if we receive 200 statuscode
	assert.Equal(t, 200, w.Code)
}
