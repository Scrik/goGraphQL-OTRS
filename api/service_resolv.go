package api

type ResolverService struct {
	s TypeService
}

func (R *ResolverService) Set(s TypeService) {
	R.s = s
}

func (R ResolverService) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverService) Name() *string {
	return R.s.Name
}

func (R ResolverService) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R ResolverService) Comments() *string {
	return R.s.Comments
}

func (R *ResolverService) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverService) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverService) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverService) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeService(extra string) string {
	return "type Service { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	Comments: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
