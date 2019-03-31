package api

type ResolverAutoResponse struct {
	s TypeAutoResponse
}

func (R *ResolverAutoResponse) Set(s TypeAutoResponse) {
	R.s = s
}

func (R ResolverAutoResponse) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverAutoResponse) Name() *string {
	return R.s.Name
}

func (R ResolverAutoResponse) Text0() *string {
	return R.s.Text0
}

func (R ResolverAutoResponse) Text1() *string {
	return R.s.Text1
}

func (R ResolverAutoResponse) TypeId() *int32 {
	i := int32(*R.s.TypeId)
	return &i
}

func (R ResolverAutoResponse) SystemAddressId() *int32 {
	i := int32(*R.s.SystemAddressId)
	return &i
}

func (R ResolverAutoResponse) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverAutoResponse) Comments() *string {
	return R.s.Comments
}

func (R ResolverAutoResponse) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverAutoResponse) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverAutoResponse) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverAutoResponse) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverAutoResponse) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeAutoResponse(extra string) string {
	return "type AutoResponse { " + extra + `
	Id: Int,
	Name: String,
	Text0: String,
	Text1: String,
	TypeId: Int,
	SystemAddressId: Int,
	ContentType: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
