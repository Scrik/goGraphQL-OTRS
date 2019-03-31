package api

type ResolverDynamicFieldValue struct {
	s TypeDynamicFieldValue
}

func (R *ResolverDynamicFieldValue) Set(s TypeDynamicFieldValue) {
	R.s = s
}

func (R ResolverDynamicFieldValue) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverDynamicFieldValue) FieldId() *int32 {
	return R.s.FieldId
}

func (R *ResolverDynamicFieldValue) ObjectId() *string {
	str := fmt.Sprintf("%d", R.s.ObjectId)
	return &str
}

func (R ResolverDynamicFieldValue) ValueText() *string {
	return R.s.ValueText
}

func (R *ResolverDynamicFieldValue) ValueDate() *string {
	str := R.s.ValueDate.String()
	return &str
}

func (R *ResolverDynamicFieldValue) ValueInt() *string {
	str := fmt.Sprintf("%d", R.s.ValueInt)
	return &str
}

func GenGqlTypeDynamicFieldValue(extra string) string {
	return "type DynamicFieldValue { " + extra + `
	Id: Int,
	FieldId: Int,
	ObjectId: String,
	ValueText: String,
	ValueDate: String,
	ValueInt: String,
	}`
}
