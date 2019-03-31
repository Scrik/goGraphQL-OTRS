package api

type ResolverDynamicField struct {
	s TypeDynamicField
}

func (R *ResolverDynamicField) Set(s TypeDynamicField) {
	R.s = s
}

func (R ResolverDynamicField) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverDynamicField) InternalField() *int32 {
	i := int32(*R.s.InternalField)
	return &i
}

func (R ResolverDynamicField) Name() *string {
	return R.s.Name
}

func (R ResolverDynamicField) Label() *string {
	return R.s.Label
}

func (R ResolverDynamicField) FieldOrder() *int32 {
	return R.s.FieldOrder
}

func (R ResolverDynamicField) FieldType() *string {
	return R.s.FieldType
}

func (R ResolverDynamicField) ObjectType() *string {
	return R.s.ObjectType
}

func (R *ResolverDynamicField) Config() *string {
	str := string(*R.s.Config)
	return &str
}

func (R ResolverDynamicField) ValidId() *int32 {
	i := int32(*R.s.ValidId)
	return &i
}

func (R *ResolverDynamicField) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverDynamicField) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverDynamicField) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverDynamicField) ChangeBy() *int32 {
	return R.s.ChangeBy
}

func GenGqlTypeDynamicField(extra string) string {
	return "type DynamicField { " + extra + `
	Id: Int,
	InternalField: Int,
	Name: String,
	Label: String,
	FieldOrder: Int,
	FieldType: String,
	ObjectType: String,
	Config: String,
	ValidId: Int,
	CreateTime: String,
	CreateBy: Int,
	ChangeTime: String,
	ChangeBy: Int,
	}`
}
