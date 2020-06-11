package models

import (
	"net/http"

	"github.com/VIVelev/goApod/database"

	"github.com/VIVelev/goApod/errors"
)

// Like struct
type Like struct {
	AuthorID  int `json:"authorId"`
	ArticleID int `json:"articleId"`
}

var likeStatements = map[string]string{
	"Save": `
		insert into likes (id, author_id, article_id)
		values (default, $1, $2)`,

	"DeleteByAuthorAndArticle": `
		delete from likes
		where author_id = $1 and article_id = $2`,
}

// Save - insert new like
func (l *Like) Save() errors.DatabaseError {
	_, err := database.Db.Exec(likeStatements["Save"], l.AuthorID, l.ArticleID)
	if err != nil {
		return &errors.InternalDatabaseError{Message: "Database error"}
	}

	return nil
}

// DeleteByAuthorAndArticle - delete like
func (l *Like) DeleteByAuthorAndArticle() errors.DatabaseError {
	result, err := database.Db.Exec(likeStatements["DeleteByAuthorAndArticle"])
	if err != nil {
		return &errors.InternalDatabaseError{Message: "Database error"}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: "Database error"}
	}
	if rowsAffected == 0 {
		return &errors.GenericError{HTTPCode: http.StatusNotFound, Message: "Invalid ids"}
	}

	return nil
}
