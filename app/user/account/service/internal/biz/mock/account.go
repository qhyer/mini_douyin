// Code generated by MockGen. DO NOT EDIT.
// Source: account.go
//
// Generated by this command:
//
//	mockgen -source=account.go -destination=mock/account.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	entity "douyin/app/user/account/common/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepo is a mock of AccountRepo interface.
type MockAccountRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepoMockRecorder
}

// MockAccountRepoMockRecorder is the mock recorder for MockAccountRepo.
type MockAccountRepoMockRecorder struct {
	mock *MockAccountRepo
}

// NewMockAccountRepo creates a new mock instance.
func NewMockAccountRepo(ctrl *gomock.Controller) *MockAccountRepo {
	mock := &MockAccountRepo{ctrl: ctrl}
	mock.recorder = &MockAccountRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepo) EXPECT() *MockAccountRepoMockRecorder {
	return m.recorder
}

// GetFollowListByUserId mocks base method.
func (m *MockAccountRepo) GetFollowListByUserId(ctx context.Context, userId, toUserId int64) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowListByUserId", ctx, userId, toUserId)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowListByUserId indicates an expected call of GetFollowListByUserId.
func (mr *MockAccountRepoMockRecorder) GetFollowListByUserId(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowListByUserId", reflect.TypeOf((*MockAccountRepo)(nil).GetFollowListByUserId), ctx, userId, toUserId)
}

// GetFollowerListByUserId mocks base method.
func (m *MockAccountRepo) GetFollowerListByUserId(ctx context.Context, userId, toUserId int64) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowerListByUserId", ctx, userId, toUserId)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowerListByUserId indicates an expected call of GetFollowerListByUserId.
func (mr *MockAccountRepoMockRecorder) GetFollowerListByUserId(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowerListByUserId", reflect.TypeOf((*MockAccountRepo)(nil).GetFollowerListByUserId), ctx, userId, toUserId)
}

// GetFriendListByUserId mocks base method.
func (m *MockAccountRepo) GetFriendListByUserId(ctx context.Context, userId int64) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendListByUserId", ctx, userId)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendListByUserId indicates an expected call of GetFriendListByUserId.
func (mr *MockAccountRepoMockRecorder) GetFriendListByUserId(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendListByUserId", reflect.TypeOf((*MockAccountRepo)(nil).GetFriendListByUserId), ctx, userId)
}

// GetUserInfoByUserId mocks base method.
func (m *MockAccountRepo) GetUserInfoByUserId(ctx context.Context, userId, toUserId int64) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoByUserId", ctx, userId, toUserId)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoByUserId indicates an expected call of GetUserInfoByUserId.
func (mr *MockAccountRepoMockRecorder) GetUserInfoByUserId(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoByUserId", reflect.TypeOf((*MockAccountRepo)(nil).GetUserInfoByUserId), ctx, userId, toUserId)
}

// MGetUserInfoByUserId mocks base method.
func (m *MockAccountRepo) MGetUserInfoByUserId(ctx context.Context, toUserIds []int64) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MGetUserInfoByUserId", ctx, toUserIds)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGetUserInfoByUserId indicates an expected call of MGetUserInfoByUserId.
func (mr *MockAccountRepoMockRecorder) MGetUserInfoByUserId(ctx, toUserIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGetUserInfoByUserId", reflect.TypeOf((*MockAccountRepo)(nil).MGetUserInfoByUserId), ctx, toUserIds)
}
