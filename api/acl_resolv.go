package api

type ResolverAcl struct {
	s TypeAcl
}

func (R *ResolverAcl) Set(s TypeAcl) {
	R.s = s
}

func (R ResolverAcl) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverAcl) Name() *string {
	return R.s.Name
}

func (R ResolverAcl) Comments() *string {
	return R.s.Comments
}

func (R ResolverAcl) Description() *string {
	return R.s.Description
}

func (R ResolverAcl) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R ResolverAcl) StopAfterMatch() *int32 {
	i := int32(*R.s.StopAfterMatch)
	return &i
}

func (R *ResolverAcl) ConfigMatch() *string {
	str := string(*R.s.ConfigMatch)
	return &str
}

func (R *ResolverAcl) ConfigChange() *string {
	str := string(*R.s.ConfigChange)
	return &str
}

func (R *ResolverAcl) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverAcl) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverAcl) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverAcl) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeAcl(extra string) string {
	return "type Acl { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	Description: String,
	ValidId: Int,
	StopAfterMatch: Int,
	ConfigMatch: String,
	ConfigChange: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
