package api

import "time"

type TypeServicePreferences struct {
	ServiceId        int32   `db:"service_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *string `db:"preferences_value"`
}
