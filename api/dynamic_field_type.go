package api

import "time"

type TypeDynamicField struct {
	Id            int32      `db:"id"`
	InternalField *int8      `db:"internal_field"`
	Name          *string    `db:"name"`
	Label         *string    `db:"label"`
	FieldOrder    *int32     `db:"field_order"`
	FieldType     *string    `db:"field_type"`
	ObjectType    *string    `db:"object_type"`
	Config        *[]byte    `db:"config"`
	ValidId       *int8      `db:"valid_id"`
	CreateTime    *time.Time `db:"create_time"`
	CreateBy      *int32     `db:"create_by"`
	ChangeTime    *time.Time `db:"change_time"`
	ChangeBy      *int32     `db:"change_by"`
}
