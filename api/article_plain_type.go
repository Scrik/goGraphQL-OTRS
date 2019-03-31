package api

import "time"

type TypeArticlePlain struct {
	Id         int64      `db:"id"`
	ArticleId  *int64     `db:"article_id"`
	Body       *[]byte    `db:"body"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
