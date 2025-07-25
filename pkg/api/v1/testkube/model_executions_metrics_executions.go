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

import (
	"time"
)

type ExecutionsMetricsExecutions struct {
	ExecutionId string    `json:"executionId,omitempty"`
	GroupId     string    `json:"groupId,omitempty"`
	Duration    string    `json:"duration,omitempty"`
	DurationMs  int32     `json:"durationMs,omitempty"`
	Status      string    `json:"status,omitempty"`
	Name        string    `json:"name,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	RunnerId    string    `json:"runnerId,omitempty"`
}
