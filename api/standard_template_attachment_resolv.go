package api

type ResolverStandardTemplateAttachment struct {
	s TypeStandardTemplateAttachment
}

func (R *ResolverStandardTemplateAttachment) Set(s TypeStandardTemplateAttachment) {
	R.s = s
}

func (R ResolverStandardTemplateAttachment) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverStandardTemplateAttachment) StandardAttachmentId() *int32 {
	return R.s.StandardAttachmentId
}

func (R ResolverStandardTemplateAttachment) StandardTemplateId() *int32 {
	return R.s.StandardTemplateId
}

func (R *ResolverStandardTemplateAttachment) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverStandardTemplateAttachment) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverStandardTemplateAttachment) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverStandardTemplateAttachment) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeStandardTemplateAttachment(extra string) string {
	return "type StandardTemplateAttachment { " + extra + `
	Id: Int,
	StandardAttachmentId: Int,
	StandardTemplateId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
