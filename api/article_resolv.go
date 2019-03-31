package api

type ResolverArticle struct {
	s TypeArticle
}

func (R *ResolverArticle) Set(s TypeArticle) {
	R.s = s
}

func (R ResolverArticle) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverArticle) TicketId() *string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R ResolverArticle) ArticleTypeId() *int32 {
	i := int32(*R.s.ArticleTypeId)
	return &i
}

func (R ResolverArticle) ArticleSenderTypeId() *int32 {
	i := int32(*R.s.ArticleSenderTypeId)
	return &i
}

func (R ResolverArticle) AFrom() *string {
	return R.s.AFrom
}

func (R ResolverArticle) AReplyTo() *string {
	return R.s.AReplyTo
}

func (R ResolverArticle) ATo() *string {
	return R.s.ATo
}

func (R ResolverArticle) ACc() *string {
	return R.s.ACc
}

func (R ResolverArticle) ASubject() *string {
	return R.s.ASubject
}

func (R ResolverArticle) AMessageId() *string {
	return R.s.AMessageId
}

func (R ResolverArticle) AMessageIdMd5() *string {
	return R.s.AMessageIdMd5
}

func (R ResolverArticle) AInReplyTo() *string {
	return R.s.AInReplyTo
}

func (R ResolverArticle) AReferences() *string {
	return R.s.AReferences
}

func (R ResolverArticle) AContentType() *string {
	return R.s.AContentType
}

func (R ResolverArticle) ABody() *string {
	return R.s.ABody
}

func (R ResolverArticle) IncomingTime() *int32 {
	return R.s.IncomingTime
}

func (R ResolverArticle) ContentPath() *string {
	return R.s.ContentPath
}

func (R ResolverArticle) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverArticle) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticle) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverArticle) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverArticle) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeArticle(extra string) string {
	return "type Article { " + extra + `
	Id: String,
	TicketId: String,
	ArticleTypeId: Int,
	ArticleSenderTypeId: Int,
	AFrom: String,
	AReplyTo: String,
	ATo: String,
	ACc: String,
	ASubject: String,
	AMessageId: String,
	AMessageIdMd5: String,
	AInReplyTo: String,
	AReferences: String,
	AContentType: String,
	ABody: String,
	IncomingTime: Int,
	ContentPath: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
