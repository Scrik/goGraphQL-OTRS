package api

type ResolverVirtualFsPreferences struct {
	s TypeVirtualFsPreferences
}

func (R *ResolverVirtualFsPreferences) Set(s TypeVirtualFsPreferences) {
	R.s = s
}

func (R ResolverVirtualFsPreferences) VirtualFsId() string {
	str := fmt.Sprintf("%d", R.s.VirtualFsId)
	return &str
}

func (R ResolverVirtualFsPreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R ResolverVirtualFsPreferences) PreferencesValue() *string {
	return R.s.PreferencesValue
}

func GenGqlTypeVirtualFsPreferences(extra string) string {
	return "type VirtualFsPreferences { " + extra + `
	VirtualFsId: String,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
