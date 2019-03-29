package api

import "fmt"

type ResolverArticleAttachment struct {
	s TypeArticleAttachment
}

func (R *ResolverArticleAttachment) Set(s TypeArticleAttachment) {
	R.s = s
}

func (R ResolverArticleAttachment) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R *ResolverArticleAttachment) ArticleId() *string {
	str := fmt.Sprintf("%d", R.s.ArticleId)
	return &str
}

func (R ResolverArticleAttachment) Filename() *string {
	return R.s.Filename
}

func (R ResolverArticleAttachment) ContentSize() *string {
	return R.s.ContentSize
}

func (R ResolverArticleAttachment) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverArticleAttachment) ContentId() *string {
	return R.s.ContentId
}

func (R ResolverArticleAttachment) ContentAlternative() *string {
	return R.s.ContentAlternative
}

func (R ResolverArticleAttachment) Disposition() *string {
	return R.s.Disposition
}

func (R *ResolverArticleAttachment) Content() *string {
	str := string(*R.s.Content)
	return &str
}

func (R *ResolverArticleAttachment) CreateTime() *string {
	str := R.s.CreateTime.String()
	return &str
}

func (R ResolverArticleAttachment) CreateBy() *int32 {
	return R.s.CreateBy
}

func (R *ResolverArticleAttachment) ChangeTime() *string {
	str := R.s.ChangeTime.String()
	return &str
}

func (R ResolverArticleAttachment) ChangeBy() *int32 {
	return R.s.ChangeBy
}
