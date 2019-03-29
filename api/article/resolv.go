package article

import (
	"github.com/goGraphQL-OTRS/api/article"
	"github.com/goGraphQL-OTRS/api/article_attachment"
	"github.com/goGraphQL-OTRS/internal/db"
)

// Articles Articles
func (R *Resolver) Attachment() (res *[]*article.Resolver, err error) {
	rows, err := db.DB.Queryx("select * from `otrs`.`article_attachment` where `article_id`=?", R.s.Id)
	if err != nil {
		return
	}
	result := []*article_attachment.Resolver{}
	for rows.Next() {
		r := &article_attachment.Resolver{}
		a := article_attachment.TypeArticleAttachment{}
		err = rows.StructScan(&a)
		if err != nil {
			return
		}
		r.Set(a)
		result = append(result, r)
	}
	res = &result
	return
}
