package api

type ResolverPmActivity struct {
	s TypePmActivity
}

func (R *ResolverPmActivity) Set(s TypePmActivity) {
	R.s = s
}

func (R ResolverPmActivity) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPmActivity) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmActivity) Name() *string {
	return R.s.Name
}

func (R *ResolverPmActivity) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R *ResolverPmActivity) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPmActivity) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPmActivity) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPmActivity) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePmActivity(extra string) string {
	return "type PmActivity { " + extra + `
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
