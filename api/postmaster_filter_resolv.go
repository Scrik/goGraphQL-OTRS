package api

type ResolverPostmasterFilter struct {
	s TypePostmasterFilter
}

func (R *ResolverPostmasterFilter) Set(s TypePostmasterFilter) {
	R.s = s
}

func (R ResolverPostmasterFilter) FName() string {
	return R.s.FName
}

func (R ResolverPostmasterFilter) FStop() *int32 {
	i := int32(*R.s.FStop)
	return &i
}

func (R ResolverPostmasterFilter) FType() *string {
	return R.s.FType
}

func (R ResolverPostmasterFilter) FKey() *string {
	return R.s.FKey
}

func (R ResolverPostmasterFilter) FValue() *string {
	return R.s.FValue
}

func (R ResolverPostmasterFilter) FNot() *int32 {
	i := int32(*R.s.FNot)
	return &i
}

func GenGqlTypePostmasterFilter(extra string) string {
	return "type PostmasterFilter { " + extra + `
	FName: String,
	FStop: Int,
	FType: String,
	FKey: String,
	FValue: String,
	FNot: Int,
	}`
}
