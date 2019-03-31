package api

import "time"

type TypeQueuePreferences struct {
	QueueId          int32   `db:"queue_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *string `db:"preferences_value"`
}
