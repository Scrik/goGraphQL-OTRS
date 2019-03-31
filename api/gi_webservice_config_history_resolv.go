package api

type ResolverGiWebserviceConfigHistory struct {
	s TypeGiWebserviceConfigHistory
}

func (R *ResolverGiWebserviceConfigHistory) Set(s TypeGiWebserviceConfigHistory) {
	R.s = s
}

func (R ResolverGiWebserviceConfigHistory) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverGiWebserviceConfigHistory) ConfigId() *int32 {
	return R.s.ConfigId
}

func (R *ResolverGiWebserviceConfigHistory) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R ResolverGiWebserviceConfigHistory) ConfigMd5() *string {
	return R.s.ConfigMd5
}

func (R *ResolverGiWebserviceConfigHistory) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGiWebserviceConfigHistory) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGiWebserviceConfigHistory) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGiWebserviceConfigHistory) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGiWebserviceConfigHistory(extra string) string {
	return "type GiWebserviceConfigHistory { " + extra + `
	Id: String,
	ConfigId: Int,
	Config: String,
	ConfigMd5: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
