package dto

type PersonDto struct {
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
	Phone      string `db:"phone"`
	Age        int    `db:"age"`
	Password   string `db:"password"`
}
