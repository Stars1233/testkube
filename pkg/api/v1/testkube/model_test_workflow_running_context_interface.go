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

// running context interface for test workflow execution
type TestWorkflowRunningContextInterface struct {
	// interface name
	Name  string                                   `json:"name,omitempty"`
	Type_ *TestWorkflowRunningContextInterfaceType `json:"type"`
}
