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

// Test with latest Execution result
type TestWithExecution struct {
	Test            *Test      `json:"test"`
	LatestExecution *Execution `json:"latestExecution,omitempty"`
}
