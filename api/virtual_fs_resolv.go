package api

type ResolverVirtualFs struct {
	s TypeVirtualFs
}

func (R *ResolverVirtualFs) Set(s TypeVirtualFs) {
	R.s = s
}

func (R ResolverVirtualFs) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverVirtualFs) Filename() *string {
	return R.s.Filename
}

func (R ResolverVirtualFs) Backend() *string {
	return R.s.Backend
}

func (R ResolverVirtualFs) BackendKey() *string {
	return R.s.BackendKey
}

func (R *ResolverVirtualFs) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeVirtualFs(extra string) string {
	return "type VirtualFs { " + extra + `
	Id: String,
	Filename: String,
	Backend: String,
	BackendKey: String,
	CreateTime: String,
	}`
}
