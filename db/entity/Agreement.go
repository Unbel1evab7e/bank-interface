package entity

import (
	"database/sql"
)

type Agreement struct {
	ID            int             `db:"id"`
	PersonID      int             `db:"person_id"`
	TariffID      int             `db:"tariff_id"`
	RequestMoney  float64         `db:"request_money"`
	RequestDays   int             `db:"request_days"`
	ApprovedMoney sql.NullFloat64 `db:"approved_money"`
	ApprovedDays  sql.NullInt64   `db:"approved_days"`
}
