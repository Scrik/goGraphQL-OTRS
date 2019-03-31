package api

import "time"

type TypeSystemMaintenance struct {
	Id               int32      `db:"id"`
	StartDate        *int32     `db:"start_date"`
	StopDate         *int32     `db:"stop_date"`
	Comments         *string    `db:"comments"`
	LoginMessage     *string    `db:"login_message"`
	ShowLoginMessage *int8      `db:"show_login_message"`
	NotifyMessage    *string    `db:"notify_message"`
	ValidId          *int8      `db:"valid_id"`
	CreateTime       *time.Time `db:"create_time"`
	CreateBy         *int32     `db:"create_by"`
	ChangeTime       *time.Time `db:"change_time"`
	ChangeBy         *int32     `db:"change_by"`
}
