package api

import (
	"github.com/goGraphQL-OTRS/internal/db"
)

// Articles Articles
func (R *ResolverTicket) Articles() (res *[]*ResolverArticle, err error) {
	rows, err := db.DB.Queryx("select * from `otrs`.`article` where `ticket_id`=?", R.s.Id)
	if err != nil {
		return
	}
	result := []*ResolverArticle{}
	for rows.Next() {
		r := &ResolverArticle{}
		a := TypeArticle{}
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
