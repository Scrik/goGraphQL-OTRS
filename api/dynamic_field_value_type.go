package api

import "time"

type TypeDynamicFieldValue struct {
	Id        int32      `db:"id"`
	FieldId   *int32     `db:"field_id"`
	ObjectId  *int64     `db:"object_id"`
	ValueText *string    `db:"value_text"`
	ValueDate *time.Time `db:"value_date"`
	ValueInt  *int64     `db:"value_int"`
}
