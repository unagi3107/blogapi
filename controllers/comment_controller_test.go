package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostCommentHandler(t *testing.T) {
	var tests = []struct {
		name       string
		comment    []byte
		statusCode int
	}{
		{
			name:       "good request body",
			comment:    []byte(`{"article_id":1,"message":"first comment"}`),
			statusCode: http.StatusOK,
		},
		{
			name:       "bad request body",
			comment:    []byte(``),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url := "http://localhost:8080/comment"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(tt.comment))

			res := httptest.NewRecorder()

			cCon.PostCommentHandler(res, req)

			if res.Code != tt.statusCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.statusCode, res.Code)
			}
		})
	}
}
