package entity

import (
	"database/sql"
)

type Person struct {
	ID          int           `db:"id"`
	Name        string        `db:"name"`
	Surname     string        `db:"surname"`
	Patronymic  string        `db:"patronymic"`
	Phone       string        `db:"phone"`
	Age         int           `db:"age"`
	CreditScore sql.NullInt64 `db:"credit_score"`
	Password    string        `db:"password"`
}
