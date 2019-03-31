package api

import "time"

type TypeStandardTemplate struct {
	Id           int32      `db:"id"`
	Name         *string    `db:"name"`
	Text         *string    `db:"text"`
	ContentType  *string    `db:"content_type"`
	TemplateType *string    `db:"template_type"`
	Comments     *string    `db:"comments"`
	ValidId      *int8      `db:"valid_id"`
	CreateTime   *time.Time `db:"create_time"`
	CreateBy     *int32     `db:"create_by"`
	ChangeTime   *time.Time `db:"change_time"`
	ChangeBy     *int32     `db:"change_by"`
}
