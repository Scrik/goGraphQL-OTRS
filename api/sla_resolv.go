package api

type ResolverSla struct {
	s TypeSla
}

func (R *ResolverSla) Set(s TypeSla) {
	R.s = s
}

func (R ResolverSla) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSla) Name() *string {
	return R.s.Name
}

func (R ResolverSla) CalendarName() *string {
	return R.s.CalendarName
}

func (R ResolverSla) FirstResponseTime() *int32 {
	return R.s.FirstResponseTime
}

func (R ResolverSla) FirstResponseNotify() *int32 {
	i := int32(*R.s.FirstResponseNotify)
	return &i
}

func (R ResolverSla) UpdateTime() *int32 {
	return R.s.UpdateTime
}

func (R ResolverSla) UpdateNotify() *int32 {
	i := int32(*R.s.UpdateNotify)
	return &i
}

func (R ResolverSla) SolutionTime() *int32 {
	return R.s.SolutionTime
}

func (R ResolverSla) SolutionNotify() *int32 {
	i := int32(*R.s.SolutionNotify)
	return &i
}

func (R ResolverSla) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R ResolverSla) Comments() *string {
	return R.s.Comments
}

func (R *ResolverSla) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSla) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSla) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSla) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSla(extra string) string {
	return "type Sla { " + extra + `
	Id: Int,
	Name: String,
	CalendarName: String,
	FirstResponseTime: Int,
	FirstResponseNotify: Int,
	UpdateTime: Int,
	UpdateNotify: Int,
	SolutionTime: Int,
	SolutionNotify: Int,
	ValidId: Int,
	Comments: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
