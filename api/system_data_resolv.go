package api

type ResolverSystemData struct {
	s TypeSystemData
}

func (R *ResolverSystemData) Set(s TypeSystemData) {
	R.s = s
}

func (R ResolverSystemData) DataKey() string {
	return R.s.DataKey
}

func (R *ResolverSystemData) DataValue() *string {
	str := string(*R.s.DataValue)
	return &str
}

func (R *ResolverSystemData) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSystemData) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSystemData) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSystemData) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSystemData(extra string) string {
	return "type SystemData { " + extra + `
	DataKey: String,
	DataValue: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
