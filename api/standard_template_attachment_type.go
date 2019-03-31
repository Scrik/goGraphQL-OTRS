package api

import "time"

type TypeStandardTemplateAttachment struct {
	Id                   int32      `db:"id"`
	StandardAttachmentId *int32     `db:"standard_attachment_id"`
	StandardTemplateId   *int32     `db:"standard_template_id"`
	CreateTime           *time.Time `db:"create_time"`
	CreateBy             *int32     `db:"create_by"`
	ChangeTime           *time.Time `db:"change_time"`
	ChangeBy             *int32     `db:"change_by"`
}
