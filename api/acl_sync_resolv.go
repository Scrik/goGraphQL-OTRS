package api

type ResolverAclSync struct {
	s TypeAclSync
}

func (R *ResolverAclSync) Set(s TypeAclSync) {
	R.s = s
}

func (R ResolverAclSync) AclId() string {
	return R.s.AclId
}

func (R ResolverAclSync) SyncState() *string {
	return R.s.SyncState
}

func (R *ResolverAclSync) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R *ResolverAclSync) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func GenGqlTypeAclSync(extra string) string {
	return "type AclSync { " + extra + `
	AclId: String,
	SyncState: String,
	CreateTime: String,
	ChangeTime: String,
	}`
}
