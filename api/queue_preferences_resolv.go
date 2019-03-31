package api

type ResolverQueuePreferences struct {
	s TypeQueuePreferences
}

func (R *ResolverQueuePreferences) Set(s TypeQueuePreferences) {
	R.s = s
}

func (R ResolverQueuePreferences) QueueId() int32 {
	return R.s.QueueId
}

func (R ResolverQueuePreferences) PreferencesKey() *string {
	return R.s.PreferencesKey
}

func (R ResolverQueuePreferences) PreferencesValue() *string {
	return R.s.PreferencesValue
}

func GenGqlTypeQueuePreferences(extra string) string {
	return "type QueuePreferences { " + extra + `
	QueueId: Int,
	PreferencesKey: String,
	PreferencesValue: String,
	}`
}
