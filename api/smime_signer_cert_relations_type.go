package api

import "time"

type TypeSmimeSignerCertRelations struct {
	Id              int32      `db:"id"`
	CertHash        *string    `db:"cert_hash"`
	CertFingerprint *string    `db:"cert_fingerprint"`
	CaHash          *string    `db:"ca_hash"`
	CaFingerprint   *string    `db:"ca_fingerprint"`
	CreateTime      *time.Time `db:"create_time"`
	CreateBy        *int32     `db:"create_by"`
	ChangeTime      *time.Time `db:"change_time"`
	ChangeBy        *int32     `db:"change_by"`
}
