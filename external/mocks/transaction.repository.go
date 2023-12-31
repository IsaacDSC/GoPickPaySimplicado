// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infra/contracts/transaction.contract.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockTransactionRepositoriesInterface is a mock of TransactionRepositoriesInterface interface.
type MockTransactionRepositoriesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoriesInterfaceMockRecorder
}

// MockTransactionRepositoriesInterfaceMockRecorder is the mock recorder for MockTransactionRepositoriesInterface.
type MockTransactionRepositoriesInterfaceMockRecorder struct {
	mock *MockTransactionRepositoriesInterface
}

// NewMockTransactionRepositoriesInterface creates a new mock instance.
func NewMockTransactionRepositoriesInterface(ctrl *gomock.Controller) *MockTransactionRepositoriesInterface {
	mock := &MockTransactionRepositoriesInterface{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoriesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepositoriesInterface) EXPECT() *MockTransactionRepositoriesInterfaceMockRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockTransactionRepositoriesInterface) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockTransactionRepositoriesInterfaceMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockTransactionRepositoriesInterface)(nil).Done))
}

// GetTransactionsByUserID mocks base method.
func (m *MockTransactionRepositoriesInterface) GetTransactionsByUserID(ctx context.Context, UserID uuid.UUID) ([]domain.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByUserID", ctx, UserID)
	ret0, _ := ret[0].([]domain.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByUserID indicates an expected call of GetTransactionsByUserID.
func (mr *MockTransactionRepositoriesInterfaceMockRecorder) GetTransactionsByUserID(ctx, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByUserID", reflect.TypeOf((*MockTransactionRepositoriesInterface)(nil).GetTransactionsByUserID), ctx, UserID)
}

// InsertTransaction mocks base method.
func (m *MockTransactionRepositoriesInterface) InsertTransaction(ctx context.Context, input domain.TransactionEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTransaction", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTransaction indicates an expected call of InsertTransaction.
func (mr *MockTransactionRepositoriesInterfaceMockRecorder) InsertTransaction(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockTransactionRepositoriesInterface)(nil).InsertTransaction), ctx, input)
}

// Rollback mocks base method.
func (m *MockTransactionRepositoriesInterface) Rollback() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback")
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionRepositoriesInterfaceMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransactionRepositoriesInterface)(nil).Rollback))
}

// UpdateStatusTransaction mocks base method.
func (m *MockTransactionRepositoriesInterface) UpdateStatusTransaction(ctx context.Context, transactionID uuid.UUID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusTransaction", ctx, transactionID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusTransaction indicates an expected call of UpdateStatusTransaction.
func (mr *MockTransactionRepositoriesInterfaceMockRecorder) UpdateStatusTransaction(ctx, transactionID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusTransaction", reflect.TypeOf((*MockTransactionRepositoriesInterface)(nil).UpdateStatusTransaction), ctx, transactionID, status)
}
