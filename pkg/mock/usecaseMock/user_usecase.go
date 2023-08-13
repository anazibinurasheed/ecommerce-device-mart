// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	response "github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddNewAddress mocks base method.
func (m *MockUserUseCase) AddNewAddress(userId int, address request.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewAddress", userId, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewAddress indicates an expected call of AddNewAddress.
func (mr *MockUserUseCaseMockRecorder) AddNewAddress(userId, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddNewAddress), userId, address)
}

// ChangeUserPassword mocks base method.
func (m *MockUserUseCase) ChangeUserPassword(password request.ChangePassword, userId int, c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", password, userId, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockUserUseCaseMockRecorder) ChangeUserPassword(password, userId, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserUseCase)(nil).ChangeUserPassword), password, userId, c)
}

// CheckUserOldPassword mocks base method.
func (m *MockUserUseCase) CheckUserOldPassword(password request.OldPassword, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserOldPassword", password, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserOldPassword indicates an expected call of CheckUserOldPassword.
func (mr *MockUserUseCaseMockRecorder) CheckUserOldPassword(password, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserOldPassword", reflect.TypeOf((*MockUserUseCase)(nil).CheckUserOldPassword), password, userId)
}

// DeleteUserAddress mocks base method.
func (m *MockUserUseCase) DeleteUserAddress(addressId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserAddress", addressId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserAddress indicates an expected call of DeleteUserAddress.
func (mr *MockUserUseCaseMockRecorder) DeleteUserAddress(addressId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserAddress", reflect.TypeOf((*MockUserUseCase)(nil).DeleteUserAddress), addressId)
}

// DisplayListOfStates mocks base method.
func (m *MockUserUseCase) DisplayListOfStates() ([]response.States, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayListOfStates")
	ret0, _ := ret[0].([]response.States)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayListOfStates indicates an expected call of DisplayListOfStates.
func (mr *MockUserUseCaseMockRecorder) DisplayListOfStates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayListOfStates", reflect.TypeOf((*MockUserUseCase)(nil).DisplayListOfStates))
}

// FindUserById mocks base method.
func (m *MockUserUseCase) FindUserById(id int) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserUseCaseMockRecorder) FindUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserUseCase)(nil).FindUserById), id)
}

// ForgotPassword mocks base method.
func (m *MockUserUseCase) ForgotPassword(userId int, c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", userId, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserUseCaseMockRecorder) ForgotPassword(userId, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserUseCase)(nil).ForgotPassword), userId, c)
}

// GetProfile mocks base method.
func (m *MockUserUseCase) GetProfile(userId int) (response.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", userId)
	ret0, _ := ret[0].(response.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockUserUseCaseMockRecorder) GetProfile(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockUserUseCase)(nil).GetProfile), userId)
}

// GetUserAddresses mocks base method.
func (m *MockUserUseCase) GetUserAddresses(userId int) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAddresses", userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAddresses indicates an expected call of GetUserAddresses.
func (mr *MockUserUseCaseMockRecorder) GetUserAddresses(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAddresses", reflect.TypeOf((*MockUserUseCase)(nil).GetUserAddresses), userId)
}

// SetDefaultAddress mocks base method.
func (m *MockUserUseCase) SetDefaultAddress(userID, addressID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultAddress", userID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultAddress indicates an expected call of SetDefaultAddress.
func (mr *MockUserUseCaseMockRecorder) SetDefaultAddress(userID, addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultAddress", reflect.TypeOf((*MockUserUseCase)(nil).SetDefaultAddress), userID, addressID)
}

// SignUp mocks base method.
func (m *MockUserUseCase) SignUp(user request.SignUpData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserUseCaseMockRecorder) SignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserUseCase)(nil).SignUp), user)
}

// UpdateUserAddress mocks base method.
func (m *MockUserUseCase) UpdateUserAddress(address request.Address, addressID, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAddress", address, addressID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAddress indicates an expected call of UpdateUserAddress.
func (mr *MockUserUseCaseMockRecorder) UpdateUserAddress(address, addressID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAddress", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUserAddress), address, addressID, userID)
}

// UpdateUserName mocks base method.
func (m *MockUserUseCase) UpdateUserName(username string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserName", username, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserName indicates an expected call of UpdateUserName.
func (mr *MockUserUseCaseMockRecorder) UpdateUserName(username, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserName", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUserName), username, userID)
}

// ValidateUserLoginCredentials mocks base method.
func (m *MockUserUseCase) ValidateUserLoginCredentials(user request.LoginData) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserLoginCredentials", user)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUserLoginCredentials indicates an expected call of ValidateUserLoginCredentials.
func (mr *MockUserUseCaseMockRecorder) ValidateUserLoginCredentials(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserLoginCredentials", reflect.TypeOf((*MockUserUseCase)(nil).ValidateUserLoginCredentials), user)
}