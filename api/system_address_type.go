package api

import "time"

type TypeSystemAddress struct {
	Id         int8       `db:"id"`
	Value0     *string    `db:"value0"`
	Value1     *string    `db:"value1"`
	Value2     *string    `db:"value2"`
	Value3     *string    `db:"value3"`
	QueueId    *int32     `db:"queue_id"`
	Comments   *string    `db:"comments"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
