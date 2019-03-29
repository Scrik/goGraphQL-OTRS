package api

import "time"

type TypeCustomerCompany struct {
	CustomerId string     `db:"customer_id"`
	Name       *string    `db:"name"`
	Street     *string    `db:"street"`
	Zip        *string    `db:"zip"`
	City       *string    `db:"city"`
	Country    *string    `db:"country"`
	Url        *string    `db:"url"`
	Comments   *string    `db:"comments"`
	ValidId    *int8      `db:"valid_id"`
	CreateTime *time.Time `db:"create_time"`
	CreateBy   *int32     `db:"create_by"`
	ChangeTime *time.Time `db:"change_time"`
	ChangeBy   *int32     `db:"change_by"`
}
