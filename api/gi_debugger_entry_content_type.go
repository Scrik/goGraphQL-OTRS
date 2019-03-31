package api

import "time"

type TypeGiDebuggerEntryContent struct {
	Id                int64      `db:"id"`
	GiDebuggerEntryId *int64     `db:"gi_debugger_entry_id"`
	DebugLevel        *string    `db:"debug_level"`
	Subject           *string    `db:"subject"`
	Content           *[]byte    `db:"content"`
	CreateTime        *time.Time `db:"create_time"`
}
