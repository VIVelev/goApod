package models

import (
	"database/sql"
	"net/http"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Article model
type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
	Text     string `json:"text"`
	AuthorID int    `json:"authorId"`
	Date     string `json:"date"`

	// Foreign
	AuthorName string  `json:"authorName"`
	EventName  string  `json:"eventName"`
	LocLat     float64 `json:"locLat"`
	LocLong    float64 `json:"locLong"`
	LikesCount int     `json:"likesCount"`
}

// SQL statements
var articleStatements = map[string]string{
	"GetAllArticles": `
		select art.*, aut.name, eve.name, loc.lat, loc.long, count(l)
		from articles art
		left join authors aut on art.author_id = aut.id
		left join events eve on art.id = eve.article_id
		left join locations loc on eve.location_id = loc.id
		left join likes l on art.id = l.article_id
		group by art.id`,

	"GetArticleByID": `
		select art.*, aut.name, eve.name, loc.lat, loc.long, count(l)
		from articles art
		left join authors aut on art.author_id = aut.id
		left join events eve on art.id = eve.article_id
		left join locations loc on eve.location_id = loc.id
		left join likes l on art.id = l.article_id
		group by art.id
		where art.id = $1`,

	"GetArticleByDate": `
		select art.*, aut.name, eve.name, loc.lat, loc.long, count(l)
		from articles art
		left join authors aut on art.author_id = aut.id
		left join events eve on art.id = eve.article_id
		left join locations loc on eve.location_id = loc.id
		left join likes l on art.id = l.article_id
		group by art.id
		where art.date = $1`,

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
			&article.AuthorName, &article.EventName,
			&article.LocLat, &article.LocLong,
			&article.LikesCount); err != nil {

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

	switch err := row.Scan(
		&ret.ID, &ret.Title,
		&ret.ImageURL, &ret.Text,
		&ret.AuthorID, &ret.Date,
		&ret.AuthorName, &ret.EventName,
		&ret.LocLat, &ret.LocLong,
		&ret.LikesCount); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "articles", ID: id}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// GetArticleByDate - returns an article by date
func GetArticleByDate(date string) (Article, errors.DatabaseError) {
	row := database.Db.QueryRow(articleStatements["GetArticleByDate"], date)
	var ret Article

	switch err := row.Scan(
		&ret.ID, &ret.Title,
		&ret.ImageURL, &ret.Text,
		&ret.AuthorID, &ret.Date,
		&ret.AuthorName, &ret.EventName,
		&ret.LocLat, &ret.LocLong,
		&ret.LikesCount); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.GenericError{Message: "No article for the today ;(", HTTPCode: http.StatusNotFound}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// Save (insert) a new article
func (a *Article) Save() errors.DatabaseError {
	if _, err := database.Db.Exec(articleStatements["Save"],
		a.Title, a.ImageURL,
		a.Text, a.AuthorID,
		a.Date); err != nil {

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
