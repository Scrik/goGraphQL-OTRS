package api

type ResolverTimeAccounting struct {
	s TypeTimeAccounting
}

func (R *ResolverTimeAccounting) Set(s TypeTimeAccounting) {
	R.s = s
}

func (R ResolverTimeAccounting) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverTimeAccounting) TicketId() *string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R *ResolverTimeAccounting) ArticleId() *string {
	str := fmt.Sprintf("%d", R.s.ArticleId)
	return &str
}

func (R ResolverTimeAccounting) TimeUnit() *string {
	return R.s.TimeUnit
}

func (R *ResolverTimeAccounting) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTimeAccounting) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTimeAccounting) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTimeAccounting) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTimeAccounting(extra string) string {
	return "type TimeAccounting { " + extra + `
	Id: String,
	TicketId: String,
	ArticleId: String,
	TimeUnit: decimal(10,2),
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
