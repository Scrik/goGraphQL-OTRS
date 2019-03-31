package api

import "time"

type TypeVirtualFsPreferences struct {
	VirtualFsId      int64   `db:"virtual_fs_id"`
	PreferencesKey   *string `db:"preferences_key"`
	PreferencesValue *string `db:"preferences_value"`
}
