// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/services/profile (interfaces: AvatarStorage)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	reflect "reflect"
)

// MockAvatarStorage is a mock of AvatarStorage interface
type MockAvatarStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAvatarStorageMockRecorder
}

// MockAvatarStorageMockRecorder is the mock recorder for MockAvatarStorage
type MockAvatarStorageMockRecorder struct {
	mock *MockAvatarStorage
}

// NewMockAvatarStorage creates a new mock instance
func NewMockAvatarStorage(ctrl *gomock.Controller) *MockAvatarStorage {
	mock := &MockAvatarStorage{ctrl: ctrl}
	mock.recorder = &MockAvatarStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAvatarStorage) EXPECT() *MockAvatarStorageMockRecorder {
	return m.recorder
}

// UploadAvatar mocks base method
func (m *MockAvatarStorage) UploadAvatar(arg0 *multipart.FileHeader, arg1 *models.UserAvatar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatar indicates an expected call of UploadAvatar
func (mr *MockAvatarStorageMockRecorder) UploadAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockAvatarStorage)(nil).UploadAvatar), arg0, arg1)
}