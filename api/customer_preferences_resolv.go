package api

type ResolverCustomerPreferences struct {
	s TypeCustomerPreferences
}

func (R *ResolverCustomerPreferences) Set(s TypeCustomerPreferences) {
	R.s = s
}

func (R ResolverCustomerPreferences) UserId() string {
	return R.s.UserId
}

func (R ResolverCustomerPreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R ResolverCustomerPreferences) PreferencesValue() *string {
	return R.s.PreferencesValue
}

func GenGqlTypeCustomerPreferences(extra string) string {
	return "type CustomerPreferences { " + extra + `
	UserId: String,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
