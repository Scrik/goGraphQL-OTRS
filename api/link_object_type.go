package api

import "time"

type TypeLinkObject struct {
	Id   int8    `db:"id"`
	Name *string `db:"name"`
}
