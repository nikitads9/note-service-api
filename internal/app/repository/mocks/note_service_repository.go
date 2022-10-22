// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nikitads9/note-service-api/internal/app/repository (interfaces: INoteRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nikitads9/note-service-api/internal/app/model"
)

// MockINoteRepository is a mock of INoteRepository interface.
type MockINoteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockINoteRepositoryMockRecorder
}

// MockINoteRepositoryMockRecorder is the mock recorder for MockINoteRepository.
type MockINoteRepositoryMockRecorder struct {
	mock *MockINoteRepository
}

// NewMockINoteRepository creates a new mock instance.
func NewMockINoteRepository(ctrl *gomock.Controller) *MockINoteRepository {
	mock := &MockINoteRepository{ctrl: ctrl}
	mock.recorder = &MockINoteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINoteRepository) EXPECT() *MockINoteRepositoryMockRecorder {
	return m.recorder
}

// AddNote mocks base method.
func (m *MockINoteRepository) AddNote(arg0 context.Context, arg1 *model.NoteInfo) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNote indicates an expected call of AddNote.
func (mr *MockINoteRepositoryMockRecorder) AddNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockINoteRepository)(nil).AddNote), arg0, arg1)
}

// GetList mocks base method.
func (m *MockINoteRepository) GetList(arg0 context.Context) ([]*model.NoteInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", arg0)
	ret0, _ := ret[0].([]*model.NoteInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockINoteRepositoryMockRecorder) GetList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockINoteRepository)(nil).GetList), arg0)
}

// GetNote mocks base method.
func (m *MockINoteRepository) GetNote(arg0 context.Context, arg1 int64) (*model.NoteInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", arg0, arg1)
	ret0, _ := ret[0].(*model.NoteInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockINoteRepositoryMockRecorder) GetNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockINoteRepository)(nil).GetNote), arg0, arg1)
}

// MultiAdd mocks base method.
func (m *MockINoteRepository) MultiAdd(arg0 context.Context, arg1 []*model.NoteInfo) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiAdd", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiAdd indicates an expected call of MultiAdd.
func (mr *MockINoteRepositoryMockRecorder) MultiAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiAdd", reflect.TypeOf((*MockINoteRepository)(nil).MultiAdd), arg0, arg1)
}

// RemoveNote mocks base method.
func (m *MockINoteRepository) RemoveNote(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNote", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveNote indicates an expected call of RemoveNote.
func (mr *MockINoteRepositoryMockRecorder) RemoveNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNote", reflect.TypeOf((*MockINoteRepository)(nil).RemoveNote), arg0, arg1)
}

// UpdateNote mocks base method.
func (m *MockINoteRepository) UpdateNote(arg0 context.Context, arg1 *model.NoteInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockINoteRepositoryMockRecorder) UpdateNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockINoteRepository)(nil).UpdateNote), arg0, arg1)
}
