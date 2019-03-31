package api

type ResolverGiDebuggerEntryContent struct {
	s TypeGiDebuggerEntryContent
}

func (R *ResolverGiDebuggerEntryContent) Set(s TypeGiDebuggerEntryContent) {
	R.s = s
}

func (R ResolverGiDebuggerEntryContent) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverGiDebuggerEntryContent) GiDebuggerEntryId() *string {
	str := fmt.Sprintf("%d", R.s.GiDebuggerEntryId)
	return &str
}

func (R ResolverGiDebuggerEntryContent) DebugLevel() *string {
	return R.s.DebugLevel
}

func (R ResolverGiDebuggerEntryContent) Subject() *string {
	return R.s.Subject
}

func (R *ResolverGiDebuggerEntryContent) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R *ResolverGiDebuggerEntryContent) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeGiDebuggerEntryContent(extra string) string {
	return "type GiDebuggerEntryContent { " + extra + `
	Id: String,
	GiDebuggerEntryId: String,
	DebugLevel: String,
	Subject: String,
	Content: String,
	CreateTime: String,
	}`
}
