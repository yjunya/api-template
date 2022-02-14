package user

import (
	"github.com/yjunya/api-example/models"
	"github.com/yjunya/api-example/utils"
)

type User struct {
	*models.User
	Age int
}

const (
	NameMaxLength = 20
	AgeMin        = 0
	AgeMax        = 150
)

func BindUser(user *models.User) *User {
	return &User{
		User: user,
		Age:  utils.GetAge(user.Birthday),
	}
}
