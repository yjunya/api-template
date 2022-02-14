package user

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/yjunya/api-example/utils"
)

type SignupParams struct {
	FirebaseUID string     `json:"firebase_uid"`
	Name        string     `json:"name" binding:"required"`
	Birthday    *time.Time `json:"birthday" binding:"required"`
}

func (p *SignupParams) Validate() error {
	if err := p.validateName(); err != nil {
		return err
	}
	if err := p.validateBirthday(); err != nil {
		return err
	}
	return nil
}

func (p *SignupParams) validateName() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if utf8.RuneCountInString(p.Name) > NameMaxLength {
		return errors.New("name length is invalid")
	}
	return nil
}

func (p *SignupParams) validateBirthday() error {
	if p.Birthday == nil {
		return errors.New("birthday is required")
	}
	age := utils.GetAge(*p.Birthday)
	if age < AgeMin || age > AgeMax {
		return errors.New("age is invalid")
	}
	return nil
}
