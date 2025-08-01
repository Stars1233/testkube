package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSQLCTestWorkflowExecutionQueries_GetTestWorkflowExecution(t *testing.T) {
	// Create mock database connection
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	// Define expected query pattern
	expectedQuery := `SELECT
	    e\.id, e\.group_id, e\.runner_id, e\.runner_target, e\.runner_original_target, e\.name, e\.namespace, e\.number, e\.scheduled_at, e\.assigned_at, e\.status_at, e\.test_workflow_execution_name, e\.disable_webhooks, e\.tags, e\.running_context, e\.config_params, e\.created_at, e\.updated_at,
	    r\.status, r\.predicted_status, r\.queued_at, r\.started_at, r\.finished_at,
	    r\.duration, r\.total_duration, r\.duration_ms, r\.paused_ms, r\.total_duration_ms,
	    r\.pauses, r\.initialization, r\.steps,
	    w\.name as workflow_name, w\.namespace as workflow_namespace, w\.description as workflow_description,
	    w\.labels as workflow_labels, w\.annotations as workflow_annotations, w\.created as workflow_created,
	    w\.updated as workflow_updated, w\.spec as workflow_spec, w\.read_only as workflow_read_only,
	    w\.status as workflow_status,
	    rw\.name as resolved_workflow_name, rw\.namespace as resolved_workflow_namespace,
	    rw\.description as resolved_workflow_description, rw\.labels as resolved_workflow_labels,
	    rw\.annotations as resolved_workflow_annotations, rw\.created as resolved_workflow_created,
	    rw\.updated as resolved_workflow_updated, rw\.spec as resolved_workflow_spec,
	    rw\.read_only as resolved_workflow_read_only, rw\.status as resolved_workflow_status,
	    COALESCE\(
	        \(SELECT json_agg\(
	            json_build_object\(
	                'id', s\.id,
	                'ref', s\.ref,
	                'name', s\.name,
	                'category', s\.category,
	                'optional', s\.optional,
	                'negative', s\.negative,
	                'parent_id', s\.parent_id
	            \) ORDER BY s\.id
	        \) FROM test_workflow_signatures s WHERE s\.execution_id = e\.id\),
	        '\[\]'::json
	    \)::json as signatures_json,
	    COALESCE\(
	        \(SELECT json_agg\(
	            json_build_object\(
	                'id', o\.id,
	                'ref', o\.ref,
	                'name', o\.name,
	                'value', o\.value
	            \) ORDER BY o\.id
	        \) FROM test_workflow_outputs o WHERE o\.execution_id = e\.id\),
	        '\[\]'::json
	    \)::json as outputs_json,
	    COALESCE\(
	        \(SELECT json_agg\(
	            json_build_object\(
	                'id', rep\.id,
	                'ref', rep\.ref,
	                'kind', rep\.kind,
	                'file', rep\.file,
	                'summary', rep\.summary
	            \) ORDER BY rep\.id
	        \) FROM test_workflow_reports rep WHERE rep\.execution_id = e\.id\),
	        '\[\]'::json
	    \)::json as reports_json,
	    ra\.global as resource_aggregations_global,
	    ra\.step as resource_aggregations_step
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
LEFT JOIN test_workflows rw ON e\.id = rw\.execution_id AND rw\.workflow_type = 'resolved_workflow'
LEFT JOIN test_workflow_resource_aggregations ra ON e\.id = ra\.execution_id
WHERE e\.id = \$1 OR e\.name = \$1`

	// Mock expected result
	rows := mock.NewRows([]string{
		"id", "group_id", "runner_id", "runner_target", "runner_original_target", "name", "namespace", "number",
		"scheduled_at", "assigned_at", "status_at", "test_workflow_execution_name", "disable_webhooks",
		"tags", "running_context", "config_params", "created_at", "updated_at",
		"status", "predicted_status", "queued_at", "started_at", "finished_at",
		"duration", "total_duration", "duration_ms", "paused_ms", "total_duration_ms",
		"pauses", "initialization", "steps",
		"workflow_name", "workflow_namespace", "workflow_description", "workflow_labels", "workflow_annotations",
		"workflow_created", "workflow_updated", "workflow_spec", "workflow_read_only", "workflow_status",
		"resolved_workflow_name", "resolved_workflow_namespace", "resolved_workflow_description",
		"resolved_workflow_labels", "resolved_workflow_annotations", "resolved_workflow_created",
		"resolved_workflow_updated", "resolved_workflow_spec", "resolved_workflow_read_only", "resolved_workflow_status",
		"signatures_json", "outputs_json", "reports_json", "resource_aggregations_global", "resource_aggregations_step",
	}).AddRow(
		"test-id", "group-1", "runner-1", []byte(`{}`), []byte(`{}`), "test-execution", "default", int64(1),
		time.Now(), time.Now(), time.Now(), "test-execution-name", false,
		[]byte(`{"env":"test"}`), []byte(`{}`), []byte(`{}`), time.Now(), time.Now(),
		"passed", "passed", time.Now(), time.Now(), time.Now(),
		"5m", "5m", int64(300000), int64(0), int64(300000),
		[]byte(`[]`), []byte(`{}`), []byte(`{}`),
		"test-workflow", "default", "Test workflow", []byte(`{}`), []byte(`{}`),
		time.Now(), time.Now(), []byte(`{}`), false, []byte(`{}`),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		[]byte(`[]`), []byte(`[]`), []byte(`[]`), []byte(`{}`), []byte(`{}`),
	)

	mock.ExpectQuery(expectedQuery).WithArgs("test-id").WillReturnRows(rows)

	// Execute query
	result, err := queries.GetTestWorkflowExecution(ctx, "test-id")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result.ID)
	assert.Equal(t, "test-execution", result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetTestWorkflowExecutionByNameAndTestWorkflow(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT 
    e\.id, e\.group_id, e\.runner_id, e\.runner_target, e\.runner_original_target, e\.name, e\.namespace, e\.number, e\.scheduled_at, e\.assigned_at, e\.status_at, e\.test_workflow_execution_name, e\.disable_webhooks, e\.tags, e\.running_context, e\.config_params, e\.created_at, e\.updated_at,
    r\.status, r\.predicted_status, r\.queued_at, r\.started_at, r\.finished_at,
    r\.duration, r\.total_duration, r\.duration_ms, r\.paused_ms, r\.total_duration_ms,
    r\.pauses, r\.initialization, r\.steps,
    w\.name as workflow_name, w\.namespace as workflow_namespace, w\.description as workflow_description,
    w\.labels as workflow_labels, w\.annotations as workflow_annotations, w\.created as workflow_created,
    w\.updated as workflow_updated, w\.spec as workflow_spec, w\.read_only as workflow_read_only,
    w\.status as workflow_status,
    rw\.name as resolved_workflow_name, rw\.namespace as resolved_workflow_namespace, 
    rw\.description as resolved_workflow_description, rw\.labels as resolved_workflow_labels,
    rw\.annotations as resolved_workflow_annotations, rw\.created as resolved_workflow_created,
    rw\.updated as resolved_workflow_updated, rw\.spec as resolved_workflow_spec,
    rw\.read_only as resolved_workflow_read_only, rw\.status as resolved_workflow_status,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', s\.id,
                'ref', s\.ref,
                'name', s\.name,
                'category', s\.category,
                'optional', s\.optional,
                'negative', s\.negative,
                'parent_id', s\.parent_id
            \) ORDER BY s\.id
        \) FROM test_workflow_signatures s WHERE s\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as signatures_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', o\.id,
                'ref', o\.ref,
                'name', o\.name,
                'value', o\.value
            \) ORDER BY o\.id
        \) FROM test_workflow_outputs o WHERE o\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as outputs_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', rep\.id,
                'ref', rep\.ref,
                'kind', rep\.kind,
                'file', rep\.file,
                'summary', rep\.summary
            \) ORDER BY rep\.id
        \) FROM test_workflow_reports rep WHERE rep\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as reports_json,
    ra\.global as resource_aggregations_global,
    ra\.step as resource_aggregations_step
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
LEFT JOIN test_workflows rw ON e\.id = rw\.execution_id AND rw\.workflow_type = 'resolved_workflow'
LEFT JOIN test_workflow_resource_aggregations ra ON e\.id = ra\.execution_id
WHERE \(e\.id = \$1 OR e\.name = \$1\) AND w\.name = \$2`

	rows := mock.NewRows([]string{
		"id", "group_id", "runner_id", "runner_target", "runner_original_target", "name", "namespace", "number",
		"scheduled_at", "assigned_at", "status_at", "test_workflow_execution_name", "disable_webhooks",
		"tags", "running_context", "config_params", "created_at", "updated_at",
		"status", "predicted_status", "queued_at", "started_at", "finished_at",
		"duration", "total_duration", "duration_ms", "paused_ms", "total_duration_ms",
		"pauses", "initialization", "steps",
		"workflow_name", "workflow_namespace", "workflow_description", "workflow_labels", "workflow_annotations",
		"workflow_created", "workflow_updated", "workflow_spec", "workflow_read_only", "workflow_status",
		"resolved_workflow_name", "resolved_workflow_namespace", "resolved_workflow_description",
		"resolved_workflow_labels", "resolved_workflow_annotations", "resolved_workflow_created",
		"resolved_workflow_updated", "resolved_workflow_spec", "resolved_workflow_read_only", "resolved_workflow_status",
		"signatures_json", "outputs_json", "reports_json", "resource_aggregations_global", "resource_aggregations_step",
	}).AddRow(
		"test-id", "group-1", "runner-1", []byte(`{}`), []byte(`{}`), "test-execution", "default", int64(1),
		time.Now(), time.Now(), time.Now(), "test-execution-name", false,
		[]byte(`{"env":"test"}`), []byte(`{}`), []byte(`{}`), time.Now(), time.Now(),
		"passed", "passed", time.Now(), time.Now(), time.Now(),
		"5m", "5m", int64(300000), int64(0), int64(300000),
		[]byte(`[]`), []byte(`{}`), []byte(`{}`),
		"test-workflow", "default", "Test workflow", []byte(`{}`), []byte(`{}`),
		time.Now(), time.Now(), []byte(`{}`), false, []byte(`{}`),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		[]byte(`[]`), []byte(`[]`), []byte(`[]`), []byte(`{}`), []byte(`{}`),
	)

	mock.ExpectQuery(expectedQuery).WithArgs("test-execution", "test-workflow").WillReturnRows(rows)

	// Execute query
	params := GetTestWorkflowExecutionByNameAndTestWorkflowParams{
		Name:         "test-execution",
		WorkflowName: "test-workflow",
	}
	result, err := queries.GetTestWorkflowExecutionByNameAndTestWorkflow(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result.ID)
	assert.Equal(t, "test-execution", result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetLatestTestWorkflowExecutionByTestWorkflow(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT 
    e\.id, e\.group_id, e\.runner_id, e\.runner_target, e\.runner_original_target, e\.name, e\.namespace, e\.number, e\.scheduled_at, e\.assigned_at, e\.status_at, e\.test_workflow_execution_name, e\.disable_webhooks, e\.tags, e\.running_context, e\.config_params, e\.created_at, e\.updated_at,
    r\.status, r\.predicted_status, r\.queued_at, r\.started_at, r\.finished_at,
    r\.duration, r\.total_duration, r\.duration_ms, r\.paused_ms, r\.total_duration_ms,
    r\.pauses, r\.initialization, r\.steps,
    w\.name as workflow_name, w\.namespace as workflow_namespace, w\.description as workflow_description,
    w\.labels as workflow_labels, w\.annotations as workflow_annotations, w\.created as workflow_created,
    w\.updated as workflow_updated, w\.spec as workflow_spec, w\.read_only as workflow_read_only,
    w\.status as workflow_status,
    rw\.name as resolved_workflow_name, rw\.namespace as resolved_workflow_namespace, 
    rw\.description as resolved_workflow_description, rw\.labels as resolved_workflow_labels,
    rw\.annotations as resolved_workflow_annotations, rw\.created as resolved_workflow_created,
    rw\.updated as resolved_workflow_updated, rw\.spec as resolved_workflow_spec,
    rw\.read_only as resolved_workflow_read_only, rw\.status as resolved_workflow_status,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', s\.id,
                'ref', s\.ref,
                'name', s\.name,
                'category', s\.category,
                'optional', s\.optional,
                'negative', s\.negative,
                'parent_id', s\.parent_id
            \) ORDER BY s\.id
        \) FROM test_workflow_signatures s WHERE s\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as signatures_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', o\.id,
                'ref', o\.ref,
                'name', o\.name,
                'value', o\.value
            \) ORDER BY o\.id
        \) FROM test_workflow_outputs o WHERE o\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as outputs_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', rep\.id,
                'ref', rep\.ref,
                'kind', rep\.kind,
                'file', rep\.file,
                'summary', rep\.summary
            \) ORDER BY rep\.id
        \) FROM test_workflow_reports rep WHERE rep\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as reports_json,
    ra\.global as resource_aggregations_global,
    ra\.step as resource_aggregations_step
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
LEFT JOIN test_workflows rw ON e\.id = rw\.execution_id AND rw\.workflow_type = 'resolved_workflow'
LEFT JOIN test_workflow_resource_aggregations ra ON e\.id = ra\.execution_id
WHERE w\.name = \$1::text
ORDER BY
    CASE
        WHEN \$2::boolean = true THEN e\.number
        WHEN \$2::boolean = false THEN EXTRACT\(EPOCH FROM e\.status_at\)::integer
    END DESC
LIMIT 1
`

	rows := mock.NewRows([]string{
		"id", "group_id", "runner_id", "runner_target", "runner_original_target", "name", "namespace", "number",
		"scheduled_at", "assigned_at", "status_at", "test_workflow_execution_name", "disable_webhooks",
		"tags", "running_context", "config_params", "created_at", "updated_at",
		"status", "predicted_status", "queued_at", "started_at", "finished_at",
		"duration", "total_duration", "duration_ms", "paused_ms", "total_duration_ms",
		"pauses", "initialization", "steps",
		"workflow_name", "workflow_namespace", "workflow_description", "workflow_labels", "workflow_annotations",
		"workflow_created", "workflow_updated", "workflow_spec", "workflow_read_only", "workflow_status",
		"resolved_workflow_name", "resolved_workflow_namespace", "resolved_workflow_description",
		"resolved_workflow_labels", "resolved_workflow_annotations", "resolved_workflow_created",
		"resolved_workflow_updated", "resolved_workflow_spec", "resolved_workflow_read_only", "resolved_workflow_status",
		"signatures_json", "outputs_json", "reports_json", "resource_aggregations_global", "resource_aggregations_step",
	}).AddRow(
		"test-id", "group-1", "runner-1", []byte(`{}`), []byte(`{}`), "test-execution", "default", int64(1),
		time.Now(), time.Now(), time.Now(), "test-execution-name", false,
		[]byte(`{"env":"test"}`), []byte(`{}`), []byte(`{}`), time.Now(), time.Now(),
		"passed", "passed", time.Now(), time.Now(), time.Now(),
		"5m", "5m", int64(300000), int64(0), int64(300000),
		[]byte(`[]`), []byte(`{}`), []byte(`{}`),
		"test-workflow", "default", "Test workflow", []byte(`{}`), []byte(`{}`),
		time.Now(), time.Now(), []byte(`{}`), false, []byte(`{}`),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		[]byte(`[]`), []byte(`[]`), []byte(`[]`), []byte(`{}`), []byte(`{}`),
	)

	mock.ExpectQuery(expectedQuery).WithArgs("test-workflow", true).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetLatestTestWorkflowExecutionByTestWorkflow(ctx, GetLatestTestWorkflowExecutionByTestWorkflowParams{
		WorkflowName: "test-workflow",
		SortByNumber: true,
	})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result.ID)
	assert.Equal(t, "test-execution", result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetLatestTestWorkflowExecutionsByTestWorkflows(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT DISTINCT ON \(w\.name\)
    e\.id, e\.group_id, e\.runner_id, e\.runner_target, e\.runner_original_target, e\.name, e\.namespace, e\.number, e\.scheduled_at, e\.assigned_at, e\.status_at, e\.test_workflow_execution_name, e\.disable_webhooks, e\.tags, e\.running_context, e\.config_params, e\.created_at, e\.updated_at,
    r\.status, r\.predicted_status, r\.queued_at, r\.started_at, r\.finished_at,
    r\.duration, r\.total_duration, r\.duration_ms, r\.paused_ms, r\.total_duration_ms,
    r\.pauses, r\.initialization, r\.steps,
    w\.name as workflow_name, w\.namespace as workflow_namespace, w\.description as workflow_description,
    w\.labels as workflow_labels, w\.annotations as workflow_annotations, w\.created as workflow_created,
    w\.updated as workflow_updated, w\.spec as workflow_spec, w\.read_only as workflow_read_only,
    w\.status as workflow_status,
    rw\.name as resolved_workflow_name, rw\.namespace as resolved_workflow_namespace, 
    rw\.description as resolved_workflow_description, rw\.labels as resolved_workflow_labels,
    rw\.annotations as resolved_workflow_annotations, rw\.created as resolved_workflow_created,
    rw\.updated as resolved_workflow_updated, rw\.spec as resolved_workflow_spec,
    rw\.read_only as resolved_workflow_read_only, rw\.status as resolved_workflow_status,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', s\.id,
                'ref', s\.ref,
                'name', s\.name,
                'category', s\.category,
                'optional', s\.optional,
                'negative', s\.negative,
                'parent_id', s\.parent_id
            \) ORDER BY s\.id
        \) FROM test_workflow_signatures s WHERE s\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as signatures_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', o\.id,
                'ref', o\.ref,
                'name', o\.name,
                'value', o\.value
            \) ORDER BY o\.id
        \) FROM test_workflow_outputs o WHERE o\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as outputs_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', rep\.id,
                'ref', rep\.ref,
                'kind', rep\.kind,
                'file', rep\.file,
                'summary', rep\.summary
            \) ORDER BY rep\.id
        \) FROM test_workflow_reports rep WHERE rep\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as reports_json,
    ra\.global as resource_aggregations_global,
    ra\.step as resource_aggregations_step
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
LEFT JOIN test_workflows rw ON e\.id = rw\.execution_id AND rw\.workflow_type = 'resolved_workflow'
LEFT JOIN test_workflow_resource_aggregations ra ON e\.id = ra\.execution_id
WHERE w\.name = ANY\(\$1::text\[\]\)
ORDER BY w\.name, e\.status_at DESC`

	rows := mock.NewRows([]string{
		"id", "group_id", "runner_id", "runner_target", "runner_original_target", "name", "namespace", "number",
		"scheduled_at", "assigned_at", "status_at", "test_workflow_execution_name", "disable_webhooks",
		"tags", "running_context", "config_params", "created_at", "updated_at",
		"status", "predicted_status", "queued_at", "started_at", "finished_at",
		"duration", "total_duration", "duration_ms", "paused_ms", "total_duration_ms",
		"pauses", "initialization", "steps",
		"workflow_name", "workflow_namespace", "workflow_description", "workflow_labels", "workflow_annotations",
		"workflow_created", "workflow_updated", "workflow_spec", "workflow_read_only", "workflow_status",
		"resolved_workflow_name", "resolved_workflow_namespace", "resolved_workflow_description",
		"resolved_workflow_labels", "resolved_workflow_annotations", "resolved_workflow_created",
		"resolved_workflow_updated", "resolved_workflow_spec", "resolved_workflow_read_only", "resolved_workflow_status",
		"signatures_json", "outputs_json", "reports_json", "resource_aggregations_global", "resource_aggregations_step",
	}).AddRow(
		"test-id", "group-1", "runner-1", []byte(`{}`), []byte(`{}`), "test-execution", "default", int64(1),
		time.Now(), time.Now(), time.Now(), "test-execution-name", false,
		[]byte(`{"env":"test"}`), []byte(`{}`), []byte(`{}`), time.Now(), time.Now(),
		"passed", "passed", time.Now(), time.Now(), time.Now(),
		"5m", "5m", int64(300000), int64(0), int64(300000),
		[]byte(`[]`), []byte(`{}`), []byte(`{}`),
		"test-workflow", "default", "Test workflow", []byte(`{}`), []byte(`{}`),
		time.Now(), time.Now(), []byte(`{}`), false, []byte(`{}`),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		[]byte(`[]`), []byte(`[]`), []byte(`[]`), []byte(`{}`), []byte(`{}`),
	)

	workflowNames := []string{"workflow1", "workflow2"}
	mock.ExpectQuery(expectedQuery).WithArgs(workflowNames).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetLatestTestWorkflowExecutionsByTestWorkflows(ctx, workflowNames)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id", result[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetRunningTestWorkflowExecutions(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT 
    e\.id, e\.group_id, e\.runner_id, e\.runner_target, e\.runner_original_target, e\.name, e\.namespace, e\.number, e\.scheduled_at, e\.assigned_at, e\.status_at, e\.test_workflow_execution_name, e\.disable_webhooks, e\.tags, e\.running_context, e\.config_params, e\.created_at, e\.updated_at,
    r\.status, r\.predicted_status, r\.queued_at, r\.started_at, r\.finished_at,
    r\.duration, r\.total_duration, r\.duration_ms, r\.paused_ms, r\.total_duration_ms,
    r\.pauses, r\.initialization, r\.steps,
    w\.name as workflow_name, w\.namespace as workflow_namespace, w\.description as workflow_description,
    w\.labels as workflow_labels, w\.annotations as workflow_annotations, w\.created as workflow_created,
    w\.updated as workflow_updated, w\.spec as workflow_spec, w\.read_only as workflow_read_only,
    w\.status as workflow_status,
    rw\.name as resolved_workflow_name, rw\.namespace as resolved_workflow_namespace, 
    rw\.description as resolved_workflow_description, rw\.labels as resolved_workflow_labels,
    rw\.annotations as resolved_workflow_annotations, rw\.created as resolved_workflow_created,
    rw\.updated as resolved_workflow_updated, rw\.spec as resolved_workflow_spec,
    rw\.read_only as resolved_workflow_read_only, rw\.status as resolved_workflow_status,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', s\.id,
                'ref', s\.ref,
                'name', s\.name,
                'category', s\.category,
                'optional', s\.optional,
                'negative', s\.negative,
                'parent_id', s\.parent_id
            \) ORDER BY s\.id
        \) FROM test_workflow_signatures s WHERE s\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as signatures_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', o\.id,
                'ref', o\.ref,
                'name', o\.name,
                'value', o\.value
            \) ORDER BY o\.id
        \) FROM test_workflow_outputs o WHERE o\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as outputs_json,
    COALESCE\(
        \(SELECT json_agg\(
            json_build_object\(
                'id', rep\.id,
                'ref', rep\.ref,
                'kind', rep\.kind,
                'file', rep\.file,
                'summary', rep\.summary
            \) ORDER BY rep\.id
        \) FROM test_workflow_reports rep WHERE rep\.execution_id = e\.id\),
        '\[\]'::json
    \)::json as reports_json,
    ra\.global as resource_aggregations_global,
    ra\.step as resource_aggregations_step
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
LEFT JOIN test_workflows rw ON e\.id = rw\.execution_id AND rw\.workflow_type = 'resolved_workflow'
LEFT JOIN test_workflow_resource_aggregations ra ON e\.id = ra\.execution_id
WHERE r\.status IN \('queued', 'assigned', 'starting', 'running', 'pausing', 'paused', 'resuming'\)
ORDER BY e\.id DESC`

	rows := mock.NewRows([]string{
		"id", "group_id", "runner_id", "runner_target", "runner_original_target", "name", "namespace", "number",
		"scheduled_at", "assigned_at", "status_at", "test_workflow_execution_name", "disable_webhooks",
		"tags", "running_context", "config_params", "created_at", "updated_at",
		"status", "predicted_status", "queued_at", "started_at", "finished_at",
		"duration", "total_duration", "duration_ms", "paused_ms", "total_duration_ms",
		"pauses", "initialization", "steps",
		"workflow_name", "workflow_namespace", "workflow_description", "workflow_labels", "workflow_annotations",
		"workflow_created", "workflow_updated", "workflow_spec", "workflow_read_only", "workflow_status",
		"resolved_workflow_name", "resolved_workflow_namespace", "resolved_workflow_description",
		"resolved_workflow_labels", "resolved_workflow_annotations", "resolved_workflow_created",
		"resolved_workflow_updated", "resolved_workflow_spec", "resolved_workflow_read_only", "resolved_workflow_status",
		"signatures_json", "outputs_json", "reports_json", "resource_aggregations_global", "resource_aggregations_step",
	}).AddRow(
		"test-id", "group-1", "runner-1", []byte(`{}`), []byte(`{}`), "test-execution", "default", int64(1),
		time.Now(), time.Now(), time.Now(), "test-execution-name", false,
		[]byte(`{"env":"test"}`), []byte(`{}`), []byte(`{}`), time.Now(), time.Now(),
		"running", "running", time.Now(), time.Now(), time.Now(),
		"5m", "5m", int64(300000), int64(0), int64(300000),
		[]byte(`[]`), []byte(`{}`), []byte(`{}`),
		"test-workflow", "default", "Test workflow", []byte(`{}`), []byte(`{}`),
		time.Now(), time.Now(), []byte(`{}`), false, []byte(`{}`),
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
		[]byte(`[]`), []byte(`[]`), []byte(`[]`), []byte(`{}`), []byte(`{}`),
	)

	mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetRunningTestWorkflowExecutions(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id", result[0].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetTestWorkflowExecutionsTotals(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT
	    r\.status,
	    COUNT\(\*\) as count
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
WHERE 1=1       
    AND \(COALESCE\(\$1::text, ''\) = '' OR w.name = \$1::text\)
    AND \(COALESCE\(\$2::text\[\], ARRAY\[\]::text\[\]\) = ARRAY\[\]::text\[\] OR w.name = ANY\(\$2::text\[\]\)\)
    AND \(COALESCE\(\$3::text, ''\) = '' OR e.name ILIKE '%' \|\| \$3::text \|\| '%'\)
    AND \(COALESCE\(\$4::timestamptz, '1900-01-01'::timestamptz\) = '1900-01-01'::timestamptz OR e.scheduled_at >= \$4::timestamptz\)
    AND \(COALESCE\(\$5::timestamptz, '2100-01-01'::timestamptz\) = '2100-01-01'::timestamptz OR e.scheduled_at <= \$5::timestamptz\)
    AND \(COALESCE\(\$6::integer, 0\) = 0 OR e.scheduled_at >= NOW\(\) - \(COALESCE\(\$6::integer, 0\) \|\| ' days'\)::interval\)
    AND \(COALESCE\(\$7::text\[\], ARRAY\[\]::text\[\]\) = ARRAY\[\]::text\[\] OR r.status = ANY\(\$7::text\[\]\)\)
    AND \(COALESCE\(\$8::text, ''\) = '' OR e.runner_id = \$8::text\)
    AND \(COALESCE\(\$9, NULL\) IS NULL OR 
         \(\$9::boolean = true AND e.runner_id IS NOT NULL AND e.runner_id != ''\) OR 
         \(\$9::boolean = false AND \(e.runner_id IS NULL OR e.runner_id = ''\)\)\)
    AND \(COALESCE\(\$10::text, ''\) = '' OR e.running_context->'actor'->>'name' = \$10::text\)
    AND \(COALESCE\(\$11::text, ''\) = '' OR e.running_context->'actor'->>'type_' = \$11::text\)
    AND \(COALESCE\(\$12::text, ''\) = '' OR e.id = \$12::text OR e.group_id = \$12::text\)
    AND \(COALESCE\(\$13, NULL\) IS NULL OR 
         \(\$13::boolean = true AND \(r.status != 'queued' OR r.steps IS NOT NULL\)\) OR
         \(\$13::boolean = false AND r.status = 'queued' AND \(r.steps IS NULL OR r.steps = '\{\}'::jsonb\)\)\)
    AND \(     
        \(COALESCE\(\$14::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$14::jsonb\) AS key_condition
                WHERE 
                CASE 
                    WHEN key_condition->>'operator' = 'not_exists' THEN
                        NOT \(e.tags \? \(key_condition->>'key'\)\)
                    ELSE
                        e.tags \? \(key_condition->>'key'\)
                END
            \) = jsonb_array_length\(\$14::jsonb\)
        \)
        AND
        \(COALESCE\(\$15::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$15::jsonb\) AS condition
                WHERE e.tags->>\(condition->>'key'\) = ANY\(
                    SELECT jsonb_array_elements_text\(condition->'values'\)
                \)
            \) > 0
        \)
    \)
    AND \(
        \(COALESCE\(\$16::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$16::jsonb\) AS key_condition
                WHERE 
                CASE 
                    WHEN key_condition->>'operator' = 'not_exists' THEN
                        NOT \(w.labels \? \(key_condition->>'key'\)\)
                    ELSE
                        w.labels \? \(key_condition->>'key'\)
                END
            \) > 0
        \)
        OR
        \(COALESCE\(\$17::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$17::jsonb\) AS condition
                WHERE w.labels->>\(condition->>'key'\) = ANY\(
                    SELECT jsonb_array_elements_text\(condition->'values'\)
                \)
            \) > 0
        \)
    \)
    AND \(
        \(COALESCE\(\$18::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$18::jsonb\) AS key_condition
                WHERE 
                CASE 
                    WHEN key_condition->>'operator' = 'not_exists' THEN
                        NOT \(w.labels \? \(key_condition->>'key'\)\)
                    ELSE
                        w.labels \? \(key_condition->>'key'\)
                END
            \) = jsonb_array_length\(\$18::jsonb\)
        \)
        AND
        \(COALESCE\(\$19::jsonb, '\[\]'::jsonb\) = '\[\]'::jsonb OR 
            \(SELECT COUNT\(\*\) FROM jsonb_array_elements\(\$19::jsonb\) AS condition
                WHERE w.labels->>\(condition->>'key'\) = ANY\(
                    SELECT jsonb_array_elements_text\(condition->'values'\)
                \)
            \) = jsonb_array_length\(\$19::jsonb\)
        \)
    \)
