package api

import "time"

type TypeCustomerPreferences struct {
	UserId           string  `db:"user_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *string `db:"preferences_value"`
}
