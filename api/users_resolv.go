package api

type ResolverUsers struct {
	s TypeUsers
}

func (R *ResolverUsers) Set(s TypeUsers) {
	R.s = s
}

func (R ResolverUsers) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverUsers) Login() *string {
	return R.s.Login
}

func (R ResolverUsers) Pw() *string {
	return R.s.Pw
}

func (R ResolverUsers) Title() *string {
	return R.s.Title
}

func (R ResolverUsers) FirstName() *string {
	return R.s.FirstName
}

func (R ResolverUsers) LastName() *string {
	return R.s.LastName
}

func (R ResolverUsers) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverUsers) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverUsers) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverUsers) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverUsers) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeUsers(extra string) string {
	return "type Users { " + extra + `
	Id: Int,
	Login: String,
	Pw: String,
	Title: String,
	FirstName: String,
	LastName: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
