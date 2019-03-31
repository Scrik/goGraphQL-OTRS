package api

type ResolverQueue struct {
	s TypeQueue
}

func (R *ResolverQueue) Set(s TypeQueue) {
	R.s = s
}

func (R ResolverQueue) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverQueue) Name() *string {
	return R.s.Name
}

func (R ResolverQueue) GroupId() *int32 {
	return R.s.GroupId
}

func (R ResolverQueue) UnlockTimeout() *int32 {
	return R.s.UnlockTimeout
}

func (R ResolverQueue) FirstResponseTime() *int32 {
	return R.s.FirstResponseTime
}

func (R ResolverQueue) FirstResponseNotify() *int32 {
	i := int32(*R.s.FirstResponseNotify)
	return &i
}

func (R ResolverQueue) UpdateTime() *int32 {
	return R.s.UpdateTime
}

func (R ResolverQueue) UpdateNotify() *int32 {
	i := int32(*R.s.UpdateNotify)
	return &i
}

func (R ResolverQueue) SolutionTime() *int32 {
	return R.s.SolutionTime
}

func (R ResolverQueue) SolutionNotify() *int32 {
	i := int32(*R.s.SolutionNotify)
	return &i
}

func (R ResolverQueue) SystemAddressId() *int32 {
	i := int32(*R.s.SystemAddressId)
	return &i
}

func (R ResolverQueue) CalendarName() *string {
	return R.s.CalendarName
}

func (R ResolverQueue) DefaultSignKey() *string {
	return R.s.DefaultSignKey
}

func (R ResolverQueue) SalutationId() *int32 {
	i := int32(*R.s.SalutationId)
	return &i
}

func (R ResolverQueue) SignatureId() *int32 {
	i := int32(*R.s.SignatureId)
	return &i
}

func (R ResolverQueue) FollowUpId() *int32 {
	i := int32(*R.s.FollowUpId)
	return &i
}

func (R ResolverQueue) FollowUpLock() *int32 {
	i := int32(*R.s.FollowUpLock)
	return &i
}

func (R ResolverQueue) Comments() *string {
	return R.s.Comments
}

func (R ResolverQueue) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverQueue) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverQueue) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverQueue) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverQueue) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeQueue(extra string) string {
	return "type Queue { " + extra + `
	Id: Int,
	Name: String,
	GroupId: Int,
	UnlockTimeout: Int,
	FirstResponseTime: Int,
	FirstResponseNotify: Int,
	UpdateTime: Int,
	UpdateNotify: Int,
	SolutionTime: Int,
	SolutionNotify: Int,
	SystemAddressId: Int,
	CalendarName: String,
	DefaultSignKey: String,
	SalutationId: Int,
	SignatureId: Int,
	FollowUpId: Int,
	FollowUpLock: Int,
	Comments: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
