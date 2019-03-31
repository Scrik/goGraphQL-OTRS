package api

type ResolverSystemMaintenance struct {
	s TypeSystemMaintenance
}

func (R *ResolverSystemMaintenance) Set(s TypeSystemMaintenance) {
	R.s = s
}

func (R ResolverSystemMaintenance) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSystemMaintenance) StartDate() *int32 {
	return R.s.StartDate
}

func (R ResolverSystemMaintenance) StopDate() *int32 {
	return R.s.StopDate
}

func (R ResolverSystemMaintenance) Comments() *string {
	return R.s.Comments
}

func (R ResolverSystemMaintenance) LoginMessage() *string {
	return R.s.LoginMessage
}

func (R ResolverSystemMaintenance) ShowLoginMessage() *int32 {
	i := int32(*R.s.ShowLoginMessage)
	return &i
}

func (R ResolverSystemMaintenance) NotifyMessage() *string {
	return R.s.NotifyMessage
}

func (R ResolverSystemMaintenance) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverSystemMaintenance) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverSystemMaintenance) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverSystemMaintenance) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverSystemMaintenance) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeSystemMaintenance(extra string) string {
	return "type SystemMaintenance { " + extra + `
	Id: Int,
	StartDate: Int,
	StopDate: Int,
	Comments: String,
	LoginMessage: String,
	ShowLoginMessage: Int,
	NotifyMessage: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
