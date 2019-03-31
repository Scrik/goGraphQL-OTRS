package api

type ResolverTicketState struct {
	s TypeTicketState
}

func (R *ResolverTicketState) Set(s TypeTicketState) {
	R.s = s
}

func (R ResolverTicketState) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketState) Name() *string {
	return R.s.Name
}

func (R ResolverTicketState) Comments() *string {
	return R.s.Comments
}

func (R ResolverTicketState) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R ResolverTicketState) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverTicketState) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketState) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketState) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketState) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketState(extra string) string {
	return "type TicketState { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	TypeId: Int,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
