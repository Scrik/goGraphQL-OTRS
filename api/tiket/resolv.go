package tiket

import (
	"github.com/goGraphQL-OTRS/api/article"
	"github.com/goGraphQL-OTRS/internal/db"
)

// Articles Articles
func (R *Resolver) Articles() (res *[]*article.Resolver, err error) {
	rows, err := db.DB.Queryx("select * from `otrs`.`article` where `ticket_id`=?", R.s.Id)
	if err != nil {
		return
	}
	result := []*article.Resolver{}
	for rows.Next() {
		r := &article.Resolver{}
		a := article.TypeArticle{}
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
