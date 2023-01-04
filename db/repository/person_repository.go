package repository

import (
	"database/sql"
	"errors"
	"github.com/Unbel1evab7e/bank-interface/db/entity"
	"github.com/Unbel1evab7e/bank-interface/domain/dto"
)

type PersonRepository struct {
	Db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) GetByPhoneAndPassword(phone string) (*entity.Person, error) {
	query, err := r.Db.Query("select * from public.person p where p.phone = $1", phone)

	if err != nil {
		return nil, err
	}

	var person entity.Person

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Phone,
			&person.Age,
			&person.CreditScore,
			&person.Password,
		)
	}

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) CreatePerson(dto *dto.PersonDto) error {
	res, err := r.Db.Exec(`insert into public.person (name, surname, patronymic, phone, age,password)
					values ($1, $2, $3, $4, $5, $6)`,
		dto.Name, dto.Surname, dto.Patronymic, dto.Phone, dto.Age, dto.Password)

	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if ra < 1 {
		return errors.New("person was not saved")
	}

	return nil
}
