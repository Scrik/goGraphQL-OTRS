package api

type ResolverWebUploadCache struct {
	s TypeWebUploadCache
}

func (R *ResolverWebUploadCache) Set(s TypeWebUploadCache) {
	R.s = s
}

func (R ResolverWebUploadCache) FormId() string {
	return R.s.FormId
}

func (R ResolverWebUploadCache) Filename() *string {
	return R.s.Filename
}

func (R ResolverWebUploadCache) ContentId() *string {
	return R.s.ContentId
}

func (R ResolverWebUploadCache) ContentSize() *string {
	return R.s.ContentSize
}

func (R ResolverWebUploadCache) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverWebUploadCache) Disposition() *string {
	return R.s.Disposition
}

func (R *ResolverWebUploadCache) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R *ResolverWebUploadCache) CreateTimeUnix() *string {
	str := fmt.Sprintf("%d", R.s.CreateTimeUnix)
	return &str
}

func GenGqlTypeWebUploadCache(extra string) string {
	return "type WebUploadCache { " + extra + `
	FormId: String,
	Filename: String,
	ContentId: String,
	ContentSize: String,
	ContentType: String,
	Disposition: String,
	Content: String,
	CreateTimeUnix: String,
	}`
}
