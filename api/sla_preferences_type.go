package api

import "time"

type TypeSlaPreferences struct {
	SlaId            int32   `db:"sla_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *string `db:"preferences_value"`
}
