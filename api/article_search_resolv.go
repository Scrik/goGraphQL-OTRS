package api

type ResolverArticleSearch struct {
	s TypeArticleSearch
}

func (R *ResolverArticleSearch) Set(s TypeArticleSearch) {
	R.s = s
}

func (R ResolverArticleSearch) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverArticleSearch) TicketId() *string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R ResolverArticleSearch) ArticleTypeId() *int32 {
	i := int32(*R.s.ArticleTypeId)
	return &i
}

func (R ResolverArticleSearch) ArticleSenderTypeId() *int32 {
	i := int32(*R.s.ArticleSenderTypeId)
	return &i
}

func (R ResolverArticleSearch) AFrom() *string {
	return R.s.AFrom
}

func (R ResolverArticleSearch) ATo() *string {
	return R.s.ATo
}

func (R ResolverArticleSearch) ACc() *string {
	return R.s.ACc
}

func (R ResolverArticleSearch) ASubject() *string {
	return R.s.ASubject
}

func (R ResolverArticleSearch) ABody() *string {
	return R.s.ABody
}

func (R ResolverArticleSearch) IncomingTime() *int32 {
	return R.s.IncomingTime
}

func GenGqlTypeArticleSearch(extra string) string {
	return "type ArticleSearch { " + extra + `
	Id: String,
	TicketId: String,
	ArticleTypeId: Int,
	ArticleSenderTypeId: Int,
	AFrom: String,
	ATo: String,
	ACc: String,
	ASubject: String,
	ABody: String,
	IncomingTime: Int,
	}`
}
