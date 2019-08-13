package model

type Articlecomment struct {
	ID    int64  `db:"id" json:"id"`
	ArticleID int64 `db:"article_id" json:"article_id"`
	Body  string `db:"body" json:"commentbody"`
	UserID int64 `db:"user_id" json:"user_id"`
}
