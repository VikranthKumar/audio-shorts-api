// Code generated by MockGen. DO NOT EDIT.
// Source: shorts.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nooble/task/audio-short-api/pkg/api/model"
)

// MockAudioShortsStore is a mock of AudioShortsStore interface.
type MockAudioShortsStore struct {
	ctrl     *gomock.Controller
	recorder *MockAudioShortsStoreMockRecorder
}

// MockAudioShortsStoreMockRecorder is the mock recorder for MockAudioShortsStore.
type MockAudioShortsStoreMockRecorder struct {
	mock *MockAudioShortsStore
}

// NewMockAudioShortsStore creates a new mock instance.
func NewMockAudioShortsStore(ctrl *gomock.Controller) *MockAudioShortsStore {
	mock := &MockAudioShortsStore{ctrl: ctrl}
	mock.recorder = &MockAudioShortsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAudioShortsStore) EXPECT() *MockAudioShortsStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAudioShortsStore) Create(ctx context.Context, input *model.AudioShortInput) (*model.AudioShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*model.AudioShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAudioShortsStoreMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAudioShortsStore)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockAudioShortsStore) Delete(ctx context.Context, id string) (*model.AudioShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(*model.AudioShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAudioShortsStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAudioShortsStore)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *MockAudioShortsStore) GetAll(ctx context.Context, page, limit uint16) ([]*model.AudioShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page, limit)
	ret0, _ := ret[0].([]*model.AudioShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAudioShortsStoreMockRecorder) GetAll(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAudioShortsStore)(nil).GetAll), ctx, page, limit)
}

// GetByID mocks base method.
func (m *MockAudioShortsStore) GetByID(ctx context.Context, id string) (*model.AudioShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*model.AudioShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAudioShortsStoreMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAudioShortsStore)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *MockAudioShortsStore) Update(ctx context.Context, id string, input *model.AudioShortInput) (*model.AudioShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, input)
	ret0, _ := ret[0].(*model.AudioShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAudioShortsStoreMockRecorder) Update(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAudioShortsStore)(nil).Update), ctx, id, input)
}