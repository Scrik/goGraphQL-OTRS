package api

type ResolverSlaPreferences struct {
	s TypeSlaPreferences
}

func (R *ResolverSlaPreferences) Set(s TypeSlaPreferences) {
	R.s = s
}

func (R ResolverSlaPreferences) SlaId() int32 {
	return R.s.SlaId
}

func (R ResolverSlaPreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R ResolverSlaPreferences) PreferencesValue() *string {
	return R.s.PreferencesValue
}

func GenGqlTypeSlaPreferences(extra string) string {
	return "type SlaPreferences { " + extra + `
	SlaId: Int,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
