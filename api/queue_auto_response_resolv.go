package api

type ResolverQueueAutoResponse struct {
	s TypeQueueAutoResponse
}

func (R *ResolverQueueAutoResponse) Set(s TypeQueueAutoResponse) {
	R.s = s
}

func (R ResolverQueueAutoResponse) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverQueueAutoResponse) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverQueueAutoResponse) AutoResponseId() *int32 {
	return R.s.AutoResponseId
}

func (R *ResolverQueueAutoResponse) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverQueueAutoResponse) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverQueueAutoResponse) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverQueueAutoResponse) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeQueueAutoResponse(extra string) string {
	return "type QueueAutoResponse { " + extra + `
	Id: Int,
	QueueId: Int,
	AutoResponseId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
