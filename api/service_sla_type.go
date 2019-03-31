package api

import "time"

type TypeServiceSla struct {
	ServiceId int32  `db:"service_id"`
	SlaId     *int32 `db:"sla_id"`
}
