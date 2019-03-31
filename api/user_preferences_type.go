package api

import "time"

type TypeUserPreferences struct {
	UserId           int32   `db:"user_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *[]byte `db:"preferences_value"`
}
