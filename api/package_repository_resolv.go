package api

type ResolverPackageRepository struct {
	s TypePackageRepository
}

func (R *ResolverPackageRepository) Set(s TypePackageRepository) {
	R.s = s
}

func (R ResolverPackageRepository) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverPackageRepository) Name() *string {
	return R.s.Name
}

func (R ResolverPackageRepository) Version() *string {
	return R.s.Version
}

func (R ResolverPackageRepository) Vendor() *string {
	return R.s.Vendor
}

func (R ResolverPackageRepository) InstallStatus() *string {
	return R.s.InstallStatus
}

func (R ResolverPackageRepository) Filename() *string {
	return R.s.Filename
}

func (R ResolverPackageRepository) ContentType() *string {
	return R.s.ContentType
}

func (R *ResolverPackageRepository) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R *ResolverPackageRepository) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverPackageRepository) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverPackageRepository) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverPackageRepository) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypePackageRepository(extra string) string {
	return "type PackageRepository { " + extra + `
	Id: Int,
	Name: String,
	Version: String,
	Vendor: String,
	InstallStatus: String,
	Filename: String,
	ContentType: String,
	Content: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
