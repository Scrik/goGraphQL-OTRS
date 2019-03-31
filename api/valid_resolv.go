package api

type ResolverValid struct {
	s TypeValid
}

func (R *ResolverValid) Set(s TypeValid) {
	R.s = s
}

func (R ResolverValid) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverValid) Name() *string {
	return R.s.Name
}

func (R *ResolverValid) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverValid) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverValid) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverValid) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeValid(extra string) string {
	return "type Valid { " + extra + `
	Id: Int,
	Name: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
