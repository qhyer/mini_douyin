// Code generated by MockGen. DO NOT EDIT.
// Source: relation.go
//
// Generated by this command:
//
//	mockgen -source=relation.go -destination=mock/relation.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	v1 "douyin/api/user/account/service/v1"
	v10 "douyin/api/user/relation/service/v1"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRelationRepo is a mock of RelationRepo interface.
type MockRelationRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRelationRepoMockRecorder
}

// MockRelationRepoMockRecorder is the mock recorder for MockRelationRepo.
type MockRelationRepoMockRecorder struct {
	mock *MockRelationRepo
}

// NewMockRelationRepo creates a new mock instance.
func NewMockRelationRepo(ctrl *gomock.Controller) *MockRelationRepo {
	mock := &MockRelationRepo{ctrl: ctrl}
	mock.recorder = &MockRelationRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelationRepo) EXPECT() *MockRelationRepoMockRecorder {
	return m.recorder
}

// GetUserFollowList mocks base method.
func (m *MockRelationRepo) GetUserFollowList(ctx context.Context, userId, toUserId int64) (*v1.GetFollowListByUserIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFollowList", ctx, userId, toUserId)
	ret0, _ := ret[0].(*v1.GetFollowListByUserIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFollowList indicates an expected call of GetUserFollowList.
func (mr *MockRelationRepoMockRecorder) GetUserFollowList(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFollowList", reflect.TypeOf((*MockRelationRepo)(nil).GetUserFollowList), ctx, userId, toUserId)
}

// GetUserFollowerList mocks base method.
func (m *MockRelationRepo) GetUserFollowerList(ctx context.Context, userId, toUserId int64) (*v1.GetFollowerListByUserIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFollowerList", ctx, userId, toUserId)
	ret0, _ := ret[0].(*v1.GetFollowerListByUserIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFollowerList indicates an expected call of GetUserFollowerList.
func (mr *MockRelationRepoMockRecorder) GetUserFollowerList(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFollowerList", reflect.TypeOf((*MockRelationRepo)(nil).GetUserFollowerList), ctx, userId, toUserId)
}

// GetUserFriendList mocks base method.
func (m *MockRelationRepo) GetUserFriendList(ctx context.Context, userId int64) (*v1.GetFriendListByUserIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFriendList", ctx, userId)
	ret0, _ := ret[0].(*v1.GetFriendListByUserIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFriendList indicates an expected call of GetUserFriendList.
func (mr *MockRelationRepoMockRecorder) GetUserFriendList(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFriendList", reflect.TypeOf((*MockRelationRepo)(nil).GetUserFriendList), ctx, userId)
}

// RelationAction mocks base method.
func (m *MockRelationRepo) RelationAction(ctx context.Context, userId, toUserId int64, actionType int32) (*v10.RelationActionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationAction", ctx, userId, toUserId, actionType)
	ret0, _ := ret[0].(*v10.RelationActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationAction indicates an expected call of RelationAction.
func (mr *MockRelationRepoMockRecorder) RelationAction(ctx, userId, toUserId, actionType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationAction", reflect.TypeOf((*MockRelationRepo)(nil).RelationAction), ctx, userId, toUserId, actionType)
}
