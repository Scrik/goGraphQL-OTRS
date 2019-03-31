package api

import "time"

type TypeSessions struct {
	Id         int64   `db:"id"`
	SessionId  *string `db:"session_id"`
	DataKey    *string `db:"data_key"`
	DataValue  *string `db:"data_value"`
	Serialized *int8   `db:"serialized"`
}