GROUP BY r\.status`

	rows := mock.NewRows([]string{"status", "count"}).
		AddRow("passed", int64(5)).
		AddRow("failed", int64(2))

	// Create parameters struct with all required fields
	params := GetTestWorkflowExecutionsTotalsParams{
		WorkflowName:       "",
		WorkflowNames:      []string{},
		TextSearch:         "",
		StartDate:          pgtype.Timestamptz{Valid: false},
		EndDate:            pgtype.Timestamptz{Valid: false},
		LastNDays:          0,
		Statuses:           []string{},
		RunnerID:           "",
		Assigned:           pgtype.Bool{Valid: false},
		ActorName:          "",
		ActorType:          "",
		GroupID:            "",
		Initialized:        pgtype.Bool{Valid: false},
		TagKeys:            []byte{},
		TagConditions:      []byte{},
		LabelKeys:          []byte{},
		LabelConditions:    []byte{},
		SelectorKeys:       []byte{},
		SelectorConditions: []byte{},
	}

	mock.ExpectQuery(expectedQuery).WithArgs(
		params.WorkflowName,
		params.WorkflowNames,
		params.TextSearch,
		params.StartDate,
		params.EndDate,
		params.LastNDays,
		params.Statuses,
		params.RunnerID,
		params.Assigned,
		params.ActorName,
		params.ActorType,
		params.GroupID,
		params.Initialized,
		params.TagKeys,
		params.TagConditions,
		params.LabelKeys,
		params.LabelConditions,
		params.SelectorKeys,
		params.SelectorConditions,
	).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetTestWorkflowExecutionsTotals(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "passed", result[0].Status.String)
	assert.Equal(t, int64(5), result[0].Count)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_InsertTestWorkflowExecution(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `INSERT INTO test_workflow_executions \(
    id, group_id, runner_id, runner_target, runner_original_target, name, namespace, number,
    scheduled_at, assigned_at, status_at, test_workflow_execution_name, disable_webhooks, 
    tags, running_context, config_params
\) VALUES \(
    \$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8,
    \$9, \$10, \$11, \$12, \$13,
    \$14, \$15, \$16
\)`

	params := InsertTestWorkflowExecutionParams{
		ID:                        "test-id",
		GroupID:                   pgtype.Text{String: "group-1", Valid: true},
		RunnerID:                  pgtype.Text{String: "runner-1", Valid: true},
		RunnerTarget:              []byte(`{}`),
		RunnerOriginalTarget:      []byte(`{}`),
		Name:                      "test-execution",
		Namespace:                 pgtype.Text{String: "default", Valid: true},
		Number:                    pgtype.Int4{Int32: 1, Valid: true},
		ScheduledAt:               pgtype.Timestamptz{Time: time.Now(), Valid: true},
		AssignedAt:                pgtype.Timestamptz{Valid: false},
		StatusAt:                  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		TestWorkflowExecutionName: pgtype.Text{Valid: false},
		DisableWebhooks:           pgtype.Bool{Bool: false, Valid: true},
		Tags:                      []byte(`{"env":"test"}`),
		RunningContext:            []byte(`{}`),
		ConfigParams:              []byte(`{}`),
	}

	mock.ExpectExec(expectedQuery).WithArgs(
		params.ID,
		params.GroupID,
		params.RunnerID,
		params.RunnerTarget,
		params.RunnerOriginalTarget,
		params.Name,
		params.Namespace,
		params.Number,
		params.ScheduledAt,
		params.AssignedAt,
		params.StatusAt,
		params.TestWorkflowExecutionName,
		params.DisableWebhooks,
		params.Tags,
		params.RunningContext,
		params.ConfigParams,
	).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	// Execute query
	err = queries.InsertTestWorkflowExecution(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_UpdateTestWorkflowExecutionResult(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `UPDATE test_workflow_results 
