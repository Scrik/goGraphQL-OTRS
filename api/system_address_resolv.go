package api

type ResolverSystemAddress struct {
	s TypeSystemAddress
}

func (R *ResolverSystemAddress) Set(s TypeSystemAddress) {
	R.s = s
}

func (R ResolverSystemAddress) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSystemAddress) Value0() *string {
	return R.s.Value0
}

func (R ResolverSystemAddress) Value1() *string {
	return R.s.Value1
}

func (R ResolverSystemAddress) Value2() *string {
	return R.s.Value2
}

func (R ResolverSystemAddress) Value3() *string {
	return R.s.Value3
}

func (R ResolverSystemAddress) QueueId() *int32 {
	return R.s.QueueId
}

func (R ResolverSystemAddress) Comments() *string {
	return R.s.Comments
}

func (R ResolverSystemAddress) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverSystemAddress) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSystemAddress) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSystemAddress) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSystemAddress) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSystemAddress(extra string) string {
	return "type SystemAddress { " + extra + `
	Id: Int,
	Value0: String,
	Value1: String,
	Value2: String,
	Value3: String,
	QueueId: Int,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
