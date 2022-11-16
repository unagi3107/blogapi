package repositories

import (
	"database/sql"

	"github.com/ura3107/blogapi/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values (?, ?, now());
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	newComment.CommentID = int(commentID)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select * 
		from comments
		where article_id = ?;
	`

	var commentArray []models.Comment

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			return nil, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
