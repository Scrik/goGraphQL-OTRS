package api

import "time"

type TypeArticleAttachment struct {
	Id                 int64      `db:"id"`
	ArticleId          *int64     `db:"article_id"`
	Filename           *string    `db:"filename"`
	ContentSize        *string    `db:"content_size"`
	ContentType        *string    `db:"content_type"`
	ContentId          *string    `db:"content_id"`
	ContentAlternative *string    `db:"content_alternative"`
	Disposition        *string    `db:"disposition"`
	Content            *[]byte    `db:"content"`
	CreateTime         *time.Time `db:"create_time"`
	CreateBy           *int32     `db:"create_by"`
	ChangeTime         *time.Time `db:"change_time"`
	ChangeBy           *int32     `db:"change_by"`
}
