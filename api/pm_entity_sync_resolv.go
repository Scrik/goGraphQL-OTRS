package api

type ResolverPmEntitySync struct {
	s TypePmEntitySync
}

func (R *ResolverPmEntitySync) Set(s TypePmEntitySync) {
	R.s = s
}

func (R ResolverPmEntitySync) EntityType() string {
	return R.s.EntityType
}

func (R ResolverPmEntitySync) EntityId() *string {
	return R.s.EntityId
}

func (R ResolverPmEntitySync) SyncState() *string {
	return R.s.SyncState
}

func (R *ResolverPmEntitySync) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R *ResolverPmEntitySync) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func GenGqlTypePmEntitySync(extra string) string {
	return "type PmEntitySync { " + extra + `
	EntityType: String,
	EntityId: String,
	SyncState: String,
	CreateTime: String,
	ChangeTime: String,
	}`
}
