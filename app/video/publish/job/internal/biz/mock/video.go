// Code generated by MockGen. DO NOT EDIT.
// Source: video.go
//
// Generated by this command:
//
//	mockgen -source=video.go -destination=mock/video.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	entity "douyin/app/video/publish/common/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockVideoRepo is a mock of VideoRepo interface.
type MockVideoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockVideoRepoMockRecorder
}

// MockVideoRepoMockRecorder is the mock recorder for MockVideoRepo.
type MockVideoRepoMockRecorder struct {
	mock *MockVideoRepo
}

// NewMockVideoRepo creates a new mock instance.
func NewMockVideoRepo(ctrl *gomock.Controller) *MockVideoRepo {
	mock := &MockVideoRepo{ctrl: ctrl}
	mock.recorder = &MockVideoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoRepo) EXPECT() *MockVideoRepoMockRecorder {
	return m.recorder
}

// BatchCreateVideo mocks base method.
func (m *MockVideoRepo) BatchCreateVideo(ctx context.Context, videos []*entity.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateVideo", ctx, videos)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateVideo indicates an expected call of BatchCreateVideo.
func (mr *MockVideoRepoMockRecorder) BatchCreateVideo(ctx, videos any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateVideo", reflect.TypeOf((*MockVideoRepo)(nil).BatchCreateVideo), ctx, videos)
}

// CreateVideo mocks base method.
func (m *MockVideoRepo) CreateVideo(ctx context.Context, video *entity.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVideo", ctx, video)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVideo indicates an expected call of CreateVideo.
func (mr *MockVideoRepoMockRecorder) CreateVideo(ctx, video any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVideo", reflect.TypeOf((*MockVideoRepo)(nil).CreateVideo), ctx, video)
}

// GetVideo mocks base method.
func (m *MockVideoRepo) GetVideo(ctx context.Context, objectName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideo", ctx, objectName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideo indicates an expected call of GetVideo.
func (mr *MockVideoRepoMockRecorder) GetVideo(ctx, objectName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideo", reflect.TypeOf((*MockVideoRepo)(nil).GetVideo), ctx, objectName)
}

// UploadCover mocks base method.
func (m *MockVideoRepo) UploadCover(ctx context.Context, data []byte, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadCover", ctx, data, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadCover indicates an expected call of UploadCover.
func (mr *MockVideoRepoMockRecorder) UploadCover(ctx, data, objectName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadCover", reflect.TypeOf((*MockVideoRepo)(nil).UploadCover), ctx, data, objectName)
}
