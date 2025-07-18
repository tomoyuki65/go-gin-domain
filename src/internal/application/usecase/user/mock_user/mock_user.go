// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/application/usecase/user/user.go
//
// Generated by this command:
//
//	mockgen -source=./internal/application/usecase/user/user.go -destination=./internal/application/usecase/user/mock_user/mock_user.go
//

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	user "go-gin-domain/internal/domain/user"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
	isgomock struct{}
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserUsecase) Create(ctx context.Context, lastName, firstName, email string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, lastName, firstName, email)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserUsecaseMockRecorder) Create(ctx, lastName, firstName, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserUsecase)(nil).Create), ctx, lastName, firstName, email)
}

// Delete mocks base method.
func (m *MockUserUsecase) Delete(ctx context.Context, uid string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, uid)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserUsecaseMockRecorder) Delete(ctx, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserUsecase)(nil).Delete), ctx, uid)
}

// FindAll mocks base method.
func (m *MockUserUsecase) FindAll(ctx context.Context) ([]*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserUsecaseMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserUsecase)(nil).FindAll), ctx)
}

// FindByUID mocks base method.
func (m *MockUserUsecase) FindByUID(ctx context.Context, uid string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUID", ctx, uid)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUID indicates an expected call of FindByUID.
func (mr *MockUserUsecaseMockRecorder) FindByUID(ctx, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUID", reflect.TypeOf((*MockUserUsecase)(nil).FindByUID), ctx, uid)
}

// Update mocks base method.
func (m *MockUserUsecase) Update(ctx context.Context, uid, lastName, firstName, email string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, uid, lastName, firstName, email)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserUsecaseMockRecorder) Update(ctx, uid, lastName, firstName, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserUsecase)(nil).Update), ctx, uid, lastName, firstName, email)
}
