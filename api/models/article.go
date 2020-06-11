package models

import (
	"database/sql"
	"time"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Article model
type Article struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	ImageURL string    `json:"imageUrl"`
	Text     string    `json:"text"`
	AuthorID int       `json:"authorId"`
	Date     time.Time `json:"date"`
	EventID  int       `json:"eventId"`

	LikesCount int `json:"likesCount"`
}

// SQL statements
var articleStatements = map[string]string{
	"GetAllArticles": `
		select a.*, count(l)
		from articles a
		left join likes l on a.id = l.article_id
		group by a.id`,

	"GetArticleByID": `
		select a.*, count(l)
		from articles a
		left join likes l on a.id = l.article_id
		where a.id = $1
		group by a.id`,

	"Save": `
		insert into articles
		values (default, $1, $2, $3, $4, $5, $6)`,

	"DeleteArticleByID": `
		delete from articles
		where id = $1`,

	"CountLikesByID": `
		select count(*)
		from likes
		where article_id = $1`,
}

// GetAllArticles - returns all articles
func GetAllArticles() ([]Article, errors.DatabaseError) {
	rows, err := database.Db.Query(articleStatements["GetAllArticles"])
	defer rows.Close()

	if err != nil {
		return nil, &errors.InternalDatabaseError{Message: err.Error()}
	}

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(
			&article.ID, &article.Title,
			&article.ImageURL, &article.Text,
			&article.AuthorID, &article.Date,
			&article.EventID, &article.LikesCount); err != nil {

			return nil, &errors.InternalDatabaseError{Message: err.Error()}
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// GetArticleByID - get single article by id
func GetArticleByID(id int) (Article, errors.DatabaseError) {
	row := database.Db.QueryRow(articleStatements["GetArticleByID"], id)
	var ret Article

	switch err := row.Scan(&ret.ID, &ret.Title, &ret.ImageURL,
		&ret.Text, &ret.AuthorID, &ret.Date, &ret.EventID, &ret.LikesCount); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "articles", ID: id}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// Save (insert) a new article
func (a *Article) Save() errors.DatabaseError {
	if _, err := database.Db.Exec(articleStatements["Save"],
		a.Title, a.ImageURL,
		a.Text, a.AuthorID,
		a.Date, a.EventID); err != nil {

		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	return nil
}

// Delete an article
func (a *Article) Delete() errors.DatabaseError {
	return DeleteArticleByID(a.ID)
}

// DeleteArticleByID - delete by id
func DeleteArticleByID(id int) errors.DatabaseError {
	result, err := database.Db.Exec(articleStatements["DeleteArticleByID"], id)
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "articles", ID: id}
	}

	return nil
}
