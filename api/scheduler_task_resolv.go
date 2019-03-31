package api

type ResolverSchedulerTask struct {
	s TypeSchedulerTask
}

func (R *ResolverSchedulerTask) Set(s TypeSchedulerTask) {
	R.s = s
}

func (R ResolverSchedulerTask) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverSchedulerTask) Ident() *string {
	str := fmt.Sprintf("%d", R.s.Ident)
	return &str
}

func (R ResolverSchedulerTask) Name() *string {
	return R.s.Name
}

func (R ResolverSchedulerTask) TaskType() *string {
	return R.s.TaskType
}

func (R *ResolverSchedulerTask) TaskData() *string {
	str := string(*R.s.TaskData)
	return &str
}

func (R ResolverSchedulerTask) Attempts() *int32 {
	i := int32(*R.s.Attempts)
	return &i
}

func (R *ResolverSchedulerTask) LockKey() *string {
	str := fmt.Sprintf("%d", R.s.LockKey)
	return &str
}

func (R *ResolverSchedulerTask) LockTime() *string {
	str := R.s.LockTime.String()
	return &str
}

func (R *ResolverSchedulerTask) LockUpdateTime() *string {
	str := R.s.LockUpdateTime.String()
	return &str
}

func (R *ResolverSchedulerTask) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeSchedulerTask(extra string) string {
	return "type SchedulerTask { " + extra + `
	Id: String,
	Ident: String,
	Name: String,
	TaskType: String,
	TaskData: String,
	Attempts: Int,
	LockKey: String,
	LockTime: String,
	LockUpdateTime: String,
	CreateTime: String,
	}`
}