SET 
    status = \$1,
    predicted_status = \$2,
    queued_at = \$3,
    started_at = \$4,
    finished_at = \$5,
    duration = \$6,
    total_duration = \$7,
    duration_ms = \$8,
    paused_ms = \$9,
    total_duration_ms = \$10,
    pauses = \$11,
    initialization = \$12,
    steps = \$13
WHERE execution_id = \$14`

	params := UpdateTestWorkflowExecutionResultParams{
		Status:          pgtype.Text{String: "passed", Valid: true},
		PredictedStatus: pgtype.Text{String: "passed", Valid: true},
		QueuedAt:        pgtype.Timestamptz{Time: time.Now(), Valid: true},
		StartedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		FinishedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Duration:        pgtype.Text{String: "5m", Valid: true},
		TotalDuration:   pgtype.Text{String: "5m", Valid: true},
		DurationMs:      pgtype.Int4{Int32: 300000, Valid: true},
		PausedMs:        pgtype.Int4{Int32: 0, Valid: true},
		TotalDurationMs: pgtype.Int4{Int32: 300000, Valid: true},
		Pauses:          []byte(`[]`),
		Initialization:  []byte(`{}`),
		Steps:           []byte(`{}`),
		ExecutionID:     "test-id",
	}

	mock.ExpectExec(expectedQuery).WithArgs(
		params.Status,
		params.PredictedStatus,
		params.QueuedAt,
		params.StartedAt,
		params.FinishedAt,
		params.Duration,
		params.TotalDuration,
		params.DurationMs,
		params.PausedMs,
		params.TotalDurationMs,
		params.Pauses,
		params.Initialization,
		params.Steps,
		params.ExecutionID,
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// Execute query
	err = queries.UpdateTestWorkflowExecutionResult(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_DeleteTestWorkflowExecutionsByTestWorkflow(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `DELETE FROM test_workflow_executions e
