package api

type ResolverServicePreferences struct {
	s TypeServicePreferences
}

func (R *ResolverServicePreferences) Set(s TypeServicePreferences) {
	R.s = s
}

func (R ResolverServicePreferences) ServiceId() int32 {
	return R.s.ServiceId
}

func (R ResolverServicePreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R ResolverServicePreferences) PreferencesValue() *string {
	return R.s.PreferencesValue
}

func GenGqlTypeServicePreferences(extra string) string {
	return "type ServicePreferences { " + extra + `
	ServiceId: Int,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
