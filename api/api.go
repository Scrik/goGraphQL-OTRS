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

//     Articles: [Article],
/**
article(ID: String!): Article!
    articleAttachment(ID: String!): ArticleAttachment!


	Attachments: [ArticleAttachment],



{
  tiket(ID: "1") {
    Title
    CreateBy
    Tn
    Articles {
      AFrom
      ATo
      ASubject
      ABody
    }
  }
  article(ID: "1") {
    Id
    AFrom
    ATo
    AReplyTo
    ACc
    ASubject
  }
  articleAttachment(ID:"1"){
    Id
    ArticleId
  }
}


*/
