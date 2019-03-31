package api

type ResolverStandardAttachment struct {
	s TypeStandardAttachment
}

func (R *ResolverStandardAttachment) Set(s TypeStandardAttachment) {
	R.s = s
}

func (R ResolverStandardAttachment) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverStandardAttachment) Name() *string {
	return R.s.Name
}

func (R ResolverStandardAttachment) ContentType() *string {
	return R.s.ContentType
}

func (R *ResolverStandardAttachment) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R ResolverStandardAttachment) Filename() *string {
	return R.s.Filename
}

func (R ResolverStandardAttachment) Comments() *string {
	return R.s.Comments
}

func (R ResolverStandardAttachment) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverStandardAttachment) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverStandardAttachment) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverStandardAttachment) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverStandardAttachment) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeStandardAttachment(extra string) string {
	return "type StandardAttachment { " + extra + `
	Id: Int,
	Name: String,
	ContentType: String,
	Content: String,
	Filename: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
