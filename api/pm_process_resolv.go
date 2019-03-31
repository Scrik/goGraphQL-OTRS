package api

type ResolverPmProcess struct {
	s TypePmProcess
}

func (R *ResolverPmProcess) Set(s TypePmProcess) {
	R.s = s
}

func (R ResolverPmProcess) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPmProcess) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmProcess) Name() *string {
	return R.s.Name
}

func (R ResolverPmProcess) StateEntityId() *string {
	return R.s.StateEntityId
}

func (R *ResolverPmProcess) Layout() *string {
	str := string(*R.s.Layout)
	return &str
}

func (R *ResolverPmProcess) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R *ResolverPmProcess) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPmProcess) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPmProcess) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPmProcess) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePmProcess(extra string) string {
	return "type PmProcess { " + extra + `
	Id: Int,
	EntityId: String,
	Name: String,
	StateEntityId: String,
	Layout: String,
	Config: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
