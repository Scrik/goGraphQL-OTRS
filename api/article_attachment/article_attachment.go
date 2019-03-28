package api

import (
	"fmt"
	"time"
)

type TypeArticleAttachment struct {
	Id                 int64     `json:"id"`
	ArticleId          int64     `json:"article_id"`
	Filename           string    `json:"filename"`
	ContentSize        string    `json:"content_size"`
	ContentType        string    `json:"content_type"`
	ContentId          string    `json:"content_id"`
	ContentAlternative string    `json:"content_alternative"`
	Disposition        string    `json:"disposition"`
	Content            []byte    `json:"content"`
	CreateTime         time.Time `json:"create_time"`
	CreateBy           int32     `json:"create_by"`
	ChangeTime         time.Time `json:"change_time"`
	ChangeBy           int32     `json:"change_by"`
}

type Resolver struct {
	s TypeArticleAttachment
}

func (R *Resolver) Set(s TypeArticleAttachment) {
	R.s = s
}

func (R *Resolver) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *Resolver) ArticleId() string {
	return fmt.Sprintf("%d", R.s.ArticleId)
}

func (R *Resolver) Filename() string {
	return R.s.Filename
}

func (R *Resolver) ContentSize() string {
	return R.s.ContentSize
}

func (R *Resolver) ContentType() string {
	return R.s.ContentType
}

func (R *Resolver) ContentId() string {
	return R.s.ContentId
}

func (R *Resolver) ContentAlternative() string {
	return R.s.ContentAlternative
}

func (R *Resolver) Disposition() string {
	return R.s.Disposition
}

func (R *Resolver) Content() string {
	return string(R.s.Content)
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

func GenGqlTypeArticleAttachment(extra string) string {
	return "type ArticleAttachment { " + extra + `
	Id: String!,
	ArticleId: String!,
	Filename: String!,
	ContentSize: String!,
	ContentType: String!,
	ContentId: String!,
	ContentAlternative: String!,
	Disposition: String!,
	Content: String!,
	CreateTime: String!,
	CreateBy: Int!,
	ChangeTime: String!,
	ChangeBy: Int!,
	}`
}
