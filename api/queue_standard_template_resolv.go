package api

type ResolverQueueStandardTemplate struct {
	s TypeQueueStandardTemplate
}

func (R *ResolverQueueStandardTemplate) Set(s TypeQueueStandardTemplate) {
	R.s = s
}

func (R ResolverQueueStandardTemplate) QueueId() int32 {
	return R.s.QueueId
}

func (R ResolverQueueStandardTemplate) StandardTemplateId() *int32 {
	return R.s.StandardTemplateId
}

func (R *ResolverQueueStandardTemplate) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverQueueStandardTemplate) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverQueueStandardTemplate) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverQueueStandardTemplate) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeQueueStandardTemplate(extra string) string {
	return "type QueueStandardTemplate { " + extra + `
	QueueId: Int,
	StandardTemplateId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
