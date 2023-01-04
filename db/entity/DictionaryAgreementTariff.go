package entity

import (
	"database/sql"
)

type DictionaryAgreementTariff struct {
	ID   int            `db:"id"`
	Name sql.NullString `db:"name"`
}
