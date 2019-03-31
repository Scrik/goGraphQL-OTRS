package api

import "time"

type TypePackageRepository struct {
	Id            int32      `db:"id"`
	Name          *string    `db:"name"`
	Version       *string    `db:"version"`
	Vendor        *string    `db:"vendor"`
	InstallStatus *string    `db:"install_status"`
	Filename      *string    `db:"filename"`
	ContentType   *string    `db:"content_type"`
	Content       *[]byte    `db:"content"`
	CreateTime    *time.Time `db:"create_time"`
	CreateBy      *int32     `db:"create_by"`
	ChangeTime    *time.Time `db:"change_time"`
	ChangeBy      *int32     `db:"change_by"`
}
