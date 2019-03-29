package api

import "time"

type TypeArticle struct {
	Id                  int64      `db:"id"`
	TicketId            *int64     `db:"ticket_id"`
	ArticleTypeId       *int8      `db:"article_type_id"`
	ArticleSenderTypeId *int8      `db:"article_sender_type_id"`
	AFrom               *string    `db:"a_from"`
	AReplyTo            *string    `db:"a_reply_to"`
	ATo                 *string    `db:"a_to"`
	ACc                 *string    `db:"a_cc"`
	ASubject            *string    `db:"a_subject"`
	AMessageId          *string    `db:"a_message_id"`
	AMessageIdMd5       *string    `db:"a_message_id_md5"`
	AInReplyTo          *string    `db:"a_in_reply_to"`
	AReferences         *string    `db:"a_references"`
	AContentType        *string    `db:"a_content_type"`
	ABody               *string    `db:"a_body"`
	IncomingTime        *int32     `db:"incoming_time"`
	ContentPath         *string    `db:"content_path"`
	ValidId             *int8      `db:"valid_id"`
	CreateTime          *time.Time `db:"create_time"`
	CreateBy            *int32     `db:"create_by"`
	ChangeTime          *time.Time `db:"change_time"`
	ChangeBy            *int32     `db:"change_by"`
}
