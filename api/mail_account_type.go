package api

import "time"

type TypeMailAccount struct {
	Id          int32      `db:"id"`
	Login       *string    `db:"login"`
	Pw          *string    `db:"pw"`
	Host        *string    `db:"host"`
	AccountType *string    `db:"account_type"`
	QueueId     *int32     `db:"queue_id"`
	Trusted     *int8      `db:"trusted"`
	ImapFolder  *string    `db:"imap_folder"`
	Comments    *string    `db:"comments"`
	ValidId     *int8      `db:"valid_id"`
	CreateTime  *time.Time `db:"create_time"`
	CreateBy    *int32     `db:"create_by"`
	ChangeTime  *time.Time `db:"change_time"`
	ChangeBy    *int32     `db:"change_by"`
}
