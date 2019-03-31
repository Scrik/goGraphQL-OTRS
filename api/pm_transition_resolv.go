package api

type ResolverPmTransition struct {
	s TypePmTransition
}

func (R *ResolverPmTransition) Set(s TypePmTransition) {
	R.s = s
}

func (R ResolverPmTransition) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPmTransition) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmTransition) Name() *string {
	return R.s.Name
}

func (R *ResolverPmTransition) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R *ResolverPmTransition) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPmTransition) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPmTransition) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPmTransition) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePmTransition(extra string) string {
	return "type PmTransition { " + extra + `
	Id: Int,
	EntityId: String,
	Name: String,
	Config: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
