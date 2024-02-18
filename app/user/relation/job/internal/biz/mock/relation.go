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
	event "douyin/app/user/relation/common/event"
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

// BatchUpdateUserRelationStat mocks base method.
func (m *MockRelationRepo) BatchUpdateUserRelationStat(ctx context.Context, follow, follower map[int64]int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchUpdateUserRelationStat", ctx, follow, follower)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchUpdateUserRelationStat indicates an expected call of BatchUpdateUserRelationStat.
func (mr *MockRelationRepoMockRecorder) BatchUpdateUserRelationStat(ctx, follow, follower any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdateUserRelationStat", reflect.TypeOf((*MockRelationRepo)(nil).BatchUpdateUserRelationStat), ctx, follow, follower)
}

// CreateRelation mocks base method.
func (m *MockRelationRepo) CreateRelation(ctx context.Context, relation *event.RelationAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelation", ctx, relation)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRelation indicates an expected call of CreateRelation.
func (mr *MockRelationRepoMockRecorder) CreateRelation(ctx, relation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelation", reflect.TypeOf((*MockRelationRepo)(nil).CreateRelation), ctx, relation)
}

// DeleteRelation mocks base method.
func (m *MockRelationRepo) DeleteRelation(ctx context.Context, relation *event.RelationAction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRelation", ctx, relation)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRelation indicates an expected call of DeleteRelation.
func (mr *MockRelationRepoMockRecorder) DeleteRelation(ctx, relation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRelation", reflect.TypeOf((*MockRelationRepo)(nil).DeleteRelation), ctx, relation)
}

// GetUserFollowTempCount mocks base method.
func (m *MockRelationRepo) GetUserFollowTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFollowTempCount", ctx, procId)
	ret0, _ := ret[0].(map[int64]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFollowTempCount indicates an expected call of GetUserFollowTempCount.
func (mr *MockRelationRepoMockRecorder) GetUserFollowTempCount(ctx, procId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFollowTempCount", reflect.TypeOf((*MockRelationRepo)(nil).GetUserFollowTempCount), ctx, procId)
}

// GetUserFollowerTempCount mocks base method.
func (m *MockRelationRepo) GetUserFollowerTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFollowerTempCount", ctx, procId)
	ret0, _ := ret[0].(map[int64]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFollowerTempCount indicates an expected call of GetUserFollowerTempCount.
func (mr *MockRelationRepoMockRecorder) GetUserFollowerTempCount(ctx, procId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFollowerTempCount", reflect.TypeOf((*MockRelationRepo)(nil).GetUserFollowerTempCount), ctx, procId)
}

// PurgeUserFollowTempCount mocks base method.
func (m *MockRelationRepo) PurgeUserFollowTempCount(ctx context.Context, procId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeUserFollowTempCount", ctx, procId)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeUserFollowTempCount indicates an expected call of PurgeUserFollowTempCount.
func (mr *MockRelationRepoMockRecorder) PurgeUserFollowTempCount(ctx, procId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeUserFollowTempCount", reflect.TypeOf((*MockRelationRepo)(nil).PurgeUserFollowTempCount), ctx, procId)
}

// PurgeUserFollowerTempCount mocks base method.
func (m *MockRelationRepo) PurgeUserFollowerTempCount(ctx context.Context, procId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeUserFollowerTempCount", ctx, procId)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeUserFollowerTempCount indicates an expected call of PurgeUserFollowerTempCount.
func (mr *MockRelationRepoMockRecorder) PurgeUserFollowerTempCount(ctx, procId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeUserFollowerTempCount", reflect.TypeOf((*MockRelationRepo)(nil).PurgeUserFollowerTempCount), ctx, procId)
}

// UpdateUserFollowCount mocks base method.
func (m *MockRelationRepo) UpdateUserFollowCount(ctx context.Context, userId, incr int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserFollowCount", ctx, userId, incr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserFollowCount indicates an expected call of UpdateUserFollowCount.
func (mr *MockRelationRepoMockRecorder) UpdateUserFollowCount(ctx, userId, incr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserFollowCount", reflect.TypeOf((*MockRelationRepo)(nil).UpdateUserFollowCount), ctx, userId, incr)
}

// UpdateUserFollowTempCount mocks base method.
func (m *MockRelationRepo) UpdateUserFollowTempCount(ctx context.Context, procId int, userId, incr int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserFollowTempCount", ctx, procId, userId, incr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserFollowTempCount indicates an expected call of UpdateUserFollowTempCount.
func (mr *MockRelationRepoMockRecorder) UpdateUserFollowTempCount(ctx, procId, userId, incr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserFollowTempCount", reflect.TypeOf((*MockRelationRepo)(nil).UpdateUserFollowTempCount), ctx, procId, userId, incr)
}

// UpdateUserFollowerCount mocks base method.
func (m *MockRelationRepo) UpdateUserFollowerCount(ctx context.Context, userId, incr int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserFollowerCount", ctx, userId, incr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserFollowerCount indicates an expected call of UpdateUserFollowerCount.
func (mr *MockRelationRepoMockRecorder) UpdateUserFollowerCount(ctx, userId, incr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserFollowerCount", reflect.TypeOf((*MockRelationRepo)(nil).UpdateUserFollowerCount), ctx, userId, incr)
}

// UpdateUserFollowerTempCount mocks base method.
func (m *MockRelationRepo) UpdateUserFollowerTempCount(ctx context.Context, procId int, userId, incr int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserFollowerTempCount", ctx, procId, userId, incr)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserFollowerTempCount indicates an expected call of UpdateUserFollowerTempCount.
func (mr *MockRelationRepoMockRecorder) UpdateUserFollowerTempCount(ctx, procId, userId, incr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserFollowerTempCount", reflect.TypeOf((*MockRelationRepo)(nil).UpdateUserFollowerTempCount), ctx, procId, userId, incr)
}