USING test_workflows w
WHERE e\.id = w\.execution_id 
  AND w\.workflow_type = 'workflow' 
  AND w\.name = \$1`

	mock.ExpectExec(expectedQuery).WithArgs("test-workflow").WillReturnResult(pgxmock.NewResult("DELETE", 1))

	// Execute query
	err = queries.DeleteTestWorkflowExecutionsByTestWorkflow(ctx, "test-workflow")

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_DeleteAllTestWorkflowExecutions(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `DELETE FROM test_workflow_executions`

	mock.ExpectExec(expectedQuery).WillReturnResult(pgxmock.NewResult("DELETE", 5))

	// Execute query
	err = queries.DeleteAllTestWorkflowExecutions(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_AssignTestWorkflowExecution(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `UPDATE test_workflow_executions 
SET 
    runner_id = \$1::text,
    assigned_at = \$2
FROM test_workflow_results r
WHERE test_workflow_executions\.id = \$3
    AND test_workflow_executions\.id = r\.execution_id
    AND r\.status = 'queued'
    AND \(\(test_workflow_executions\.runner_id IS NULL OR test_workflow_executions\.runner_id = ''\)
         OR \(test_workflow_executions\.runner_id = \$1::text AND assigned_at < \$2\)
         OR \(test_workflow_executions\.runner_id = \$4::text AND assigned_at < NOW\(\) - INTERVAL '1 minute' AND assigned_at < \$2\)\)
