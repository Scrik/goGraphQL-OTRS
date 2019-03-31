package api

type ResolverCustomerUser struct {
	s TypeCustomerUser
}

func (R *ResolverCustomerUser) Set(s TypeCustomerUser) {
	R.s = s
}

func (R ResolverCustomerUser) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverCustomerUser) Login() *string {
	return R.s.Login
}

func (R ResolverCustomerUser) Email() *string {
	return R.s.Email
}

func (R ResolverCustomerUser) CustomerId() *string {
	return R.s.CustomerId
}

func (R ResolverCustomerUser) Pw() *string {
	return R.s.Pw
}

func (R ResolverCustomerUser) Title() *string {
	return R.s.Title
}

func (R ResolverCustomerUser) FirstName() *string {
	return R.s.FirstName
}

func (R ResolverCustomerUser) LastName() *string {
	return R.s.LastName
}

func (R ResolverCustomerUser) Phone() *string {
	return R.s.Phone
}

func (R ResolverCustomerUser) Fax() *string {
	return R.s.Fax
}

func (R ResolverCustomerUser) Mobile() *string {
	return R.s.Mobile
}

func (R ResolverCustomerUser) Street() *string {
	return R.s.Street
}

func (R ResolverCustomerUser) Zip() *string {
	return R.s.Zip
}

func (R ResolverCustomerUser) City() *string {
	return R.s.City
}

func (R ResolverCustomerUser) Country() *string {
	return R.s.Country
}

func (R ResolverCustomerUser) Comments() *string {
	return R.s.Comments
}

func (R ResolverCustomerUser) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverCustomerUser) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverCustomerUser) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverCustomerUser) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverCustomerUser) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeCustomerUser(extra string) string {
	return "type CustomerUser { " + extra + `
	Id: Int,
	Login: String,
	Email: String,
	CustomerId: String,
	Pw: String,
	Title: String,
	FirstName: String,
	LastName: String,
	Phone: String,
	Fax: String,
	Mobile: String,
	Street: String,
	Zip: String,
	City: String,
	Country: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
