package api

type ResolverSignature struct {
	s TypeSignature
}

func (R *ResolverSignature) Set(s TypeSignature) {
	R.s = s
}

func (R ResolverSignature) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSignature) Name() *string {
	return R.s.Name
}

func (R ResolverSignature) Text() *string {
	return R.s.Text
}

func (R ResolverSignature) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverSignature) Comments() *string {
	return R.s.Comments
}

func (R ResolverSignature) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverSignature) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSignature) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSignature) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSignature) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSignature(extra string) string {
	return "type Signature { " + extra + `
	Id: Int,
	Name: String,
	Text: String,
	ContentType: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
