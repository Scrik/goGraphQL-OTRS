package api

type ResolverNotificationEventItem struct {
	s TypeNotificationEventItem
}

func (R *ResolverNotificationEventItem) Set(s TypeNotificationEventItem) {
	R.s = s
}

func (R ResolverNotificationEventItem) NotificationId() int32 {
	return R.s.NotificationId
}

func (R ResolverNotificationEventItem) EventKey() *string {
	return R.s.EventKey
}

func (R ResolverNotificationEventItem) EventValue() *string {
	return R.s.EventValue
}

func GenGqlTypeNotificationEventItem(extra string) string {
	return "type NotificationEventItem { " + extra + `
	NotificationId: Int,
	EventKey: String,
	EventValue: String,
	}`
}
