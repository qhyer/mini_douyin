// Code generated by MockGen. DO NOT EDIT.
// Source: publish.go
//
// Generated by this command:
//
//	mockgen -source=publish.go -destination=mock/publish.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	v1 "douyin/api/video/feed/service/v1"
	v10 "douyin/api/video/publish/service/v1"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPublishRepo is a mock of PublishRepo interface.
type MockPublishRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPublishRepoMockRecorder
}

// MockPublishRepoMockRecorder is the mock recorder for MockPublishRepo.
type MockPublishRepoMockRecorder struct {
	mock *MockPublishRepo
}

// NewMockPublishRepo creates a new mock instance.
func NewMockPublishRepo(ctrl *gomock.Controller) *MockPublishRepo {
	mock := &MockPublishRepo{ctrl: ctrl}
	mock.recorder = &MockPublishRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublishRepo) EXPECT() *MockPublishRepoMockRecorder {
	return m.recorder
}

// GetUserPublishedVideoList mocks base method.
func (m *MockPublishRepo) GetUserPublishedVideoList(ctx context.Context, userId, toUserId int64) (*v1.GetPublishedVideoByUserIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPublishedVideoList", ctx, userId, toUserId)
	ret0, _ := ret[0].(*v1.GetPublishedVideoByUserIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPublishedVideoList indicates an expected call of GetUserPublishedVideoList.
func (mr *MockPublishRepoMockRecorder) GetUserPublishedVideoList(ctx, userId, toUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPublishedVideoList", reflect.TypeOf((*MockPublishRepo)(nil).GetUserPublishedVideoList), ctx, userId, toUserId)
}

// PublishVideo mocks base method.
func (m *MockPublishRepo) PublishVideo(ctx context.Context, userId int64, title string, data []byte) (*v10.PublishActionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishVideo", ctx, userId, title, data)
	ret0, _ := ret[0].(*v10.PublishActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishVideo indicates an expected call of PublishVideo.
func (mr *MockPublishRepoMockRecorder) PublishVideo(ctx, userId, title, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishVideo", reflect.TypeOf((*MockPublishRepo)(nil).PublishVideo), ctx, userId, title, data)
}
