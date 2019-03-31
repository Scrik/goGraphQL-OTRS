package api

import "time"

type TypeNotificationEventItem struct {
	NotificationId int32   `db:"notification_id"`
	EventKey       *string `db:"event_key"`
	EventValue     *string `db:"event_value"`
}
