package api

type ResolverCustomerCompany struct {
	s TypeCustomerCompany
}

func (R *ResolverCustomerCompany) Set(s TypeCustomerCompany) {
	R.s = s
}

func (R ResolverCustomerCompany) CustomerId() string {
	return R.s.CustomerId
}

func (R ResolverCustomerCompany) Name() *string {
	return R.s.Name
}

func (R ResolverCustomerCompany) Street() *string {
	return R.s.Street
}

func (R ResolverCustomerCompany) Zip() *string {
	return R.s.Zip
}

func (R ResolverCustomerCompany) City() *string {
	return R.s.City
}

func (R ResolverCustomerCompany) Country() *string {
	return R.s.Country
}

func (R ResolverCustomerCompany) Url() *string {
	return R.s.Url
}

func (R ResolverCustomerCompany) Comments() *string {
	return R.s.Comments
}

func (R ResolverCustomerCompany) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverCustomerCompany) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverCustomerCompany) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverCustomerCompany) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverCustomerCompany) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeCustomerCompany(extra string) string {
	return "type CustomerCompany { " + extra + `
	CustomerId: String,
	Name: String,
	Street: String,
	Zip: String,
	City: String,
	Country: String,
	Url: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
