package api

type ResolverVirtualFsDb struct {
	s TypeVirtualFsDb
}

func (R *ResolverVirtualFsDb) Set(s TypeVirtualFsDb) {
	R.s = s
}

func (R ResolverVirtualFsDb) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverVirtualFsDb) Filename() *string {
	return R.s.Filename
}

func (R *ResolverVirtualFsDb) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R *ResolverVirtualFsDb) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeVirtualFsDb(extra string) string {
	return "type VirtualFsDb { " + extra + `
	Id: String,
	Filename: String,
	Content: String,
	CreateTime: String,
	}`
}
