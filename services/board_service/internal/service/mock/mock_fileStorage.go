// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/services/board_service/internal/fileStorage (interfaces: FileStorage)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFileStorage is a mock of FileStorage interface
type MockFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageMockRecorder
}

// MockFileStorageMockRecorder is the mock recorder for MockFileStorage
type MockFileStorageMockRecorder struct {
	mock *MockFileStorage
}

// NewMockFileStorage creates a new mock instance
func NewMockFileStorage(ctrl *gomock.Controller) *MockFileStorage {
	mock := &MockFileStorage{ctrl: ctrl}
	mock.recorder = &MockFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStorage) EXPECT() *MockFileStorageMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method
func (m *MockFileStorage) DeleteFile(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockFileStorageMockRecorder) DeleteFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileStorage)(nil).DeleteFile), arg0, arg1)
}

// UploadFile mocks base method
func (m *MockFileStorage) UploadFile(arg0 models.AttachmentFileInternal, arg1 *models.AttachmentInternal, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockFileStorageMockRecorder) UploadFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileStorage)(nil).UploadFile), arg0, arg1, arg2)
}