RETURNING test_workflow_executions\.id`

	params := AssignTestWorkflowExecutionParams{
		NewRunnerID:  "new-runner",
		AssignedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		ID:           "test-id",
		PrevRunnerID: "old-runner",
	}

	rows := mock.NewRows([]string{"id"}).AddRow("test-id")
	mock.ExpectQuery(expectedQuery).WithArgs(
		params.NewRunnerID,
		params.AssignedAt,
		params.ID,
		params.PrevRunnerID,
	).WillReturnRows(rows)

	// Execute query
	result, err := queries.AssignTestWorkflowExecution(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_AbortTestWorkflowExecutionIfQueued(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `UPDATE test_workflow_executions 
SET status_at = \$1
FROM test_workflow_results r
WHERE test_workflow_executions\.id = \$2
    AND test_workflow_executions\.id = r\.execution_id
    AND r\.status IN \('queued', 'assigned', 'starting', 'running', 'paused', 'resuming'\)
    AND \(test_workflow_executions\.runner_id IS NULL OR test_workflow_executions\.runner_id = ''\)
RETURNING test_workflow_executions\.id`

	params := AbortTestWorkflowExecutionIfQueuedParams{
		AbortTime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		ID:        "test-id",
	}

	rows := mock.NewRows([]string{"id"}).AddRow("test-id")
	mock.ExpectQuery(expectedQuery).WithArgs(params.AbortTime, params.ID).WillReturnRows(rows)

	// Execute query
	result, err := queries.AbortTestWorkflowExecutionIfQueued(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "test-id", result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_AbortTestWorkflowResultIfQueued(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `UPDATE test_workflow_results 
