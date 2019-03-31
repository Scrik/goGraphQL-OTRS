package api

type ResolverTicketHistory struct {
	s TypeTicketHistory
}

func (R *ResolverTicketHistory) Set(s TypeTicketHistory) {
	R.s = s
}

func (R ResolverTicketHistory) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketHistory) Name() *string {
	return R.s.Name
}

func (R ResolverTicketHistory) HistoryTypeId() *int32 {
	i := int32(*R.s.HistoryTypeId)
	return &i
}

func (R *ResolverTicketHistory) TicketId() *string {
	str := fmt.Sprintf("%d", R.s.TicketId)
	return &str
}

func (R *ResolverTicketHistory) ArticleId() *string {
	str := fmt.Sprintf("%d", R.s.ArticleId)
	return &str
}

func (R ResolverTicketHistory) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R ResolverTicketHistory) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverTicketHistory) OwnerId() *int32 {
	return R.s.OwnerId
}

func (R ResolverTicketHistory) PriorityId() *int32 {
	i := int32(*R.s.PriorityId)
	return &i
}

func (R ResolverTicketHistory) StateId() *int32 {
	i := int32(*R.s.StateId)
	return &i
}

func (R *ResolverTicketHistory) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketHistory) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketHistory) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketHistory) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketHistory(extra string) string {
	return "type TicketHistory { " + extra + `
	Id: String,
	Name: String,
	HistoryTypeId: Int,
	TicketId: String,
	ArticleId: String,
	TypeId: Int,
	QueueId: Int,
	OwnerId: Int,
	PriorityId: Int,
	StateId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
