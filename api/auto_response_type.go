package api

import "time"

type TypeAutoResponse struct {
	Id              int32      `db:"id"`
	Name            *string    `db:"name"`
	Text0           *string    `db:"text0"`
	Text1           *string    `db:"text1"`
	TypeId          *int8      `db:"type_id"`
	SystemAddressId *int8      `db:"system_address_id"`
	ContentType     *string    `db:"content_type"`
	Comments        *string    `db:"comments"`
	ValidId         *int8      `db:"valid_id"`
	CreateTime      *time.Time `db:"create_time"`
	CreateBy        *int32     `db:"create_by"`
	ChangeTime      *time.Time `db:"change_time"`
	ChangeBy        *int32     `db:"change_by"`
}
