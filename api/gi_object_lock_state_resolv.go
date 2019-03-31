package api

type ResolverGiObjectLockState struct {
	s TypeGiObjectLockState
}

func (R *ResolverGiObjectLockState) Set(s TypeGiObjectLockState) {
	R.s = s
}

func (R ResolverGiObjectLockState) WebserviceId() int32 {
	return R.s.WebserviceId
}

func (R ResolverGiObjectLockState) ObjectType() *string {
	return R.s.ObjectType
}

func (R *ResolverGiObjectLockState) ObjectId() *string {
	str := fmt.Sprintf("%d", R.s.ObjectId)
	return &str
}

func (R ResolverGiObjectLockState) LockState() *string {
	return R.s.LockState
}

func (R ResolverGiObjectLockState) LockStateCounter() *int32 {
	return R.s.LockStateCounter
}

func (R *ResolverGiObjectLockState) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R *ResolverGiObjectLockState) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func GenGqlTypeGiObjectLockState(extra string) string {
	return "type GiObjectLockState { " + extra + `
	WebserviceId: Int,
	ObjectType: String,
	ObjectId: String,
	LockState: String,
	LockStateCounter: Int,
	CreateTime: String,
	ChangeTime: String,
	}`
}
