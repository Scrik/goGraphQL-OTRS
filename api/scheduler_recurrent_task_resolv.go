package api

type ResolverSchedulerRecurrentTask struct {
	s TypeSchedulerRecurrentTask
}

func (R *ResolverSchedulerRecurrentTask) Set(s TypeSchedulerRecurrentTask) {
	R.s = s
}

func (R ResolverSchedulerRecurrentTask) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverSchedulerRecurrentTask) Name() *string {
	return R.s.Name
}

func (R ResolverSchedulerRecurrentTask) TaskType() *string {
	return R.s.TaskType
}

func (R *ResolverSchedulerRecurrentTask) LastExecutionTime() *string {
	str := R.s.LastExecutionTime.String()
	return &str
}

func (R *ResolverSchedulerRecurrentTask) LastWorkerTaskId() *string {
	str := fmt.Sprintf("%d", R.s.LastWorkerTaskId)
	return &str
}

func (R ResolverSchedulerRecurrentTask) LastWorkerStatus() *int32 {
	i := int32(*R.s.LastWorkerStatus)
	return &i
}

func (R ResolverSchedulerRecurrentTask) LastWorkerRunningTime() *int32 {
	return R.s.LastWorkerRunningTime
}

func (R *ResolverSchedulerRecurrentTask) LockKey() *string {
	str := fmt.Sprintf("%d", R.s.LockKey)
	return &str
}

func (R *ResolverSchedulerRecurrentTask) LockTime() *string {
	str := R.s.LockTime.String()
	return &str
}

func (R *ResolverSchedulerRecurrentTask) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R *ResolverSchedulerRecurrentTask) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func GenGqlTypeSchedulerRecurrentTask(extra string) string {
	return "type SchedulerRecurrentTask { " + extra + `
	Id: String,
	Name: String,
	TaskType: String,
	LastExecutionTime: String,
	LastWorkerTaskId: String,
	LastWorkerStatus: Int,
	LastWorkerRunningTime: Int,
	LockKey: String,
	LockTime: String,
	CreateTime: String,
	ChangeTime: String,
	}`
}
