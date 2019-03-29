package query

import (
	"github.com/goGraphQL-OTRS/api"
)

// Query Query
type Query struct{}

// Hello Hello
func (Q *Query) Hello() string {
	return "Hello, world!"
}

// Tiket Tiket
func (Q *Query) Tiket(args struct{ ID string }) (result *api.ResolverTicket, err error) {
	return api.OneTicket(args.ID)
}

// Article Article
func (Q *Query) Article(args struct{ ID string }) (result *api.ResolverArticle, err error) {
	return api.OneArticle(args.ID)
}

// // ArticleAttachment Attachment
func (Q *Query) ArticleAttachment(args struct{ ID string }) (result *api.ResolverArticleAttachment, err error) {
	return api.OneAttachment(args.ID)
}
