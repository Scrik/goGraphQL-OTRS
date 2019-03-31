package api

type ResolverSessions struct {
	s TypeSessions
}

func (R *ResolverSessions) Set(s TypeSessions) {
	R.s = s
}

func (R ResolverSessions) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSessions) SessionId() *string {
	return R.s.SessionId
}

func (R ResolverSessions) DataKey() *string {
	return R.s.DataKey
}

func (R ResolverSessions) DataValue() *string {
	return R.s.DataValue
}

func (R ResolverSessions) Serialized() *int32 {
	i := int32(*R.s.Serialized)
	return &i
}

func GenGqlTypeSessions(extra string) string {
	return "type Sessions { " + extra + `
	Id: String,
	SessionId: String,
	DataKey: String,
	DataValue: String,
	Serialized: Int,
	}`
}
