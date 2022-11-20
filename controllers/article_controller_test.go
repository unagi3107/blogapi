package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestArticleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			aCon.ArticleListHandler(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		articleID  string
		statusCode int
	}{
		{name: "number query", articleID: "1", statusCode: http.StatusOK},
		{name: "alphabet query", articleID: "aaa", statusCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.statusCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.statusCode, res.Code)
			}
		})
	}
}

func TestPostArticleHandler(t *testing.T) {
	var tests = []struct {
		name       string
		article    []byte
		statusCode int
	}{
		{
			name:       "good request body",
			article:    []byte(`{"title":"a","contents":"b","username":"c"}`),
			statusCode: http.StatusOK,
		},
		{
			name:       "bad request body",
			article:    []byte(``),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url := "http://localhost:8080/article"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(tt.article))

			res := httptest.NewRecorder()

			aCon.PostArticleHandler(res, req)

			if res.Code != tt.statusCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.statusCode, res.Code)
			}
		})
	}
}

func TestPostNiceHandler(t *testing.T) {
	var tests = []struct {
		name       string
		article    []byte
		statusCode int
	}{
		{
			name:       "good request body",
			article:    []byte(`{"title":"a","contents":"b","username":"c","nice":1}`),
			statusCode: http.StatusOK,
		},
		{
			name:       "bad request body",
			article:    []byte(``),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url := "http://localhost:8080/article/nice"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(tt.article))

			res := httptest.NewRecorder()

			aCon.PostNiceHandler(res, req)

			if res.Code != tt.statusCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.statusCode, res.Code)
			}
		})
	}
}
