package service

import (
	"errors"
	"github.com/Unbel1evab7e/bank-interface/db/entity"
	"github.com/Unbel1evab7e/bank-interface/db/repository"
	"github.com/Unbel1evab7e/bank-interface/domain"
	"github.com/Unbel1evab7e/bank-interface/domain/dto"
	"golang.org/x/crypto/bcrypt"
)

type PersonService struct {
	PersonRepository *repository.PersonRepository
}

func NewPersonService(personRepository *repository.PersonRepository) *PersonService {
	return &PersonService{PersonRepository: personRepository}
}

func (s *PersonService) GetPersonByPhoneAndPassword(phone string, password string) (*entity.Person, error) {
	person, err := s.PersonRepository.GetByPhoneAndPassword(phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (s *PersonService) SavePerson(personDto *dto.PersonDto) error {
	err := validate(personDto)

	if err != nil {
		return err
	}

	newPass, err := bcrypt.GenerateFromPassword([]byte(personDto.Password), 14)

	if err != nil {
		return err
	}

	personDto.Password = string(newPass)

	err = s.PersonRepository.CreatePerson(personDto)

	if err != nil {
		return err
	}

	return nil
}
func validate(personDto *dto.PersonDto) error {
	if !domain.PhonePattern.Match([]byte(personDto.Phone)) {
		errors.New("invalid phone")
	}

	return nil
}
