package api

type ResolverSearchProfile struct {
	s TypeSearchProfile
}

func (R *ResolverSearchProfile) Set(s TypeSearchProfile) {
	R.s = s
}

func (R ResolverSearchProfile) Login() string {
	return R.s.Login
}

func (R ResolverSearchProfile) ProfileName() *string {
	return R.s.ProfileName
}

func (R ResolverSearchProfile) ProfileType() *string {
	return R.s.ProfileType
}

func (R ResolverSearchProfile) ProfileKey() *string {
	return R.s.ProfileKey
}

func (R ResolverSearchProfile) ProfileValue() *string {
	return R.s.ProfileValue
}

func GenGqlTypeSearchProfile(extra string) string {
	return "type SearchProfile { " + extra + `
	Login: String,
	ProfileName: String,
	ProfileType: String,
	ProfileKey: String,
	ProfileValue: String,
	}`
}
