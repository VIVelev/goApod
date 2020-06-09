package models

import (
	"database/sql"
	"time"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Article model
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	ImagePath string    `json:"imagePath"`
	Text      string    `json:"text"`
	AuthorID  int       `json:"authorId"`
	Date      time.Time `json:"date"`
	EventID   int       `json:"eventId"`
}

// SQL statements
var statements = map[string]string{
	"GetAllArticles": "select * from articles",
	"GetArticleByID": "select * from articles where id = $1",
	"Save": `insert into articles values
	 (default, $1, $2, $3, $4, $5, $6)`,
	"Delete": `delete from articles where id = $1`,
}

// GetAllArticles - returns all articles
func GetAllArticles() ([]Article, error) {
	rows, err := database.Db.Query(statements["GetAllArticles"])
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(
			&article.ID, &article.Title,
			&article.ImagePath, &article.Text,
			&article.AuthorID, &article.Date,
			&article.EventID); err != nil {

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
		&ret.Text, &ret.AuthorID, &ret.Date, &ret.EventID); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "articles", ID: id}
	default:
		return ret, err
	}
}

// Save a new article
func (a *Article) Save() error {
	if _, err := database.Db.Exec(statements["Save"],
		a.Title, a.ImagePath,
		a.Text, a.AuthorID,
		a.Date, a.EventID); err != nil {

		return err
	}

	return nil
}

// Delete an article
func (a *Article) Delete() error {
	result, err := database.Db.Exec(statements["Delete"], a.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "articles", ID: a.ID}
	}

	return nil
}
