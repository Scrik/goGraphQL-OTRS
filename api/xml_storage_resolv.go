package api

type ResolverXmlStorage struct {
	s TypeXmlStorage
}

func (R *ResolverXmlStorage) Set(s TypeXmlStorage) {
	R.s = s
}

func (R ResolverXmlStorage) XmlType() string {
	return R.s.XmlType
}

func (R ResolverXmlStorage) XmlKey() *string {
	return R.s.XmlKey
}

func (R ResolverXmlStorage) XmlContentKey() *string {
	return R.s.XmlContentKey
}

func (R ResolverXmlStorage) XmlContentValue() *string {
	return R.s.XmlContentValue
}

func GenGqlTypeXmlStorage(extra string) string {
	return "type XmlStorage { " + extra + `
	XmlType: String,
	XmlKey: String,
	XmlContentKey: String,
	XmlContentValue: String,
	}`
}
