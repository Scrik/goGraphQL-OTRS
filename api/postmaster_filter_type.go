package api

import "time"

type TypePostmasterFilter struct {
	FName  string  `db:"f_name"`
	FStop  *int8   `db:"f_stop"`
	FType  *string `db:"f_type"`
	FKey   *string `db:"f_key"`
	FValue *string `db:"f_value"`
	FNot   *int8   `db:"f_not"`
}
