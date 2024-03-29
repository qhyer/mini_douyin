// Code generated by MockGen. DO NOT EDIT.
// Source: favorite.go
//
// Generated by this command:
//
//	mockgen -source=favorite.go -destination=mock/favorite.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	v1 "douyin/api/video/favorite/service/v1"
	v10 "douyin/api/video/feed/service/v1"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFavoriteRepo is a mock of FavoriteRepo interface.
type MockFavoriteRepo struct {
	ctrl     *gomock.Controller
	recorder *MockFavoriteRepoMockRecorder
}

// MockFavoriteRepoMockRecorder is the mock recorder for MockFavoriteRepo.
type MockFavoriteRepoMockRecorder struct {
	mock *MockFavoriteRepo
}

// NewMockFavoriteRepo creates a new mock instance.
func NewMockFavoriteRepo(ctrl *gomock.Controller) *MockFavoriteRepo {
	mock := &MockFavoriteRepo{ctrl: ctrl}
	mock.recorder = &MockFavoriteRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFavoriteRepo) EXPECT() *MockFavoriteRepoMockRecorder {
	return m.recorder
}

// FavoriteAction mocks base method.
func (m *MockFavoriteRepo) FavoriteAction(ctx context.Context, userId, videoId int64, actionType int32) (*v1.DouyinFavoriteActionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteAction", ctx, userId, videoId, actionType)
	ret0, _ := ret[0].(*v1.DouyinFavoriteActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FavoriteAction indicates an expected call of FavoriteAction.
func (mr *MockFavoriteRepoMockRecorder) FavoriteAction(ctx, userId, videoId, actionType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteAction", reflect.TypeOf((*MockFavoriteRepo)(nil).FavoriteAction), ctx, userId, videoId, actionType)
}

// GetUserFavoriteVideoList mocks base method.
func (m *MockFavoriteRepo) GetUserFavoriteVideoList(ctx context.Context, userId, toUserId int64) (*v10.GetUserFavoriteVideoListByUserIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavoriteVideoList", ctx, userId, toUserId)
	ret0, _ := ret[0].(*v10.GetUserFavoriteVideoListByUserIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFavoriteVideoList indicates an expected call of GetUserFavoriteVideoList.
func (mr *MockFavoriteRepoMockRecorder) GetUserFavoriteVideoList(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavoriteVideoList", reflect.TypeOf((*MockFavoriteRepo)(nil).GetUserFavoriteVideoList), ctx, userId, toUserId)
}
