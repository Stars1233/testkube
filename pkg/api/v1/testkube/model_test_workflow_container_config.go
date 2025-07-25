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

type TestWorkflowContainerConfig struct {
	WorkingDir *BoxedString `json:"workingDir,omitempty"`
	// image to be used for the container
	Image           string           `json:"image,omitempty"`
	ImagePullPolicy *ImagePullPolicy `json:"imagePullPolicy,omitempty"`
	// environment variables to append to the container
	Env []EnvVar `json:"env,omitempty"`
	// external environment variables to append to the container
	EnvFrom         []EnvFromSource        `json:"envFrom,omitempty"`
	Command         *BoxedStringList       `json:"command,omitempty"`
	Args            *BoxedStringList       `json:"args,omitempty"`
	Resources       *TestWorkflowResources `json:"resources,omitempty"`
	SecurityContext *SecurityContext       `json:"securityContext,omitempty"`
	// volumes to mount to the container
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
}
