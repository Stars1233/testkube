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

type Toleration struct {
	Key               string        `json:"key,omitempty"`
	Operator          string        `json:"operator,omitempty"`
	Value             string        `json:"value,omitempty"`
	Effect            string        `json:"effect,omitempty"`
	TolerationSeconds *BoxedInteger `json:"tolerationSeconds,omitempty"`
}
