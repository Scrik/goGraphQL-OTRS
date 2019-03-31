package api

type ResolverFollowUpPossible struct {
	s TypeFollowUpPossible
}

func (R *ResolverFollowUpPossible) Set(s TypeFollowUpPossible) {
	R.s = s
}

func (R ResolverFollowUpPossible) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverFollowUpPossible) Name() *string {
	return R.s.Name
}

func (R ResolverFollowUpPossible) Comments() *string {
	return R.s.Comments
}

func (R ResolverFollowUpPossible) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverFollowUpPossible) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverFollowUpPossible) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverFollowUpPossible) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverFollowUpPossible) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeFollowUpPossible(extra string) string {
	return "type FollowUpPossible { " + extra + `
	Id: Int,
	Name: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
