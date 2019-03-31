package api

type ResolverTicketHistoryType struct {
	s TypeTicketHistoryType
}

func (R *ResolverTicketHistoryType) Set(s TypeTicketHistoryType) {
	R.s = s
}

func (R ResolverTicketHistoryType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketHistoryType) Name() *string {
	return R.s.Name
}

func (R ResolverTicketHistoryType) Comments() *string {
	return R.s.Comments
}

func (R ResolverTicketHistoryType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverTicketHistoryType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketHistoryType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketHistoryType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketHistoryType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketHistoryType(extra string) string {
	return "type TicketHistoryType { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
