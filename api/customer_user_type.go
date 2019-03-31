package api

import "time"

type TypeCustomerUser struct {
	Id         int32      `db:"id"`
	Login      *string    `db:"login"`
	Email      *string    `db:"email"`
	CustomerId *string    `db:"customer_id"`
	Pw         *string    `db:"pw"`
	Title      *string    `db:"title"`
	FirstName  *string    `db:"first_name"`
	LastName   *string    `db:"last_name"`
	Phone      *string    `db:"phone"`
	Fax        *string    `db:"fax"`
	Mobile     *string    `db:"mobile"`
	Street     *string    `db:"street"`
	Zip        *string    `db:"zip"`
	City       *string    `db:"city"`
	Country    *string    `db:"country"`
	Comments   *string    `db:"comments"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
