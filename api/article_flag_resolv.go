package api

type ResolverArticleFlag struct {
	s TypeArticleFlag
}

func (R *ResolverArticleFlag) Set(s TypeArticleFlag) {
	R.s = s
}

func (R ResolverArticleFlag) ArticleId() string {
	str := fmt.Sprintf("%d", R.s.ArticleId)
	return &str
}

func (R ResolverArticleFlag) ArticleKey() *string {
	return R.s.ArticleKey
}

func (R ResolverArticleFlag) ArticleValue() *string {
	return R.s.ArticleValue
}

func (R *ResolverArticleFlag) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticleFlag) CreateBy() *int32 {
	return R.s.CreateBy
}

func GenGqlTypeArticleFlag(extra string) string {
	return "type ArticleFlag { " + extra + `
	ArticleId: String,
	ArticleKey: String,
	ArticleValue: String,
	CreateTime: String,
	CreateBy: Int,
	}`
}
