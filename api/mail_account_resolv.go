package api

type ResolverMailAccount struct {
	s TypeMailAccount
}

func (R *ResolverMailAccount) Set(s TypeMailAccount) {
	R.s = s
}

func (R ResolverMailAccount) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverMailAccount) Login() *string {
	return R.s.Login
}

func (R ResolverMailAccount) Pw() *string {
	return R.s.Pw
}

func (R ResolverMailAccount) Host() *string {
	return R.s.Host
}

func (R ResolverMailAccount) AccountType() *string {
	return R.s.AccountType
}

func (R ResolverMailAccount) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverMailAccount) Trusted() *int32 {
	i := int32(*R.s.Trusted)
	return &i
}

func (R ResolverMailAccount) ImapFolder() *string {
	return R.s.ImapFolder
}

func (R ResolverMailAccount) Comments() *string {
	return R.s.Comments
}

func (R ResolverMailAccount) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverMailAccount) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverMailAccount) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverMailAccount) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverMailAccount) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeMailAccount(extra string) string {
	return "type MailAccount { " + extra + `
	Id: Int,
	Login: String,
	Pw: String,
	Host: String,
	AccountType: String,
	QueueId: Int,
	Trusted: Int,
	ImapFolder: String,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
