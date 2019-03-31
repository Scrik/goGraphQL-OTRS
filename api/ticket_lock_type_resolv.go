package api

type ResolverTicketLockType struct {
	s TypeTicketLockType
}

func (R *ResolverTicketLockType) Set(s TypeTicketLockType) {
	R.s = s
}

func (R ResolverTicketLockType) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverTicketLockType) Name() *string {
	return R.s.Name
}

func (R ResolverTicketLockType) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverTicketLockType) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverTicketLockType) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverTicketLockType) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverTicketLockType) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeTicketLockType(extra string) string {
	return "type TicketLockType { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
