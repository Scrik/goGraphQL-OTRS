package api

import "time"

type TypeGenericAgentJobs struct {
	JobName  string  `db:"job_name"`
	JobKey   *string `db:"job_key"`
	JobValue *string `db:"job_value"`
}
