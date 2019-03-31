package api

type ResolverCloudServiceConfig struct {
	s TypeCloudServiceConfig
}

func (R *ResolverCloudServiceConfig) Set(s TypeCloudServiceConfig) {
	R.s = s
}

func (R ResolverCloudServiceConfig) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverCloudServiceConfig) Name() *string {
	return R.s.Name
}

func (R *ResolverCloudServiceConfig) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R ResolverCloudServiceConfig) ConfigMd5() *string {
	return R.s.ConfigMd5
}

func (R ResolverCloudServiceConfig) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverCloudServiceConfig) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverCloudServiceConfig) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverCloudServiceConfig) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverCloudServiceConfig) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeCloudServiceConfig(extra string) string {
	return "type CloudServiceConfig { " + extra + `
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