SET 
    status = 'aborted',
    predicted_status = 'aborted',
    finished_at = \$1,
    initialization = jsonb_set\(
        jsonb_set\(
            jsonb_set\(COALESCE\(initialization, '\{\}'::jsonb\), '\{status\}', '"aborted"'\),
            '\{errormessage\}', '"Aborted before initialization\."'
        \),
        '\{finishedat\}', to_jsonb\(\$1::timestamp\)
    \)
WHERE execution_id = \$2
    AND status IN \('queued', 'running', 'paused'\)`

	params := AbortTestWorkflowResultIfQueuedParams{
		AbortTime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		ID:        "test-id",
	}

	mock.ExpectExec(expectedQuery).WithArgs(params.AbortTime, params.ID).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// Execute query
	err = queries.AbortTestWorkflowResultIfQueued(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetPreviousFinishedState(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT r\.status
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
WHERE w\.name = \$1::text
    AND r\.finished_at < \$2
    AND r\.status IN \('passed', 'failed', 'skipped', 'aborted', 'canceled', 'timeout'\)
ORDER BY r\.finished_at DESC
LIMIT 1`

	params := GetPreviousFinishedStateParams{
		WorkflowName: "test-workflow",
		Date:         pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	rows := mock.NewRows([]string{"status"}).AddRow("passed")
	mock.ExpectQuery(expectedQuery).WithArgs(params.WorkflowName, params.Date).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetPreviousFinishedState(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "passed", result.String)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetTestWorkflowExecutionTags(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `WITH tag_extracts AS \(
SELECT 
        e.id,
        w.name as workflow_name,
        tag_pair.key as tag_key,
        tag_pair.value as tag_value
    FROM test_workflow_executions e
    LEFT JOIN test_workflows w ON e.id = w.execution_id AND w.workflow_type = 'workflow'
    CROSS JOIN LATERAL jsonb_each_text\(e.tags\) AS tag_pair\(key, value\)
    WHERE e.tags IS NOT NULL 
        AND e.tags != '\{\}'::jsonb
        AND jsonb_typeof\(e.tags\) = 'object'
\)
SELECT 
    tag_key::text,
    array_agg\(DISTINCT tag_value ORDER BY tag_value\)::text\[\] as values
FROM tag_extracts
WHERE \(COALESCE\(\$1::text, ''\) = '' OR workflow_name = \$1::text\)
GROUP BY tag_key
ORDER BY tag_key`

	rows := mock.NewRows([]string{"tag_key", "values"}).
		AddRow("env", []string{"test", "prod"}).
		AddRow("version", []string{"1.0", "2.0"})

	mock.ExpectQuery(expectedQuery).WithArgs("test-workflow").WillReturnRows(rows)

	// Execute query
	result, err := queries.GetTestWorkflowExecutionTags(ctx, "test-workflow")

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "env", result[0].TagKey)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_GetTestWorkflowMetrics(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `SELECT 
    e\.id as execution_id,
    e\.group_id,
    r\.duration,
    r\.duration_ms,
    r\.status,
    e\.name,
    e\.scheduled_at as start_time,
    e\.runner_id
FROM test_workflow_executions e
LEFT JOIN test_workflow_results r ON e\.id = r\.execution_id
LEFT JOIN test_workflows w ON e\.id = w\.execution_id AND w\.workflow_type = 'workflow'
WHERE w\.name = \$1::text
    AND \(\$2::integer = 0 OR e\.scheduled_at >= NOW\(\) - \(\$2::integer \|\| ' days'\)::interval\)
ORDER BY e\.scheduled_at DESC
LIMIT \$3`

	params := GetTestWorkflowMetricsParams{
		WorkflowName: "test-workflow",
		LastNDays:    7,
		Lmt:          10,
	}

	rows := mock.NewRows([]string{
		"execution_id", "group_id", "duration", "duration_ms", "status", "name", "start_time", "runner_id",
	}).AddRow(
		"exec-1", "group-1", "5m", int64(300000), "passed", "test-execution", time.Now(), "runner-1",
	)

	mock.ExpectQuery(expectedQuery).WithArgs(params.WorkflowName, params.LastNDays, params.Lmt).WillReturnRows(rows)

	// Execute query
	result, err := queries.GetTestWorkflowMetrics(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "exec-1", result[0].ExecutionID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_InsertTestWorkflowSignature(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `INSERT INTO test_workflow_signatures \(
    execution_id, ref, name, category, optional, negative, parent_id
\) VALUES \(
    \$1, \$2, \$3, \$4, \$5, \$6, \$7
\)
RETURNING test_workflow_signatures\.id`

	params := InsertTestWorkflowSignatureParams{
		ExecutionID: "test-id",
		Ref:         pgtype.Text{String: "step1", Valid: true},
		Name:        pgtype.Text{String: "Test Step", Valid: true},
		Category:    pgtype.Text{String: "test", Valid: true},
		Optional:    pgtype.Bool{Bool: false, Valid: true},
		Negative:    pgtype.Bool{Bool: false, Valid: true},
		ParentID:    pgtype.Int4{Valid: false},
	}

	rows := mock.NewRows([]string{"id"}).AddRow(int32(1))
	mock.ExpectQuery(expectedQuery).WithArgs(
		params.ExecutionID,
		params.Ref,
		params.Name,
		params.Category,
		params.Optional,
		params.Negative,
		params.ParentID,
	).WillReturnRows(rows)

	// Execute query
	result, err := queries.InsertTestWorkflowSignature(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, int32(1), result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSQLCTestWorkflowExecutionQueries_UpdateTestWorkflowExecution(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	queries := New(mock)
	ctx := context.Background()

	expectedQuery := `UPDATE test_workflow_executions
SET
    group_id = \$1,
    runner_id = \$2,
    runner_target = \$3,
    runner_original_target = \$4,
    name = \$5,
    namespace = \$6,
    number = \$7,
    scheduled_at = \$8,
    assigned_at = \$9,
    status_at = \$10,
    test_workflow_execution_name = \$11,
    disable_webhooks = \$12,
    tags = \$13,
    running_context = \$14,
    config_params = \$15
WHERE id = \$16`

	params := UpdateTestWorkflowExecutionParams{
		GroupID:                   pgtype.Text{String: "group-1", Valid: true},
		RunnerID:                  pgtype.Text{String: "runner-1", Valid: true},
		RunnerTarget:              []byte(`{}`),
		RunnerOriginalTarget:      []byte(`{}`),
		Name:                      "updated-execution",
		Namespace:                 pgtype.Text{String: "default", Valid: true},
		Number:                    pgtype.Int4{Int32: 2, Valid: true},
		ScheduledAt:               pgtype.Timestamptz{Time: time.Now(), Valid: true},
		AssignedAt:                pgtype.Timestamptz{Time: time.Now(), Valid: true},
		StatusAt:                  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		TestWorkflowExecutionName: pgtype.Text{String: "test-execution", Valid: true},
		DisableWebhooks:           pgtype.Bool{Bool: false, Valid: true},
		Tags:                      []byte(`{"env":"prod"}`),
		RunningContext:            []byte(`{}`),
		ConfigParams:              []byte(`{}`),
		ID:                        "test-id",
	}

	mock.ExpectExec(expectedQuery).WithArgs(
		params.GroupID,
		params.RunnerID,
		params.RunnerTarget,
		params.RunnerOriginalTarget,
		params.Name,
		params.Namespace,
		params.Number,
		params.ScheduledAt,
		params.AssignedAt,
		params.StatusAt,
		params.TestWorkflowExecutionName,
		params.DisableWebhooks,
		params.Tags,
		params.RunningContext,
		params.ConfigParams,
		params.ID,
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// Execute query
	err = queries.UpdateTestWorkflowExecution(ctx, params)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
