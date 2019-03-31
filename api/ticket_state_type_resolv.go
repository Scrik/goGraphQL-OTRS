package api

type ResolverTicketStateType struct {
	s TypeTicketStateType
}

func (R *ResolverTicketStateType) Set(s TypeTicketStateType) {
	R.s = s
}

func (R ResolverTicketStateType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketStateType) Name() *string {
	return R.s.Name
}

func (R ResolverTicketStateType) Comments() *string {
	return R.s.Comments
}

func (R *ResolverTicketStateType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketStateType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketStateType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketStateType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketStateType(extra string) string {
	return "type TicketStateType { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
