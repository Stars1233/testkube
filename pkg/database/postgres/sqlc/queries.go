package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TestWorkflowExecutionQueriesInterface defines the interface for sqlc generated queries
type TestWorkflowExecutionQueriesInterface interface {
	// Transaction methods
	WithTx(tx pgx.Tx) TestWorkflowExecutionQueriesInterface

	// TestWorkflowExecution queries
	GetTestWorkflowExecution(ctx context.Context, id string) (GetTestWorkflowExecutionRow, error)
	GetTestWorkflowExecutionByNameAndTestWorkflow(ctx context.Context, arg GetTestWorkflowExecutionByNameAndTestWorkflowParams) (GetTestWorkflowExecutionByNameAndTestWorkflowRow, error)
	GetLatestTestWorkflowExecutionByTestWorkflow(ctx context.Context, arg GetLatestTestWorkflowExecutionByTestWorkflowParams) (GetLatestTestWorkflowExecutionByTestWorkflowRow, error)
	GetLatestTestWorkflowExecutionsByTestWorkflows(ctx context.Context, workflowNames []string) ([]GetLatestTestWorkflowExecutionsByTestWorkflowsRow, error)
	GetRunningTestWorkflowExecutions(ctx context.Context) ([]GetRunningTestWorkflowExecutionsRow, error)
	GetTestWorkflowExecutionsTotals(ctx context.Context, arg GetTestWorkflowExecutionsTotalsParams) ([]GetTestWorkflowExecutionsTotalsRow, error)
	GetTestWorkflowExecutions(ctx context.Context, arg GetTestWorkflowExecutionsParams) ([]GetTestWorkflowExecutionsRow, error)
	GetTestWorkflowExecutionsSummary(ctx context.Context, arg GetTestWorkflowExecutionsSummaryParams) ([]GetTestWorkflowExecutionsSummaryRow, error)
	GetFinishedTestWorkflowExecutions(ctx context.Context, arg GetFinishedTestWorkflowExecutionsParams) ([]GetFinishedTestWorkflowExecutionsRow, error)
	GetUnassignedTestWorkflowExecutions(ctx context.Context) ([]GetUnassignedTestWorkflowExecutionsRow, error)
	CountTestWorkflowExecutions(ctx context.Context, arg CountTestWorkflowExecutionsParams) (int64, error)

	// Insert operations
	InsertTestWorkflowExecution(ctx context.Context, arg InsertTestWorkflowExecutionParams) error
	InsertTestWorkflowResult(ctx context.Context, arg InsertTestWorkflowResultParams) error
	InsertTestWorkflowSignature(ctx context.Context, arg InsertTestWorkflowSignatureParams) (int32, error)
	InsertTestWorkflowOutput(ctx context.Context, arg InsertTestWorkflowOutputParams) error
	InsertTestWorkflowReport(ctx context.Context, arg InsertTestWorkflowReportParams) error
	InsertTestWorkflowResourceAggregations(ctx context.Context, arg InsertTestWorkflowResourceAggregationsParams) error
	InsertTestWorkflow(ctx context.Context, arg InsertTestWorkflowParams) error

	// Update operations
	UpdateTestWorkflowExecution(ctx context.Context, arg UpdateTestWorkflowExecutionParams) error
	UpdateTestWorkflowExecutionResult(ctx context.Context, arg UpdateTestWorkflowExecutionResultParams) error
	UpdateExecutionStatusAt(ctx context.Context, arg UpdateExecutionStatusAtParams) error
	UpdateTestWorkflowExecutionReport(ctx context.Context, arg UpdateTestWorkflowExecutionReportParams) error
	UpdateTestWorkflowExecutionResourceAggregations(ctx context.Context, arg UpdateTestWorkflowExecutionResourceAggregationsParams) error

	// Delete operations
	DeleteTestWorkflowSignatures(ctx context.Context, executionID string) error
	DeleteTestWorkflowResult(ctx context.Context, executionID string) error
	DeleteTestWorkflowOutputs(ctx context.Context, executionID string) error
	DeleteTestWorkflowReports(ctx context.Context, executionID string) error
	DeleteTestWorkflowResourceAggregations(ctx context.Context, executionID string) error
	DeleteTestWorkflow(ctx context.Context, arg DeleteTestWorkflowParams) error
	DeleteTestWorkflowExecutionsByTestWorkflow(ctx context.Context, workflowName string) error
	DeleteAllTestWorkflowExecutions(ctx context.Context) error
	DeleteTestWorkflowExecutionsByTestWorkflows(ctx context.Context, workflowNames []string) error

	// Metrics and analytics
	GetTestWorkflowMetrics(ctx context.Context, arg GetTestWorkflowMetricsParams) ([]GetTestWorkflowMetricsRow, error)
	GetPreviousFinishedState(ctx context.Context, arg GetPreviousFinishedStateParams) (pgtype.Text, error)
	GetTestWorkflowExecutionTags(ctx context.Context, workflowName string) ([]GetTestWorkflowExecutionTagsRow, error)

	// Execution management
	InitTestWorkflowExecution(ctx context.Context, arg InitTestWorkflowExecutionParams) error
	AssignTestWorkflowExecution(ctx context.Context, arg AssignTestWorkflowExecutionParams) (string, error)
	AbortTestWorkflowExecutionIfQueued(ctx context.Context, arg AbortTestWorkflowExecutionIfQueuedParams) (string, error)
	AbortTestWorkflowResultIfQueued(ctx context.Context, arg AbortTestWorkflowResultIfQueuedParams) error
}

