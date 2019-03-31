package api

type ResolverPmActivityDialog struct {
	s TypePmActivityDialog
}

func (R *ResolverPmActivityDialog) Set(s TypePmActivityDialog) {
	R.s = s
}

func (R ResolverPmActivityDialog) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPmActivityDialog) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmActivityDialog) Name() *string {
	return R.s.Name
}

func (R *ResolverPmActivityDialog) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R *ResolverPmActivityDialog) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPmActivityDialog) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPmActivityDialog) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPmActivityDialog) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePmActivityDialog(extra string) string {
	return "type PmActivityDialog { " + extra + `
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
