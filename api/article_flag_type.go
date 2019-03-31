package api

import "time"

type TypeArticleFlag struct {
	ArticleId    int64      `db:"article_id"`
	ArticleKey   *string    `db:"article_key"`
	ArticleValue *string    `db:"article_value"`
	CreateTime   *time.Time `db:"create_time"`
	CreateBy     *int32     `db:"create_by"`
}
