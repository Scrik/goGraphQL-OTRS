package api

import "time"

type TypeNotificationEventMessage struct {
	Id             int32   `db:"id"`
	NotificationId *int32  `db:"notification_id"`
	Subject        *string `db:"subject"`
	Text           *string `db:"text"`
	ContentType    *string `db:"content_type"`
	Language       *string `db:"language"`
}
