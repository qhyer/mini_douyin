package biz

import (
	do "douyin/app/user/passport/common/entity"
	mock_biz "douyin/app/user/passport/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPassportUsecase_CreateUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockPassportRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil),
	)

	uc := NewPassportUseCase(repo, nil)
	_, err := uc.CreateUser(nil, &do.User{
		ID:                0,
		Name:              "test",
		Password:          "test",
		EncryptedPassword: "test",
		Avatar:            "test",
		BackgroundImage:   "test",
		Signature:         "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPassportUsecase_VerifyPassword(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockPassportRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetUserByName(gomock.Any(), gomock.Any()).Return(&do.User{
			ID:                0,
			Name:              "test",
			Password:          "test",
			EncryptedPassword: "$2a$12$Wit9DXVWgvplnCWmoudQyexpbiz2Ekt4DK1cNKSIuivMxEKGcEEAy",
			Avatar:            "test",
			BackgroundImage:   "test",
			Signature:         "test",
		}, nil),
	)

	uc := NewPassportUseCase(repo, nil)
	_, _, err := uc.VerifyPassword(nil, &do.User{
		ID:              0,
		Name:            "test",
		Password:        "123456",
		Avatar:          "test",
		BackgroundImage: "test",
		Signature:       "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPassportUsecase_GetUserByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockPassportRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetUserById(gomock.Any(), gomock.Any()).Return(&do.User{
			ID:                0,
			Name:              "test",
			Password:          "test",
			EncryptedPassword: "test",
			Avatar:            "test",
			BackgroundImage:   "test",
			Signature:         "test",
		}, nil),
	)

	uc := NewPassportUseCase(repo, nil)
	_, err := uc.GetUserByID(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPassportUsecase_MGetUserByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockPassportRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().MGetUserById(gomock.Any(), gomock.Any()).Return([]*do.User{
			{
				ID:                0,
				Name:              "test",
				Password:          "test",
				EncryptedPassword: "test",
				Avatar:            "test",
				BackgroundImage:   "test",
				Signature:         "test",
			},
		}, nil),
	)

	uc := NewPassportUseCase(repo, nil)
	_, err := uc.MGetUserByID(nil, []int64{0})
	if err != nil {
		t.Fatal(err)
	}
}
