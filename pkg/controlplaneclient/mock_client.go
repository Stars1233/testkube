// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/controlplaneclient (interfaces: Client)

// Package controlplaneclient is a generated GoMock package.
package controlplaneclient

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	testkube "github.com/kubeshop/testkube/pkg/api/v1/testkube"
	cloud "github.com/kubeshop/testkube/pkg/cloud"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AppendExecutionReport mocks base method.
func (m *MockClient) AppendExecutionReport(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string, arg6 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendExecutionReport", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendExecutionReport indicates an expected call of AppendExecutionReport.
func (mr *MockClientMockRecorder) AppendExecutionReport(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendExecutionReport", reflect.TypeOf((*MockClient)(nil).AppendExecutionReport), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// CreateTestWorkflow mocks base method.
func (m *MockClient) CreateTestWorkflow(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTestWorkflow indicates an expected call of CreateTestWorkflow.
func (mr *MockClientMockRecorder) CreateTestWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestWorkflow", reflect.TypeOf((*MockClient)(nil).CreateTestWorkflow), arg0, arg1, arg2)
}

// CreateTestWorkflowTemplate mocks base method.
func (m *MockClient) CreateTestWorkflowTemplate(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestWorkflowTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTestWorkflowTemplate indicates an expected call of CreateTestWorkflowTemplate.
func (mr *MockClientMockRecorder) CreateTestWorkflowTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestWorkflowTemplate", reflect.TypeOf((*MockClient)(nil).CreateTestWorkflowTemplate), arg0, arg1, arg2)
}

// DeleteTestWorkflow mocks base method.
func (m *MockClient) DeleteTestWorkflow(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTestWorkflow indicates an expected call of DeleteTestWorkflow.
func (mr *MockClientMockRecorder) DeleteTestWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestWorkflow", reflect.TypeOf((*MockClient)(nil).DeleteTestWorkflow), arg0, arg1, arg2)
}

// DeleteTestWorkflowTemplate mocks base method.
func (m *MockClient) DeleteTestWorkflowTemplate(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestWorkflowTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTestWorkflowTemplate indicates an expected call of DeleteTestWorkflowTemplate.
func (mr *MockClientMockRecorder) DeleteTestWorkflowTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestWorkflowTemplate", reflect.TypeOf((*MockClient)(nil).DeleteTestWorkflowTemplate), arg0, arg1, arg2)
}

// DeleteTestWorkflowTemplatesByLabels mocks base method.
func (m *MockClient) DeleteTestWorkflowTemplatesByLabels(arg0 context.Context, arg1 string, arg2 map[string]string) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestWorkflowTemplatesByLabels", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTestWorkflowTemplatesByLabels indicates an expected call of DeleteTestWorkflowTemplatesByLabels.
func (mr *MockClientMockRecorder) DeleteTestWorkflowTemplatesByLabels(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestWorkflowTemplatesByLabels", reflect.TypeOf((*MockClient)(nil).DeleteTestWorkflowTemplatesByLabels), arg0, arg1, arg2)
}

// DeleteTestWorkflowsByLabels mocks base method.
func (m *MockClient) DeleteTestWorkflowsByLabels(arg0 context.Context, arg1 string, arg2 map[string]string) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTestWorkflowsByLabels", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTestWorkflowsByLabels indicates an expected call of DeleteTestWorkflowsByLabels.
func (mr *MockClientMockRecorder) DeleteTestWorkflowsByLabels(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTestWorkflowsByLabels", reflect.TypeOf((*MockClient)(nil).DeleteTestWorkflowsByLabels), arg0, arg1, arg2)
}

// FinishExecutionResult mocks base method.
func (m *MockClient) FinishExecutionResult(arg0 context.Context, arg1, arg2 string, arg3 *testkube.TestWorkflowResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishExecutionResult", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishExecutionResult indicates an expected call of FinishExecutionResult.
func (mr *MockClientMockRecorder) FinishExecutionResult(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishExecutionResult", reflect.TypeOf((*MockClient)(nil).FinishExecutionResult), arg0, arg1, arg2, arg3)
}

// GetCredential mocks base method.
func (m *MockClient) GetCredential(arg0 context.Context, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredential", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredential indicates an expected call of GetCredential.
func (mr *MockClientMockRecorder) GetCredential(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredential", reflect.TypeOf((*MockClient)(nil).GetCredential), arg0, arg1, arg2, arg3)
}

// GetExecution mocks base method.
func (m *MockClient) GetExecution(arg0 context.Context, arg1, arg2 string) (*testkube.TestWorkflowExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecution", arg0, arg1, arg2)
	ret0, _ := ret[0].(*testkube.TestWorkflowExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecution indicates an expected call of GetExecution.
func (mr *MockClientMockRecorder) GetExecution(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecution", reflect.TypeOf((*MockClient)(nil).GetExecution), arg0, arg1, arg2)
}

// GetGitHubToken mocks base method.
func (m *MockClient) GetGitHubToken(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGitHubToken indicates an expected call of GetGitHubToken.
func (mr *MockClientMockRecorder) GetGitHubToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubToken", reflect.TypeOf((*MockClient)(nil).GetGitHubToken), arg0, arg1)
}

// GetRunnerOngoingExecutions mocks base method.
func (m *MockClient) GetRunnerOngoingExecutions(arg0 context.Context) ([]*cloud.UnfinishedExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunnerOngoingExecutions", arg0)
	ret0, _ := ret[0].([]*cloud.UnfinishedExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunnerOngoingExecutions indicates an expected call of GetRunnerOngoingExecutions.
func (mr *MockClientMockRecorder) GetRunnerOngoingExecutions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunnerOngoingExecutions", reflect.TypeOf((*MockClient)(nil).GetRunnerOngoingExecutions), arg0)
}

// GetTestWorkflow mocks base method.
func (m *MockClient) GetTestWorkflow(arg0 context.Context, arg1, arg2 string) (*testkube.TestWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(*testkube.TestWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestWorkflow indicates an expected call of GetTestWorkflow.
func (mr *MockClientMockRecorder) GetTestWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestWorkflow", reflect.TypeOf((*MockClient)(nil).GetTestWorkflow), arg0, arg1, arg2)
}

// GetTestWorkflowTemplate mocks base method.
func (m *MockClient) GetTestWorkflowTemplate(arg0 context.Context, arg1, arg2 string) (*testkube.TestWorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestWorkflowTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*testkube.TestWorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestWorkflowTemplate indicates an expected call of GetTestWorkflowTemplate.
func (mr *MockClientMockRecorder) GetTestWorkflowTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestWorkflowTemplate", reflect.TypeOf((*MockClient)(nil).GetTestWorkflowTemplate), arg0, arg1, arg2)
}

// InitExecution mocks base method.
func (m *MockClient) InitExecution(arg0 context.Context, arg1, arg2 string, arg3 []testkube.TestWorkflowSignature, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitExecution", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitExecution indicates an expected call of InitExecution.
func (mr *MockClientMockRecorder) InitExecution(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitExecution", reflect.TypeOf((*MockClient)(nil).InitExecution), arg0, arg1, arg2, arg3, arg4)
}

// IsLegacy mocks base method.
func (m *MockClient) IsLegacy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLegacy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLegacy indicates an expected call of IsLegacy.
func (mr *MockClientMockRecorder) IsLegacy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLegacy", reflect.TypeOf((*MockClient)(nil).IsLegacy))
}

// IsRunner mocks base method.
func (m *MockClient) IsRunner() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunner")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunner indicates an expected call of IsRunner.
func (mr *MockClientMockRecorder) IsRunner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunner", reflect.TypeOf((*MockClient)(nil).IsRunner))
}

// IsSuperAgent mocks base method.
func (m *MockClient) IsSuperAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSuperAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSuperAgent indicates an expected call of IsSuperAgent.
func (mr *MockClientMockRecorder) IsSuperAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSuperAgent", reflect.TypeOf((*MockClient)(nil).IsSuperAgent))
}

// ListTestWorkflowLabels mocks base method.
func (m *MockClient) ListTestWorkflowLabels(arg0 context.Context, arg1 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestWorkflowLabels", arg0, arg1)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTestWorkflowLabels indicates an expected call of ListTestWorkflowLabels.
func (mr *MockClientMockRecorder) ListTestWorkflowLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestWorkflowLabels", reflect.TypeOf((*MockClient)(nil).ListTestWorkflowLabels), arg0, arg1)
}

// ListTestWorkflowTemplateLabels mocks base method.
func (m *MockClient) ListTestWorkflowTemplateLabels(arg0 context.Context, arg1 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestWorkflowTemplateLabels", arg0, arg1)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTestWorkflowTemplateLabels indicates an expected call of ListTestWorkflowTemplateLabels.
func (mr *MockClientMockRecorder) ListTestWorkflowTemplateLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestWorkflowTemplateLabels", reflect.TypeOf((*MockClient)(nil).ListTestWorkflowTemplateLabels), arg0, arg1)
}

// ListTestWorkflowTemplates mocks base method.
func (m *MockClient) ListTestWorkflowTemplates(arg0 context.Context, arg1 string, arg2 ListTestWorkflowTemplateOptions) TestWorkflowTemplatesReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestWorkflowTemplates", arg0, arg1, arg2)
	ret0, _ := ret[0].(TestWorkflowTemplatesReader)
	return ret0
}

// ListTestWorkflowTemplates indicates an expected call of ListTestWorkflowTemplates.
func (mr *MockClientMockRecorder) ListTestWorkflowTemplates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestWorkflowTemplates", reflect.TypeOf((*MockClient)(nil).ListTestWorkflowTemplates), arg0, arg1, arg2)
}

// ListTestWorkflows mocks base method.
func (m *MockClient) ListTestWorkflows(arg0 context.Context, arg1 string, arg2 ListTestWorkflowOptions) TestWorkflowsReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestWorkflows", arg0, arg1, arg2)
	ret0, _ := ret[0].(TestWorkflowsReader)
	return ret0
}

// ListTestWorkflows indicates an expected call of ListTestWorkflows.
func (mr *MockClientMockRecorder) ListTestWorkflows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestWorkflows", reflect.TypeOf((*MockClient)(nil).ListTestWorkflows), arg0, arg1, arg2)
}

// ProcessExecutionNotificationRequests mocks base method.
func (m *MockClient) ProcessExecutionNotificationRequests(arg0 context.Context, arg1 func(context.Context, *cloud.TestWorkflowNotificationsRequest) NotificationWatcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessExecutionNotificationRequests", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessExecutionNotificationRequests indicates an expected call of ProcessExecutionNotificationRequests.
func (mr *MockClientMockRecorder) ProcessExecutionNotificationRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessExecutionNotificationRequests", reflect.TypeOf((*MockClient)(nil).ProcessExecutionNotificationRequests), arg0, arg1)
}

// ProcessExecutionParallelWorkerNotificationRequests mocks base method.
func (m *MockClient) ProcessExecutionParallelWorkerNotificationRequests(arg0 context.Context, arg1 func(context.Context, *cloud.TestWorkflowParallelStepNotificationsRequest) NotificationWatcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessExecutionParallelWorkerNotificationRequests", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessExecutionParallelWorkerNotificationRequests indicates an expected call of ProcessExecutionParallelWorkerNotificationRequests.
func (mr *MockClientMockRecorder) ProcessExecutionParallelWorkerNotificationRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessExecutionParallelWorkerNotificationRequests", reflect.TypeOf((*MockClient)(nil).ProcessExecutionParallelWorkerNotificationRequests), arg0, arg1)
}

// ProcessExecutionServiceNotificationRequests mocks base method.
func (m *MockClient) ProcessExecutionServiceNotificationRequests(arg0 context.Context, arg1 func(context.Context, *cloud.TestWorkflowServiceNotificationsRequest) NotificationWatcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessExecutionServiceNotificationRequests", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessExecutionServiceNotificationRequests indicates an expected call of ProcessExecutionServiceNotificationRequests.
func (mr *MockClientMockRecorder) ProcessExecutionServiceNotificationRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessExecutionServiceNotificationRequests", reflect.TypeOf((*MockClient)(nil).ProcessExecutionServiceNotificationRequests), arg0, arg1)
}

// SaveExecutionArtifactGetPresignedURL mocks base method.
func (m *MockClient) SaveExecutionArtifactGetPresignedURL(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveExecutionArtifactGetPresignedURL", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveExecutionArtifactGetPresignedURL indicates an expected call of SaveExecutionArtifactGetPresignedURL.
func (mr *MockClientMockRecorder) SaveExecutionArtifactGetPresignedURL(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveExecutionArtifactGetPresignedURL", reflect.TypeOf((*MockClient)(nil).SaveExecutionArtifactGetPresignedURL), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SaveExecutionLogs mocks base method.
func (m *MockClient) SaveExecutionLogs(arg0 context.Context, arg1, arg2, arg3 string, arg4 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveExecutionLogs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveExecutionLogs indicates an expected call of SaveExecutionLogs.
func (mr *MockClientMockRecorder) SaveExecutionLogs(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveExecutionLogs", reflect.TypeOf((*MockClient)(nil).SaveExecutionLogs), arg0, arg1, arg2, arg3, arg4)
}

// SaveExecutionLogsGetPresignedURL mocks base method.
func (m *MockClient) SaveExecutionLogsGetPresignedURL(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveExecutionLogsGetPresignedURL", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveExecutionLogsGetPresignedURL indicates an expected call of SaveExecutionLogsGetPresignedURL.
func (mr *MockClientMockRecorder) SaveExecutionLogsGetPresignedURL(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveExecutionLogsGetPresignedURL", reflect.TypeOf((*MockClient)(nil).SaveExecutionLogsGetPresignedURL), arg0, arg1, arg2, arg3)
}

// ScheduleExecution mocks base method.
func (m *MockClient) ScheduleExecution(arg0 context.Context, arg1 string, arg2 *cloud.ScheduleRequest) ExecutionsReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleExecution", arg0, arg1, arg2)
	ret0, _ := ret[0].(ExecutionsReader)
	return ret0
}

// ScheduleExecution indicates an expected call of ScheduleExecution.
func (mr *MockClientMockRecorder) ScheduleExecution(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleExecution", reflect.TypeOf((*MockClient)(nil).ScheduleExecution), arg0, arg1, arg2)
}

// UpdateExecutionOutput mocks base method.
func (m *MockClient) UpdateExecutionOutput(arg0 context.Context, arg1, arg2 string, arg3 []testkube.TestWorkflowOutput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExecutionOutput", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExecutionOutput indicates an expected call of UpdateExecutionOutput.
func (mr *MockClientMockRecorder) UpdateExecutionOutput(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExecutionOutput", reflect.TypeOf((*MockClient)(nil).UpdateExecutionOutput), arg0, arg1, arg2, arg3)
}

// UpdateExecutionResult mocks base method.
func (m *MockClient) UpdateExecutionResult(arg0 context.Context, arg1, arg2 string, arg3 *testkube.TestWorkflowResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExecutionResult", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExecutionResult indicates an expected call of UpdateExecutionResult.
func (mr *MockClientMockRecorder) UpdateExecutionResult(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExecutionResult", reflect.TypeOf((*MockClient)(nil).UpdateExecutionResult), arg0, arg1, arg2, arg3)
}

// UpdateTestWorkflow mocks base method.
func (m *MockClient) UpdateTestWorkflow(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestWorkflow", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTestWorkflow indicates an expected call of UpdateTestWorkflow.
func (mr *MockClientMockRecorder) UpdateTestWorkflow(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestWorkflow", reflect.TypeOf((*MockClient)(nil).UpdateTestWorkflow), arg0, arg1, arg2)
}

// UpdateTestWorkflowTemplate mocks base method.
func (m *MockClient) UpdateTestWorkflowTemplate(arg0 context.Context, arg1 string, arg2 testkube.TestWorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestWorkflowTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTestWorkflowTemplate indicates an expected call of UpdateTestWorkflowTemplate.
func (mr *MockClientMockRecorder) UpdateTestWorkflowTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestWorkflowTemplate", reflect.TypeOf((*MockClient)(nil).UpdateTestWorkflowTemplate), arg0, arg1, arg2)
}

// WatchRunnerRequests mocks base method.
func (m *MockClient) WatchRunnerRequests(arg0 context.Context) RunnerRequestsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRunnerRequests", arg0)
	ret0, _ := ret[0].(RunnerRequestsWatcher)
	return ret0
}

// WatchRunnerRequests indicates an expected call of WatchRunnerRequests.
func (mr *MockClientMockRecorder) WatchRunnerRequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRunnerRequests", reflect.TypeOf((*MockClient)(nil).WatchRunnerRequests), arg0)
}

// WatchTestWorkflowTemplateUpdates mocks base method.
func (m *MockClient) WatchTestWorkflowTemplateUpdates(arg0 context.Context, arg1 string, arg2 bool) TestWorkflowTemplateWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchTestWorkflowTemplateUpdates", arg0, arg1, arg2)
	ret0, _ := ret[0].(TestWorkflowTemplateWatcher)
	return ret0
}

// WatchTestWorkflowTemplateUpdates indicates an expected call of WatchTestWorkflowTemplateUpdates.
func (mr *MockClientMockRecorder) WatchTestWorkflowTemplateUpdates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchTestWorkflowTemplateUpdates", reflect.TypeOf((*MockClient)(nil).WatchTestWorkflowTemplateUpdates), arg0, arg1, arg2)
}

// WatchTestWorkflowUpdates mocks base method.
func (m *MockClient) WatchTestWorkflowUpdates(arg0 context.Context, arg1 string, arg2 bool) TestWorkflowWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchTestWorkflowUpdates", arg0, arg1, arg2)
	ret0, _ := ret[0].(TestWorkflowWatcher)
	return ret0
}

// WatchTestWorkflowUpdates indicates an expected call of WatchTestWorkflowUpdates.
func (mr *MockClientMockRecorder) WatchTestWorkflowUpdates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchTestWorkflowUpdates", reflect.TypeOf((*MockClient)(nil).WatchTestWorkflowUpdates), arg0, arg1, arg2)
}
