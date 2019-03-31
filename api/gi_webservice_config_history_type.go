package api

import "time"

type TypeGiWebserviceConfigHistory struct {
	Id         int64      `db:"id"`
	ConfigId   *int32     `db:"config_id"`
	Config     *[]byte    `db:"config"`
	ConfigMd5  *string    `db:"config_md5"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
