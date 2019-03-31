package api

type ResolverTicketType struct {
	s TypeTicketType
}

func (R *ResolverTicketType) Set(s TypeTicketType) {
	R.s = s
}

func (R ResolverTicketType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketType) Name() *string {
	return R.s.Name
}

func (R ResolverTicketType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverTicketType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketType(extra string) string {
	return "type TicketType { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
