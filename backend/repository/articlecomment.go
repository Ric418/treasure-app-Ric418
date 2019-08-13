package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllArticlecomment(db *sqlx.DB) ([]model.Articlecomment, error) {
	a := make([]model.Articlecomment, 0)
	if err := db.Select(&a, `SELECT id, title, body FROM article_comment`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindArticlecomment(db *sqlx.DB, id int64) (*model.Articlecomment, error) {
	a := model.Articlecomment{}
	if err := db.Get(&a, `
SELECT id, title, body FROM article_comment WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateArticlecomment(db *sqlx.Tx, a *model.Articlecomment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_comment (id, user_id, body, article_id) VALUES (?, ?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.ID, a.UserID, a.Body, a.ArticleID)
}

func UpdateArticlecomment(db *sqlx.Tx, id int64, a *model.Articlecomment) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article_comment SET body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Body, id)
}

func DestroyArticlecomment(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article_comment WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

