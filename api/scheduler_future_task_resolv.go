package api

type ResolverSchedulerFutureTask struct {
	s TypeSchedulerFutureTask
}

func (R *ResolverSchedulerFutureTask) Set(s TypeSchedulerFutureTask) {
	R.s = s
}

func (R ResolverSchedulerFutureTask) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverSchedulerFutureTask) Ident() *string {
	str := fmt.Sprintf("%d", R.s.Ident)
	return &str
}

func (R *ResolverSchedulerFutureTask) ExecutionTime() *string {
	str := R.s.ExecutionTime.String()
	return &str
}

func (R ResolverSchedulerFutureTask) Name() *string {
	return R.s.Name
}

func (R ResolverSchedulerFutureTask) TaskType() *string {
	return R.s.TaskType
}

func (R *ResolverSchedulerFutureTask) TaskData() *string {
	str := string(*R.s.TaskData)
	return &str
}

func (R ResolverSchedulerFutureTask) Attempts() *int32 {
	i := int32(*R.s.Attempts)
	return &i
}

func (R *ResolverSchedulerFutureTask) LockKey() *string {
	str := fmt.Sprintf("%d", R.s.LockKey)
	return &str
}

func (R *ResolverSchedulerFutureTask) LockTime() *string {
	str := R.s.LockTime.String()
	return &str
}

func (R *ResolverSchedulerFutureTask) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func GenGqlTypeSchedulerFutureTask(extra string) string {
	return "type SchedulerFutureTask { " + extra + `
	Id: String,
	Ident: String,
	ExecutionTime: String,
	Name: String,
	TaskType: String,
	TaskData: String,
	Attempts: Int,
	LockKey: String,
	LockTime: String,
	CreateTime: String,
	}`
}
