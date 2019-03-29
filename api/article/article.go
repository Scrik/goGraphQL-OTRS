package article

import (
	"fmt"
	"time"

	"github.com/goGraphQL-OTRS/api/article"
	"github.com/goGraphQL-OTRS/api/types"
	"github.com/goGraphQL-OTRS/internal/db"
)

type TypeArticle struct {
	Id                  int64     `db:"id"`
	TicketId            int64     `db:"ticket_id"`
	ArticleTypeId       int8      `db:"article_type_id"`
	ArticleSenderTypeId int8      `db:"article_sender_type_id"`
	AFrom               string    `db:"a_from"`
	AReplyTo            *string   `db:"a_reply_to"`
	ATo                 string    `db:"a_to"`
	ACc                 *string   `db:"a_cc"`
	ASubject            string    `db:"a_subject"`
	AMessageId          string    `db:"a_message_id"`
	AMessageIdMd5       *string   `db:"a_message_id_md5"`
	AInReplyTo          *string   `db:"a_in_reply_to"`
	AReferences         *string   `db:"a_references"`
	AContentType        *string   `db:"a_content_type"`
	ABody               string    `db:"a_body"`
	IncomingTime        int32     `db:"incoming_time"`
	ContentPath         string    `db:"content_path"`
	ValidId             int8      `db:"valid_id"`
	CreateTime          time.Time `db:"create_time"`
	CreateBy            int32     `db:"create_by"`
	ChangeTime          time.Time `db:"change_time"`
	ChangeBy            int32     `db:"change_by"`
}

// Article Article
func One(ID string) (result *article.Resolver, err error) {
	result = &article.Resolver{}
	Article := types.TypeArticle{}
	row := db.DB.QueryRowx("select * from `otrs`.`article` where id=?", args.ID)
	err = row.StructScan(&Article)
	if err != nil {
		return nil, err
	}
	result.Set(Article)
	return
}

type Resolver struct {
	s TypeArticle
}

func (R *Resolver) Set(s TypeArticle) {
	R.s = s
}

func (R *Resolver) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *Resolver) TicketId() string {
	return fmt.Sprintf("%d", R.s.TicketId)
}

func (R *Resolver) ArticleTypeId() int32 {
	return int32(R.s.ArticleTypeId)
}

func (R *Resolver) ArticleSenderTypeId() int32 {
	return int32(R.s.ArticleSenderTypeId)
}

func (R *Resolver) AFrom() string {
	return R.s.AFrom
}

func (R *Resolver) AReplyTo() *string {
	return R.s.AReplyTo
}

func (R *Resolver) ATo() string {
	return R.s.ATo
}

func (R *Resolver) ACc() *string {
	return R.s.ACc
}

func (R *Resolver) ASubject() string {
	return R.s.ASubject
}

func (R *Resolver) AMessageId() string {
	return R.s.AMessageId
}

func (R *Resolver) AMessageIdMd5() *string {
	return R.s.AMessageIdMd5
}

func (R *Resolver) AInReplyTo() *string {
	return R.s.AInReplyTo
}

func (R *Resolver) AReferences() *string {
	return R.s.AReferences
}

func (R *Resolver) AContentType() *string {
	return R.s.AContentType
}

func (R *Resolver) ABody() string {
	return R.s.ABody
}

func (R *Resolver) IncomingTime() int32 {
	return R.s.IncomingTime
}

func (R *Resolver) ContentPath() string {
	return R.s.ContentPath
}

func (R *Resolver) ValidId() int32 {
	return int32(R.s.ValidId)
}

func (R *Resolver) CreateTime() string {
	return R.s.CreateTime.String()
}

func (R *Resolver) CreateBy() int32 {
	return R.s.CreateBy
}

func (R *Resolver) ChangeTime() string {
	return R.s.ChangeTime.String()
}

func (R *Resolver) ChangeBy() int32 {
	return R.s.ChangeBy
}

func GenGqlTypeArticle(extra string) string {
	return "type Article { " + extra + `
	Id: String!,
	TicketId: String!,
	ArticleTypeId: Int!,
	ArticleSenderTypeId: Int!,
	AFrom: String!,
	AReplyTo: String!,
	ATo: String!,
	ACc: String!,
	ASubject: String!,
	AMessageId: String!,
	AMessageIdMd5: String!,
	AInReplyTo: String!,
	AReferences: String!,
	AContentType: String!,
	ABody: String!,
	IncomingTime: Int!,
	ContentPath: String!,
	ValidId: Int!,
	CreateTime: String!,
	CreateBy: Int!,
	ChangeTime: String!,
	ChangeBy: Int!,
	}`
}
