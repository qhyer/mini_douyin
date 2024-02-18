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
	event "douyin/app/video/favorite/common/event"
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

// CountUserFavoriteByUserId mocks base method.
func (m *MockFavoriteRepo) CountUserFavoriteByUserId(ctx context.Context, userId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserFavoriteByUserId", ctx, userId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserFavoriteByUserId indicates an expected call of CountUserFavoriteByUserId.
func (mr *MockFavoriteRepoMockRecorder) CountUserFavoriteByUserId(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserFavoriteByUserId", reflect.TypeOf((*MockFavoriteRepo)(nil).CountUserFavoriteByUserId), ctx, userId)
}

// CountUserFavoritedByUserId mocks base method.
func (m *MockFavoriteRepo) CountUserFavoritedByUserId(ctx context.Context, userId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserFavoritedByUserId", ctx, userId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserFavoritedByUserId indicates an expected call of CountUserFavoritedByUserId.
func (mr *MockFavoriteRepoMockRecorder) CountUserFavoritedByUserId(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserFavoritedByUserId", reflect.TypeOf((*MockFavoriteRepo)(nil).CountUserFavoritedByUserId), ctx, userId)
}

// CountVideoFavoritedByVideoId mocks base method.
func (m *MockFavoriteRepo) CountVideoFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountVideoFavoritedByVideoId", ctx, videoId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountVideoFavoritedByVideoId indicates an expected call of CountVideoFavoritedByVideoId.
func (mr *MockFavoriteRepoMockRecorder) CountVideoFavoritedByVideoId(ctx, videoId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountVideoFavoritedByVideoId", reflect.TypeOf((*MockFavoriteRepo)(nil).CountVideoFavoritedByVideoId), ctx, videoId)
}

// FavoriteVideo mocks base method.
func (m *MockFavoriteRepo) FavoriteVideo(ctx context.Context, fav *event.FavoriteAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FavoriteVideo", ctx, fav)
	ret0, _ := ret[0].(error)
	return ret0
}

// FavoriteVideo indicates an expected call of FavoriteVideo.
func (mr *MockFavoriteRepoMockRecorder) FavoriteVideo(ctx, fav any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FavoriteVideo", reflect.TypeOf((*MockFavoriteRepo)(nil).FavoriteVideo), ctx, fav)
}

// GetFavoriteVideoIdListByUserId mocks base method.
func (m *MockFavoriteRepo) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteVideoIdListByUserId", ctx, userId)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteVideoIdListByUserId indicates an expected call of GetFavoriteVideoIdListByUserId.
func (mr *MockFavoriteRepoMockRecorder) GetFavoriteVideoIdListByUserId(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteVideoIdListByUserId", reflect.TypeOf((*MockFavoriteRepo)(nil).GetFavoriteVideoIdListByUserId), ctx, userId)
}

// IsUserFavoriteVideo mocks base method.
func (m *MockFavoriteRepo) IsUserFavoriteVideo(ctx context.Context, userId, videoId int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserFavoriteVideo", ctx, userId, videoId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserFavoriteVideo indicates an expected call of IsUserFavoriteVideo.
func (mr *MockFavoriteRepoMockRecorder) IsUserFavoriteVideo(ctx, userId, videoId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserFavoriteVideo", reflect.TypeOf((*MockFavoriteRepo)(nil).IsUserFavoriteVideo), ctx, userId, videoId)
}

// IsUserFavoriteVideoList mocks base method.
func (m *MockFavoriteRepo) IsUserFavoriteVideoList(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserFavoriteVideoList", ctx, userId, videoIds)
	ret0, _ := ret[0].([]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserFavoriteVideoList indicates an expected call of IsUserFavoriteVideoList.
func (mr *MockFavoriteRepoMockRecorder) IsUserFavoriteVideoList(ctx, userId, videoIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserFavoriteVideoList", reflect.TypeOf((*MockFavoriteRepo)(nil).IsUserFavoriteVideoList), ctx, userId, videoIds)
}

// MCountVideoFavoritedByVideoId mocks base method.
func (m *MockFavoriteRepo) MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MCountVideoFavoritedByVideoId", ctx, videoId)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MCountVideoFavoritedByVideoId indicates an expected call of MCountVideoFavoritedByVideoId.
func (mr *MockFavoriteRepoMockRecorder) MCountVideoFavoritedByVideoId(ctx, videoId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MCountVideoFavoritedByVideoId", reflect.TypeOf((*MockFavoriteRepo)(nil).MCountVideoFavoritedByVideoId), ctx, videoId)
}