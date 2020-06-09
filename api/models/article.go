package models

import (
	"database/sql"
	"time"

	"github.com/VIVelev/goApod/database"
)

// Article model
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ImagePath string    `json:"imagePath"`
	Text      string    `json:"text"`
	AuthorID  int       `json:"authorId"`
	Date      time.Time `json:"date"`
}

var statements = map[string]string{
	"GetArticles":    "select * from articles",
	"GetArticleByID": "select * from articles where id = $1",
}

// GetArticles - returns all articles
func GetArticles() ([]Article, error) {
	rows, err := database.Db.Query(statements["GetArticles"])
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.ImagePath,
			&article.Text, &article.AuthorID, &article.Date); err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// GetArticleByID - get single article by id
func GetArticleByID(id int) (Article, error) {
	row := database.Db.QueryRow(statements["GetArticleByID"])
	var ret Article

	switch err := row.Scan(&ret.ID, &ret.Title, &ret.ImagePath,
		&ret.Text, &ret.AuthorID, &ret.Date); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, nil
	default:
		return ret, err
	}
}
