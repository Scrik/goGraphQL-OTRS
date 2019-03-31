package api

type ResolverSmimeSignerCertRelations struct {
	s TypeSmimeSignerCertRelations
}

func (R *ResolverSmimeSignerCertRelations) Set(s TypeSmimeSignerCertRelations) {
	R.s = s
}

func (R ResolverSmimeSignerCertRelations) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSmimeSignerCertRelations) CertHash() *string {
	return R.s.CertHash
}

func (R ResolverSmimeSignerCertRelations) CertFingerprint() *string {
	return R.s.CertFingerprint
}

func (R ResolverSmimeSignerCertRelations) CaHash() *string {
	return R.s.CaHash
}

func (R ResolverSmimeSignerCertRelations) CaFingerprint() *string {
	return R.s.CaFingerprint
}

func (R *ResolverSmimeSignerCertRelations) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSmimeSignerCertRelations) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSmimeSignerCertRelations) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSmimeSignerCertRelations) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSmimeSignerCertRelations(extra string) string {
	return "type SmimeSignerCertRelations { " + extra + `
	Id: Int,
	CertHash: String,
	CertFingerprint: String,
	CaHash: String,
	CaFingerprint: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
