package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenSuccess(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
	require.Len(t, strings.Split(responseRecorder.Body.String(), ","), 4) // Проверяем, что вернулось 4 кафе
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=spb", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "count missing", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountWrongValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=wrong&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong count value", responseRecorder.Body.String())
}
