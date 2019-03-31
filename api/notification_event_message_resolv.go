package api

type ResolverNotificationEventMessage struct {
	s TypeNotificationEventMessage
}

func (R *ResolverNotificationEventMessage) Set(s TypeNotificationEventMessage) {
	R.s = s
}

func (R ResolverNotificationEventMessage) Id() string {
	return fmt.Sprintf("%d", R.s.Id)
}

func (R ResolverNotificationEventMessage) NotificationId() *int32 {
	return R.s.NotificationId
}

func (R ResolverNotificationEventMessage) Subject() *string {
	return R.s.Subject
}

func (R ResolverNotificationEventMessage) Text() *string {
	return R.s.Text
}

func (R ResolverNotificationEventMessage) ContentType() *string {
	return R.s.ContentType
}

func (R ResolverNotificationEventMessage) Language() *string {
	return R.s.Language
}

func GenGqlTypeNotificationEventMessage(extra string) string {
	return "type NotificationEventMessage { " + extra + `
	Id: Int,
	NotificationId: Int,
	Subject: String,
	Text: String,
	ContentType: String,
	Language: String,
	}`
}
