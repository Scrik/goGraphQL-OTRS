package api

type ResolverStandardTemplate struct {
	s TypeStandardTemplate
}

func (R *ResolverStandardTemplate) Set(s TypeStandardTemplate) {
	R.s = s
}

func (R ResolverStandardTemplate) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverStandardTemplate) Name() *string {
	return R.s.Name
}

func (R ResolverStandardTemplate) Text() *string {
	return R.s.Text
}

func (R ResolverStandardTemplate) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverStandardTemplate) TemplateType() *string {
	return R.s.TemplateType
}

func (R ResolverStandardTemplate) Comments() *string {
	return R.s.Comments
}

func (R ResolverStandardTemplate) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverStandardTemplate) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverStandardTemplate) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverStandardTemplate) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverStandardTemplate) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeStandardTemplate(extra string) string {
	return "type StandardTemplate { " + extra + `
	Id: Int,
	Name: String,
	Text: String,
	ContentType: String,
	TemplateType: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
