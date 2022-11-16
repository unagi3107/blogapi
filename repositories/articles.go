package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ura3107/blogapi/models"
)

const (
	articlePerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	articleID, _ := result.LastInsertId()
	newArticle.ID = int(articleID)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	var articleArray []models.Article
	rows, err := db.Query(sqlStr, articlePerPage, (page-1)*articlePerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return nil, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	var article models.Article

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?;`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
