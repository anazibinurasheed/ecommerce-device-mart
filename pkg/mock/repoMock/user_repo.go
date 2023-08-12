// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	request "github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	response "github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAdressToDatabase mocks base method.
func (m *MockUserRepository) AddAdressToDatabase(userId int, address request.Address) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAdressToDatabase", userId, address)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAdressToDatabase indicates an expected call of AddAdressToDatabase.
func (mr *MockUserRepositoryMockRecorder) AddAdressToDatabase(userId, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdressToDatabase", reflect.TypeOf((*MockUserRepository)(nil).AddAdressToDatabase), userId, address)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(userId int, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", userId, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(userId, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), userId, newPassword)
}

// DeleteAddressFromDatabase mocks base method.
func (m *MockUserRepository) DeleteAddressFromDatabase(adressId int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddressFromDatabase", adressId)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAddressFromDatabase indicates an expected call of DeleteAddressFromDatabase.
func (mr *MockUserRepositoryMockRecorder) DeleteAddressFromDatabase(adressId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddressFromDatabase", reflect.TypeOf((*MockUserRepository)(nil).DeleteAddressFromDatabase), adressId)
}

// FindAddressByAddressID mocks base method.
func (m *MockUserRepository) FindAddressByAddressID(addressID int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddressByAddressID", addressID)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddressByAddressID indicates an expected call of FindAddressByAddressID.
func (mr *MockUserRepositoryMockRecorder) FindAddressByAddressID(addressID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddressByAddressID", reflect.TypeOf((*MockUserRepository)(nil).FindAddressByAddressID), addressID)
}

// FindDefaultAddressById mocks base method.
func (m *MockUserRepository) FindDefaultAddressById(userId int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDefaultAddressById", userId)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDefaultAddressById indicates an expected call of FindDefaultAddressById.
func (mr *MockUserRepositoryMockRecorder) FindDefaultAddressById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDefaultAddressById", reflect.TypeOf((*MockUserRepository)(nil).FindDefaultAddressById), userId)
}

// FindUserAddress mocks base method.
func (m *MockUserRepository) FindUserAddress(userID int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserAddress", userID)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserAddress indicates an expected call of FindUserAddress.
func (mr *MockUserRepositoryMockRecorder) FindUserAddress(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserAddress", reflect.TypeOf((*MockUserRepository)(nil).FindUserAddress), userID)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(email string) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), email)
}

// FindUserById mocks base method.
func (m *MockUserRepository) FindUserById(id int) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserRepositoryMockRecorder) FindUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserRepository)(nil).FindUserById), id)
}

// FindUserByPhone mocks base method.
func (m *MockUserRepository) FindUserByPhone(phone int) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByPhone", phone)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByPhone indicates an expected call of FindUserByPhone.
func (mr *MockUserRepositoryMockRecorder) FindUserByPhone(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByPhone", reflect.TypeOf((*MockUserRepository)(nil).FindUserByPhone), phone)
}

// GetAllUserAddresses mocks base method.
func (m *MockUserRepository) GetAllUserAddresses(userId int) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserAddresses", userId)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserAddresses indicates an expected call of GetAllUserAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAllUserAddresses(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAllUserAddresses), userId)
}

// GetListOfStates mocks base method.
func (m *MockUserRepository) GetListOfStates() ([]response.States, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListOfStates")
	ret0, _ := ret[0].([]response.States)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListOfStates indicates an expected call of GetListOfStates.
func (mr *MockUserRepositoryMockRecorder) GetListOfStates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListOfStates", reflect.TypeOf((*MockUserRepository)(nil).GetListOfStates))
}

// ReadCategory mocks base method.
func (m *MockUserRepository) ReadCategory() ([]response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCategory")
	ret0, _ := ret[0].([]response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCategory indicates an expected call of ReadCategory.
func (mr *MockUserRepositoryMockRecorder) ReadCategory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCategory", reflect.TypeOf((*MockUserRepository)(nil).ReadCategory))
}

// SaveUserOnDatabase mocks base method.
func (m *MockUserRepository) SaveUserOnDatabase(user request.SignUpData) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserOnDatabase", user)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveUserOnDatabase indicates an expected call of SaveUserOnDatabase.
func (mr *MockUserRepositoryMockRecorder) SaveUserOnDatabase(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserOnDatabase", reflect.TypeOf((*MockUserRepository)(nil).SaveUserOnDatabase), user)
}

// SetIsDefaultStatusOnAddress mocks base method.
func (m *MockUserRepository) SetIsDefaultStatusOnAddress(status bool, addressId, userId int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIsDefaultStatusOnAddress", status, addressId, userId)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetIsDefaultStatusOnAddress indicates an expected call of SetIsDefaultStatusOnAddress.
func (mr *MockUserRepositoryMockRecorder) SetIsDefaultStatusOnAddress(status, addressId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIsDefaultStatusOnAddress", reflect.TypeOf((*MockUserRepository)(nil).SetIsDefaultStatusOnAddress), status, addressId, userId)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(address request.Address, addressID, userID int) (response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", address, addressID, userID)
	ret0, _ := ret[0].(response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(address, addressID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), address, addressID, userID)
}

// UpdateUserName mocks base method.
func (m *MockUserRepository) UpdateUserName(name string, userID int) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserName", name, userID)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserName indicates an expected call of UpdateUserName.
func (mr *MockUserRepositoryMockRecorder) UpdateUserName(name, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserName", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserName), name, userID)
}