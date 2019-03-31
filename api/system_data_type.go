package api

import "time"

type TypeSystemData struct {
	DataKey    string     `db:"data_key"`
	DataValue  *[]byte    `db:"data_value"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
