package api

type ResolverTicketPriority struct {
	s TypeTicketPriority
}

func (R *ResolverTicketPriority) Set(s TypeTicketPriority) {
	R.s = s
}

func (R ResolverTicketPriority) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketPriority) Name() *string {
	return R.s.Name
}

func (R ResolverTicketPriority) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverTicketPriority) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketPriority) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketPriority) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketPriority) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketPriority(extra string) string {
	return "type TicketPriority { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
