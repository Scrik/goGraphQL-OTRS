package api

import "time"

type TypeServiceCustomerUser struct {
	CustomerUserLogin string     `db:"customer_user_login"`
	ServiceId         *int32     `db:"service_id"`
	CreateTime        *time.Time `db:"create_time"`
	CreateBy          *int32     `db:"create_by"`
}
