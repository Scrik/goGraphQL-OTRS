package article_attachment

import (
	"fmt"
	"time"

	"github.com/goGraphQL-OTRS/api/article_attachment"
	"github.com/goGraphQL-OTRS/api/types"
	"github.com/goGraphQL-OTRS/internal/db"
)

type TypeArticleAttachment struct {
	Id                 int64     `db:"id"`
	ArticleId          int64     `db:"article_id"`
	Filename           string    `db:"filename"`
	ContentSize        string    `db:"content_size"`
	ContentType        string    `db:"content_type"`
	ContentId          *string   `db:"content_id"`
	ContentAlternative *string   `db:"content_alternative"`
	Disposition        string    `db:"disposition"`
	Content            []byte    `db:"content"`
	CreateTime         time.Time `db:"create_time"`
	CreateBy           int32     `db:"create_by"`
	ChangeTime         time.Time `db:"change_time"`
	ChangeBy           int32     `db:"change_by"`
}

// Article Article
func One(ID string) (result *article_attachment.Resolver, err error) {
	result = &article_attachment.Resolver{}
	ArticleAttachment := types.TypeArticleAttachment{}
	row := db.DB.QueryRowx("select * from `otrs`.`article_attachment` where id=?", args.ID)
	err = row.StructScan(&ArticleAttachment)
	if err != nil {
		return nil, err
	}
	result.Set(ArticleAttachment)
	return
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

func (R *Resolver) ContentId() *string {
	return R.s.ContentId
}

func (R *Resolver) ContentAlternative() *string {
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
