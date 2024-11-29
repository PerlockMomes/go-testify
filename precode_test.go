package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое

func TestMainHandlerWhenResponseIsOk(t *testing.T) {
	request := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)

	status := responseRecorder.Code
	body := responseRecorder.Body.String()
	method := "GET"

	require.Equal(t, method, request.Method, "Method should be GET")
	assert.Equal(t, http.StatusOK, status, "Status should be 200")
	assert.NotEmpty(t, body, "Body should not be empty")
}

// город, который передаётся в параметре city, не поддерживается,
// cервис возвращает код ответа 400 и ошибку wrong city value в теле ответа

func TestMainHandlerWhenCityDoesNotExist(t *testing.T) {
	request := httptest.NewRequest("GET", "/cafe?count=4&city=samara", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)

	status := responseRecorder.Code
	body := responseRecorder.Body.String()
	expectedBody := `wrong city value`

	assert.Equal(t, http.StatusBadRequest, status, "Status should be 400")
	assert.Equal(t, expectedBody, body, "Body should be: wrong city value")
}

// если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	request := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)

	body := responseRecorder.Body.String()
	count := strings.Split(body, ",")
	totalCount := 4

	assert.Equal(t, totalCount, len(count))
}
