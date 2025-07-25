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

// TestWorkflowParameterType : type of the config parameter
type TestWorkflowParameterType string

// List of TestWorkflowParameterType
const (
	STRING__TestWorkflowParameterType TestWorkflowParameterType = "string"
	INTEGER_TestWorkflowParameterType TestWorkflowParameterType = "integer"
	NUMBER_TestWorkflowParameterType  TestWorkflowParameterType = "number"
	BOOLEAN_TestWorkflowParameterType TestWorkflowParameterType = "boolean"
)
