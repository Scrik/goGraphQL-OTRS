package api

type ResolverPmTransitionAction struct {
	s TypePmTransitionAction
}

func (R *ResolverPmTransitionAction) Set(s TypePmTransitionAction) {
	R.s = s
}

func (R ResolverPmTransitionAction) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPmTransitionAction) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmTransitionAction) Name() *string {
	return R.s.Name
}

func (R *ResolverPmTransitionAction) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R *ResolverPmTransitionAction) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPmTransitionAction) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPmTransitionAction) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPmTransitionAction) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePmTransitionAction(extra string) string {
	return "type PmTransitionAction { " + extra + `
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
