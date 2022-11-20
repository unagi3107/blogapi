package controllers_test

import (
	"testing"

	"github.com/ura3107/blogapi/controllers"
	"github.com/ura3107/blogapi/controllers/testdata"
)

var aCon *controllers.ArticleController
var cCon *controllers.CommentController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	cCon = controllers.NewCommentController(ser)

	m.Run()
}
