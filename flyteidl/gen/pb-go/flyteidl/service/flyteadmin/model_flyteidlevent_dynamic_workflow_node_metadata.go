/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// For dynamic workflow nodes we send information about the dynamic workflow definition that gets generated.
type FlyteidleventDynamicWorkflowNodeMetadata struct {
	// id represents the unique identifier of the workflow.
	Id *CoreIdentifier `json:"id,omitempty"`
	// Represents the compiled representation of the embedded dynamic workflow.
	CompiledWorkflow *CoreCompiledWorkflowClosure `json:"compiled_workflow,omitempty"`
	// dynamic_job_spec_uri is the location of the DynamicJobSpec proto message for this DynamicWorkflow. This is required to correctly recover partially completed executions where the workflow has already been compiled.
	DynamicJobSpecUri string `json:"dynamic_job_spec_uri,omitempty"`
}
