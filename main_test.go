package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetInstruction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/ff/ff-test", nil)
	if err != nil {
		t.Errorf("Erro no retorno %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/api/v1/ff/ falhou com o respectivo codigo de erro %d.", resp.Code)
	}
}
