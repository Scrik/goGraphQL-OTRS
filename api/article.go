package api

import "github.com/goGraphQL-OTRS/internal/db"

// Attachment Attachment
func (R ResolverArticle) Attachments() (res *[]*ResolverArticleAttachment, err error) {
	rows, err := db.DB.Queryx("select * from `otrs`.`article_attachment` where `article_id`=?", R.s.Id)
	if err != nil {
		return
	}
	result := []*ResolverArticleAttachment{}
	for rows.Next() {
		r := &ResolverArticleAttachment{}
		a := TypeArticleAttachment{}
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
