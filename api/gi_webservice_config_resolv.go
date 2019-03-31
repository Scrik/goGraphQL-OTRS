package api

type ResolverGiWebserviceConfig struct {
	s TypeGiWebserviceConfig
}

func (R *ResolverGiWebserviceConfig) Set(s TypeGiWebserviceConfig) {
	R.s = s
}

func (R ResolverGiWebserviceConfig) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverGiWebserviceConfig) Name() *string {
	return R.s.Name
}

func (R *ResolverGiWebserviceConfig) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R ResolverGiWebserviceConfig) ConfigMd5() *string {
	return R.s.ConfigMd5
}

func (R ResolverGiWebserviceConfig) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverGiWebserviceConfig) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverGiWebserviceConfig) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverGiWebserviceConfig) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverGiWebserviceConfig) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeGiWebserviceConfig(extra string) string {
	return "type GiWebserviceConfig { " + extra + `
	Id: Int,
	Name: String,
	Config: String,
	ConfigMd5: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
