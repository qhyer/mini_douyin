// Code generated by MockGen. DO NOT EDIT.
// Source: chat.go
//
// Generated by this command:
//
//	mockgen -source=chat.go -destination=mock/chat.go
//

// Package mock_biz is a generated GoMock package.
package mock_biz

import (
	context "context"
	entity "douyin/app/user/chat/common/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockChatRepo is a mock of ChatRepo interface.
type MockChatRepo struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepoMockRecorder
}

// MockChatRepoMockRecorder is the mock recorder for MockChatRepo.
type MockChatRepoMockRecorder struct {
	mock *MockChatRepo
}

// NewMockChatRepo creates a new mock instance.
func NewMockChatRepo(ctrl *gomock.Controller) *MockChatRepo {
	mock := &MockChatRepo{ctrl: ctrl}
	mock.recorder = &MockChatRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRepo) EXPECT() *MockChatRepoMockRecorder {
	return m.recorder
}

// GetLatestMsgByMyUserIdAndHisUserId mocks base method.
func (m *MockChatRepo) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, myUserId, hisUserId int64) (*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestMsgByMyUserIdAndHisUserId", ctx, myUserId, hisUserId)
	ret0, _ := ret[0].(*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestMsgByMyUserIdAndHisUserId indicates an expected call of GetLatestMsgByMyUserIdAndHisUserId.
func (mr *MockChatRepoMockRecorder) GetLatestMsgByMyUserIdAndHisUserId(ctx, myUserId, hisUserId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestMsgByMyUserIdAndHisUserId", reflect.TypeOf((*MockChatRepo)(nil).GetLatestMsgByMyUserIdAndHisUserId), ctx, myUserId, hisUserId)
}

// GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime mocks base method.
func (m *MockChatRepo) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int) ([]*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime", ctx, myUserId, hisUserId, preMsgTime, limit)
	ret0, _ := ret[0].([]*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime indicates an expected call of GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime.
func (mr *MockChatRepoMockRecorder) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx, myUserId, hisUserId, preMsgTime, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime", reflect.TypeOf((*MockChatRepo)(nil).GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime), ctx, myUserId, hisUserId, preMsgTime, limit)
}

// SendMessage mocks base method.
func (m *MockChatRepo) SendMessage(ctx context.Context, message *entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockChatRepoMockRecorder) SendMessage(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockChatRepo)(nil).SendMessage), ctx, message)
}
