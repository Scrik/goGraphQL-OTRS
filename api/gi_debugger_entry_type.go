package api

import "time"

type TypeGiDebuggerEntry struct {
	Id                int64      `db:"id"`
	CommunicationId   *string    `db:"communication_id"`
	CommunicationType *string    `db:"communication_type"`
	RemoteIp          *string    `db:"remote_ip"`
	WebserviceId      *int32     `db:"webservice_id"`
	CreateTime        *time.Time `db:"create_time"`
}
