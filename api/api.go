package api

import "github.com/goGraphQL-OTRS/internal/db"

// OneTiket OneTiket
func OneTicket(ID string) (result *ResolverTicket, err error) {
	result = &ResolverTicket{}
	S := TypeTicket{}
	row := db.DB.QueryRowx("select * from `otrs`.`ticket` where id=?", ID)
	err = row.StructScan(&S)
	if err != nil {
		return nil, err
	}
	result.Set(S)
	return
}

// OneTiket OneTiket
func OneArticle(ID string) (result *ResolverArticle, err error) {
	result = &ResolverArticle{}
	S := TypeArticle{}
	row := db.DB.QueryRowx("select * from `otrs`.`article` where id=?", ID)
	err = row.StructScan(&S)
	if err != nil {
		return nil, err
	}
	result.Set(S)
	return
}

// OneTiket OneTiket
func OneAttachment(ID string) (result *ResolverArticleAttachment, err error) {
	result = &ResolverArticleAttachment{}
	S := TypeArticleAttachment{}
	row := db.DB.QueryRowx("select * from `otrs`.`article_attachment` where id=?", ID)
	err = row.StructScan(&S)
	if err != nil {
		return nil, err
	}
	result.Set(S)
	return
}

// OneTiket OneTiket
func OneCustomerCompany(ID string) (result *ResolverCustomerCompany, err error) {
	result = &ResolverCustomerCompany{}
	S := TypeCustomerCompany{}
	row := db.DB.QueryRowx("select * from `otrs`.`customer_company` where customer_id=?", ID)
	err = row.StructScan(&S)
	if err != nil {
		return nil, err
	}
	result.Set(S)
	return
}
