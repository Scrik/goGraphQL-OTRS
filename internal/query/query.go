package query

import (
	"github.com/goGraphQL-OTRS/api/article"
	"github.com/goGraphQL-OTRS/api/article_attachment"
	"github.com/goGraphQL-OTRS/api/tiket"
)

// Query Query
type Query struct{}

// Hello Hello
func (Q *Query) Hello() string {
	return "Hello, world!"
}

// Tiket Tiket
func (Q *Query) Tiket(args struct{ ID string }) (result *tiket.Resolver, err error) {
	return tiket.One(args.ID)
}

// Article Article
func (Q *Query) Article(args struct{ ID string }) (result *article.Resolver, err error) {
	return article.One(arg.ID)
}

// ArticleAttachment Attachment
func (Q *Query) ArticleAttachment(args struct{ ID string }) (result *article_attachment.Resolver, err error) {
	return article_attachment.One(args.ID)
}
