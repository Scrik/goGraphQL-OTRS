package api

import "time"

type TypeSearchProfile struct {
	Login        string  `db:"login"`
	ProfileName  *string `db:"profile_name"`
	ProfileType  *string `db:"profile_type"`
	ProfileKey   *string `db:"profile_key"`
	ProfileValue *string `db:"profile_value"`
}
