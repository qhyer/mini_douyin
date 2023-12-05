package biz

import (
	"context"
	"douyin/app/user/passport/common/constants"
	do "douyin/app/user/passport/common/entity"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/singleflight"
)

type PassportRepo interface {
	CreateUser(ctx context.Context, user *do.User) error
	GetUserByName(ctx context.Context, name string) (*do.User, error)
	GetUserById(ctx context.Context, id int64) (*do.User, error)
	MGetUserById(ctx context.Context, ids []int64) ([]*do.User, error)
}

type PassportUsecase struct {
	repo PassportRepo
	log  *log.Helper
	sf   *singleflight.Group
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUsecase {
	return &PassportUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		sf:   &singleflight.Group{},
	}
}

// CreateUser 创建用户
func (u *PassportUsecase) CreateUser(ctx context.Context, user *do.User) (uid int64, err error) {
	uid = int64(0)
	user.ID = uid
	pwd := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("create user error(%v)", err)
		return 0, err
	}
	user.EncryptedPassword = string(hashedPassword)
	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		u.log.Errorf("create user error(%v)", err)
		return 0, err
	}
	return user.ID, nil
}

// VerifyPassword 验证密码
func (u *PassportUsecase) VerifyPassword(ctx context.Context, user *do.User) (verified bool, uid int64, err error) {
	pwd := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("verify password error(%v)", err)
		return false, 0, err
	}
	us, err := u.repo.GetUserByName(ctx, user.Name)
	if err != nil {
		u.log.Errorf("get user by name error(%v)", err)
		return false, 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(us.EncryptedPassword), hashedPassword)
	if err != nil {
		return false, 0, err
	}
	return true, us.ID, nil
}

// GetUserByID 通过id获取用户信息
func (u *PassportUsecase) GetUserByID(ctx context.Context, userId int64) (*do.User, error) {
	res, err, _ := u.sf.Do(constants.SFUserInfoKey(userId), func() (interface{}, error) {
		return u.repo.GetUserById(ctx, userId)
	})
	if err != nil {
		u.log.Errorf("get user by id(%v) error(%v)", userId, err)
		return nil, err
	}
	return res.(*do.User), nil
}

// MGetUserByID 批量获取用户信息
func (u *PassportUsecase) MGetUserByID(ctx context.Context, userIds []int64) ([]*do.User, error) {
	return u.repo.MGetUserById(ctx, userIds)
}
