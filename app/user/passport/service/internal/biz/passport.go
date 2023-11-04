package biz

import (
	"context"
	do "douyin/app/user/passport/common/entity"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type PassportRepo interface {
	CreateUser(ctx context.Context, user *do.User) error
	GetUserByName(ctx context.Context, name string) (*do.User, error)
	GetUserById(ctx context.Context, id int64) (*do.User, error)
	MGetUserById(ctx context.Context, ids []int64) ([]*do.User, error)
}

type PassportUsecase struct {
	repo PassportRepo
	log  *log.Logger
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUsecase {
	return &PassportUsecase{
		repo: repo,
		log:  &logger,
	}
}

func (u *PassportUsecase) CreateUser(ctx context.Context, user *do.User) error {
	pwd := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.EncryptedPassword = string(hashedPassword)
	return u.repo.CreateUser(ctx, user)
}

func (u *PassportUsecase) VerifyPassword(ctx context.Context, user *do.User) (verified bool, uid int64, err error) {
	pwd := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return false, 0, err
	}
	us, err := u.repo.GetUserByName(ctx, user.Name)
	if err != nil {
		return false, 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(us.EncryptedPassword), hashedPassword)
	if err != nil {
		return false, 0, err
	}
	return true, us.ID, nil
}

func (u *PassportUsecase) GetUserByID(ctx context.Context, id int64) (*do.User, error) {
	return u.repo.GetUserById(ctx, id)
}

func (u *PassportUsecase) MGetUserByID(ctx context.Context, ids []int64) ([]*do.User, error) {
	return u.repo.MGetUserById(ctx, ids)
}
