package api

type ResolverLinkObject struct {
	s TypeLinkObject
}

func (R *ResolverLinkObject) Set(s TypeLinkObject) {
	R.s = s
}

func (R ResolverLinkObject) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverLinkObject) Name() *string {
	return R.s.Name
}

func GenGqlTypeLinkObject(extra string) string {
	return "type LinkObject { " + extra + `
	Id: Int,
	Name: String,
	}`
}
