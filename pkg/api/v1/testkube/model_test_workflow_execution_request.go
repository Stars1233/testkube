/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: contact@testkube.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

type TestWorkflowExecutionRequest struct {
	// custom execution name
	Name   string            `json:"name,omitempty"`
	Config map[string]string `json:"config,omitempty"`
	// test workflow execution name started the test workflow execution
	TestWorkflowExecutionName string `json:"testWorkflowExecutionName,omitempty"`
	// whether webhooks on the execution of this test workflow are disabled
	DisableWebhooks bool                        `json:"disableWebhooks,omitempty"`
	Tags            map[string]string           `json:"tags,omitempty"`
	Target          *ExecutionTarget            `json:"target,omitempty"`
	RunningContext  *TestWorkflowRunningContext `json:"runningContext,omitempty"`
	// parent execution ids
	ParentExecutionIds []string `json:"parentExecutionIds,omitempty"`
}
