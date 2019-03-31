package api

type ResolverUserPreferences struct {
	s TypeUserPreferences
}

func (R *ResolverUserPreferences) Set(s TypeUserPreferences) {
	R.s = s
}

func (R ResolverUserPreferences) UserId() int32 {
	return R.s.UserId
}

func (R ResolverUserPreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R *ResolverUserPreferences) PreferencesValue() *string {
	str := string(*R.s.PreferencesValue)
	return &str
}

func GenGqlTypeUserPreferences(extra string) string {
	return "type UserPreferences { " + extra + `
	UserId: Int,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
