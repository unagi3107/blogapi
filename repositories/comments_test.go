package repositories_test

import (
	"testing"

	"github.com/ura3107/blogapi/models"
	"github.com/ura3107/blogapi/repositories"
	"github.com/ura3107/blogapi/repositories/testdata"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test",
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("article id is expect %d but get %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments where article_id = ? and Message = ?
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}

func TestSelectCommentList(t *testing.T) {
	expectedNum := len(testdata.CommentTestData)
	articleID := 1

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	if expectedNum != len(got) {
		t.Errorf("comments length is expect %d but get %d\n", expectedNum, len(got))
	}
}
