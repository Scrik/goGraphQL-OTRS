package api

import "time"

type TypeArticleSearch struct {
	Id                  int64   `db:"id"`
	TicketId            *int64  `db:"ticket_id"`
	ArticleTypeId       *int8   `db:"article_type_id"`
	ArticleSenderTypeId *int8   `db:"article_sender_type_id"`
	AFrom               *string `db:"a_from"`
	ATo                 *string `db:"a_to"`
	ACc                 *string `db:"a_cc"`
	ASubject            *string `db:"a_subject"`
	ABody               *string `db:"a_body"`
	IncomingTime        *int32  `db:"incoming_time"`
}
