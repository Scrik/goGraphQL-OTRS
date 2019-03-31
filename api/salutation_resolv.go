package api

type ResolverSalutation struct {
	s TypeSalutation
}

func (R *ResolverSalutation) Set(s TypeSalutation) {
	R.s = s
}

func (R ResolverSalutation) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSalutation) Name() *string {
	return R.s.Name
}

func (R ResolverSalutation) Text() *string {
	return R.s.Text
}

func (R ResolverSalutation) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverSalutation) Comments() *string {
	return R.s.Comments
}

func (R ResolverSalutation) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverSalutation) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSalutation) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSalutation) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSalutation) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSalutation(extra string) string {
	return "type Salutation { " + extra + `
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
