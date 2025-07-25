// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/newclients/testworkflowclient (interfaces: TestWorkflowClient)

// Package testworkflowclient is a generated GoMock package.
package testworkflowclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	testkube "github.com/kubeshop/testkube/pkg/api/v1/testkube"
	types "k8s.io/apimachinery/pkg/types"
)

// MockTestWorkflowClient is a mock of TestWorkflowClient interface.
type MockTestWorkflowClient struct {
	ctrl     *gomock.Controller
	recorder *MockTestWorkflowClientMockRecorder
}

// MockTestWorkflowClientMockRecorder is the mock recorder for MockTestWorkflowClient.
type MockTestWorkflowClientMockRecorder struct {
	mock *MockTestWorkflowClient
}

// NewMockTestWorkflowClient creates a new mock instance.
func NewMockTestWorkflowClient(ctrl *gomock.Controller) *MockTestWorkflowClient {
	mock := &MockTestWorkflowClient{ctrl: ctrl}
	mock.recorder = &MockTestWorkflowClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestWorkflowClient) EXPECT() *MockTestWorkflowClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTestWorkflowClient) Create(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTestWorkflowClientMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTestWorkflowClient)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockTestWorkflowClient) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTestWorkflowClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTestWorkflowClient)(nil).Delete), arg0, arg1, arg2)
}

// DeleteByLabels mocks base method.
func (m *MockTestWorkflowClient) DeleteByLabels(arg0 context.Context, arg1 string, arg2 map[string]string) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByLabels", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByLabels indicates an expected call of DeleteByLabels.
func (mr *MockTestWorkflowClientMockRecorder) DeleteByLabels(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByLabels", reflect.TypeOf((*MockTestWorkflowClient)(nil).DeleteByLabels), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockTestWorkflowClient) Get(arg0 context.Context, arg1, arg2 string) (*testkube.TestWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*testkube.TestWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTestWorkflowClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTestWorkflowClient)(nil).Get), arg0, arg1, arg2)
}

// GetKubernetesObjectUID mocks base method.
func (m *MockTestWorkflowClient) GetKubernetesObjectUID(arg0 context.Context, arg1, arg2 string) (types.UID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubernetesObjectUID", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.UID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubernetesObjectUID indicates an expected call of GetKubernetesObjectUID.
func (mr *MockTestWorkflowClientMockRecorder) GetKubernetesObjectUID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesObjectUID", reflect.TypeOf((*MockTestWorkflowClient)(nil).GetKubernetesObjectUID), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockTestWorkflowClient) List(arg0 context.Context, arg1 string, arg2 ListOptions) ([]testkube.TestWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]testkube.TestWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTestWorkflowClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTestWorkflowClient)(nil).List), arg0, arg1, arg2)
}

// ListLabels mocks base method.
func (m *MockTestWorkflowClient) ListLabels(arg0 context.Context, arg1 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLabels", arg0, arg1)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLabels indicates an expected call of ListLabels.
func (mr *MockTestWorkflowClientMockRecorder) ListLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabels", reflect.TypeOf((*MockTestWorkflowClient)(nil).ListLabels), arg0, arg1)
}

// Update mocks base method.
func (m *MockTestWorkflowClient) Update(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTestWorkflowClientMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTestWorkflowClient)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockTestWorkflowClient) UpdateStatus(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockTestWorkflowClientMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockTestWorkflowClient)(nil).UpdateStatus), arg0, arg1, arg2)
}

// WatchUpdates mocks base method.
func (m *MockTestWorkflowClient) WatchUpdates(arg0 context.Context, arg1 string, arg2 bool) Watcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpdates", arg0, arg1, arg2)
	ret0, _ := ret[0].(Watcher)
	return ret0
}

// WatchUpdates indicates an expected call of WatchUpdates.
func (mr *MockTestWorkflowClientMockRecorder) WatchUpdates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpdates", reflect.TypeOf((*MockTestWorkflowClient)(nil).WatchUpdates), arg0, arg1, arg2)
}
