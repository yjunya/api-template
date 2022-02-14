package usecase

import (
	"context"
	"log"

	"github.com/yjunya/api-example/domains/user"
	"github.com/yjunya/api-example/domains/user/repository"
	"github.com/yjunya/api-example/models"
)

type Usecase interface {
	Get(ctx context.Context, id int) (*user.User, error)
	Signup(ctx context.Context, params *user.SignupParams) error
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Get(ctx context.Context, id int) (*user.User, error) {
	u, err := uc.repo.Get(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user.BindUser(u), nil
}

func (uc *usecase) Signup(ctx context.Context, params *user.SignupParams) error {
	u := &models.User{
		FirebaseUID: params.FirebaseUID,
		Name:        params.Name,
		Birthday:    *params.Birthday,
	}
	if err := uc.repo.Store(ctx, u); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
