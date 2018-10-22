package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte(""))

	req, err := http.NewRequest("GET", "/health", body)
	if err != nil {
		t.Errorf("Get heartbeat failed with error %d.", err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/health failed with error code %d.", resp.Code)
	}
}

// func TestMain(m *testing.M) {
// 	go func() {
// 		os.Exit(m.Run())
// 	}()
// }
