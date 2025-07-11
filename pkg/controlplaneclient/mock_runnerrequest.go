// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/controlplaneclient (interfaces: RunnerRequest)

// Package controlplaneclient is a generated GoMock package.
package controlplaneclient

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cloud "github.com/kubeshop/testkube/pkg/cloud"
)

// MockRunnerRequest is a mock of RunnerRequest interface.
type MockRunnerRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerRequestMockRecorder
}

// MockRunnerRequestMockRecorder is the mock recorder for MockRunnerRequest.
type MockRunnerRequestMockRecorder struct {
	mock *MockRunnerRequest
}

// NewMockRunnerRequest creates a new mock instance.
func NewMockRunnerRequest(ctrl *gomock.Controller) *MockRunnerRequest {
	mock := &MockRunnerRequest{ctrl: ctrl}
	mock.recorder = &MockRunnerRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnerRequest) EXPECT() *MockRunnerRequestMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockRunnerRequest) Abort() RunnerRequestOK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abort")
	ret0, _ := ret[0].(RunnerRequestOK)
	return ret0
}

// Abort indicates an expected call of Abort.
func (mr *MockRunnerRequestMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockRunnerRequest)(nil).Abort))
}

// Cancel mocks base method.
func (m *MockRunnerRequest) Cancel() RunnerRequestOK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(RunnerRequestOK)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockRunnerRequestMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockRunnerRequest)(nil).Cancel))
}

// Consider mocks base method.
func (m *MockRunnerRequest) Consider() RunnerRequestConsider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consider")
	ret0, _ := ret[0].(RunnerRequestConsider)
	return ret0
}

// Consider indicates an expected call of Consider.
func (mr *MockRunnerRequestMockRecorder) Consider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consider", reflect.TypeOf((*MockRunnerRequest)(nil).Consider))
}

// EnvironmentID mocks base method.
func (m *MockRunnerRequest) EnvironmentID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentID")
	ret0, _ := ret[0].(string)
	return ret0
}

// EnvironmentID indicates an expected call of EnvironmentID.
func (mr *MockRunnerRequestMockRecorder) EnvironmentID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentID", reflect.TypeOf((*MockRunnerRequest)(nil).EnvironmentID))
}

// ExecutionID mocks base method.
func (m *MockRunnerRequest) ExecutionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecutionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExecutionID indicates an expected call of ExecutionID.
func (mr *MockRunnerRequestMockRecorder) ExecutionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutionID", reflect.TypeOf((*MockRunnerRequest)(nil).ExecutionID))
}

// MessageID mocks base method.
func (m *MockRunnerRequest) MessageID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageID")
	ret0, _ := ret[0].(string)
	return ret0
}

// MessageID indicates an expected call of MessageID.
func (mr *MockRunnerRequestMockRecorder) MessageID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageID", reflect.TypeOf((*MockRunnerRequest)(nil).MessageID))
}

// Pause mocks base method.
func (m *MockRunnerRequest) Pause() RunnerRequestOK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause")
	ret0, _ := ret[0].(RunnerRequestOK)
	return ret0
}

// Pause indicates an expected call of Pause.
func (mr *MockRunnerRequestMockRecorder) Pause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockRunnerRequest)(nil).Pause))
}

// Ping mocks base method.
func (m *MockRunnerRequest) Ping() RunnerRequestOK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(RunnerRequestOK)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRunnerRequestMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRunnerRequest)(nil).Ping))
}

// Resume mocks base method.
func (m *MockRunnerRequest) Resume() RunnerRequestOK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resume")
	ret0, _ := ret[0].(RunnerRequestOK)
	return ret0
}

// Resume indicates an expected call of Resume.
func (mr *MockRunnerRequestMockRecorder) Resume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockRunnerRequest)(nil).Resume))
}

// SendError mocks base method.
func (m *MockRunnerRequest) SendError(arg0 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendError", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendError indicates an expected call of SendError.
func (mr *MockRunnerRequestMockRecorder) SendError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendError", reflect.TypeOf((*MockRunnerRequest)(nil).SendError), arg0)
}

// Start mocks base method.
func (m *MockRunnerRequest) Start() RunnerRequestStart {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(RunnerRequestStart)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockRunnerRequestMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRunnerRequest)(nil).Start))
}

// Type mocks base method.
func (m *MockRunnerRequest) Type() cloud.RunnerRequestType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(cloud.RunnerRequestType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockRunnerRequestMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockRunnerRequest)(nil).Type))
}
