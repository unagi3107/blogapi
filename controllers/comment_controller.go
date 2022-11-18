package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ura3107/blogapi/controllers/services"
	"github.com/ura3107/blogapi/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
