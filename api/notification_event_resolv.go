package api

type ResolverNotificationEvent struct {
	s TypeNotificationEvent
}

func (R *ResolverNotificationEvent) Set(s TypeNotificationEvent) {
	R.s = s
}

func (R ResolverNotificationEvent) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverNotificationEvent) Name() *string {
	return R.s.Name
}

func (R ResolverNotificationEvent) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R ResolverNotificationEvent) Comments() *string {
	return R.s.Comments
}

func (R *ResolverNotificationEvent) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverNotificationEvent) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverNotificationEvent) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverNotificationEvent) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeNotificationEvent(extra string) string {
	return "type NotificationEvent { " + extra + `
	Id: Int,
	Name: String,
	ValidId: Int,
	Comments: String,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