// Ensure Queries implements TestWorkflowExecutionQueriesInterface
var _ TestWorkflowExecutionQueriesInterface = (*SQLCTestWorkflowExecutionQueriesWrapper)(nil)

// SQLCTestWorkflowExecutionQueriesWrapper wraps Queries to implement TestWorkflowExecutionQueriesInterface
type SQLCTestWorkflowExecutionQueriesWrapper struct {
	*Queries
}

// NewSQLCTestWorkflowExecutionQueriesWrapper creates a new wrapper for Queries
func NewSQLCTestWorkflowExecutionQueriesWrapper(queries *Queries) TestWorkflowExecutionQueriesInterface {
	return &SQLCTestWorkflowExecutionQueriesWrapper{Queries: queries}
}

// WithTx returns a new TestWorkflowExecutionQueriesInterface with transaction
func (w *SQLCTestWorkflowExecutionQueriesWrapper) WithTx(tx pgx.Tx) TestWorkflowExecutionQueriesInterface {
	return &SQLCTestWorkflowExecutionQueriesWrapper{Queries: w.Queries.WithTx(tx)}
}

// DatabaseInterface defines the interface for database operations
type DatabaseInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

// PgxPoolWrapper wraps pgxpool.Pool to implement DatabaseInterface
type PgxPoolWrapper struct {
	*pgxpool.Pool
}

func (w *PgxPoolWrapper) Begin(ctx context.Context) (pgx.Tx, error) {
	return w.Pool.Begin(ctx)
}

// LeaseBackendQueriesInterface defines the interface for sqlc generated queries
type LeaseBackendQueriesInterface interface {
	FindLeaseById(ctx context.Context, leaseID string) (Lease, error)
	InsertLease(ctx context.Context, arg InsertLeaseParams) (Lease, error)
	UpdateLease(ctx context.Context, arg UpdateLeaseParams) (Lease, error)
}

// ExecutionSequenceQueriesInterface defines the interface for sqlc generated queries
type ExecutionSequenceQueriesInterface interface {
	UpsertAndIncrementExecutionSequence(ctx context.Context, name string) (ExecutionSequence, error)
	DeleteExecutionSequence(ctx context.Context, name string) error
	DeleteExecutionSequences(ctx context.Context, names []string) error
	DeleteAllExecutionSequences(ctx context.Context) error
}
