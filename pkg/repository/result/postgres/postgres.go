package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/repository/result"
)

var _ result.Repository = (*PostgresRepository)(nil)

const (
	PageDefaultLimit = 100
)

type PostgresRepository struct {
	db *pgxpool.Pool
	// queries *sqlc.Queries // This would be added when implementing
}

type PostgresRepositoryOpt func(*PostgresRepository)

func NewPostgresRepository(db *pgxpool.Pool, opts ...PostgresRepositoryOpt) *PostgresRepository {
	r := &PostgresRepository{
		db: db,
		// queries: sqlc.New(db), // This would be added when implementing
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Sequences interface implementation

// GetNextExecutionNumber gets next execution number by name
func (r *PostgresRepository) GetNextExecutionNumber(ctx context.Context, name string) (int32, error) {
	return 0, errors.New("GetNextExecutionNumber not implemented")
}

// Repository interface implementation

// Get gets execution result by id or name
func (r *PostgresRepository) Get(ctx context.Context, id string) (testkube.Execution, error) {
	return testkube.Execution{}, errors.New("Get not implemented")
}

// GetExecution gets execution result without output
func (r *PostgresRepository) GetExecution(ctx context.Context, id string) (testkube.Execution, error) {
	return testkube.Execution{}, errors.New("GetExecution not implemented")
}

// GetByNameAndTest gets execution result by name and test name
func (r *PostgresRepository) GetByNameAndTest(ctx context.Context, name, testName string) (testkube.Execution, error) {
	return testkube.Execution{}, errors.New("GetByNameAndTest not implemented")
}

// GetPreviousFinishedState gets previous finished execution state by test
func (r *PostgresRepository) GetPreviousFinishedState(ctx context.Context, testName string, date time.Time) (testkube.ExecutionStatus, error) {
	return "", errors.New("GetPreviousFinishedState not implemented")
}

// GetLatestByTest gets latest execution result by test
func (r *PostgresRepository) GetLatestByTest(ctx context.Context, testName string) (*testkube.Execution, error) {
	return nil, errors.New("GetLatestByTest not implemented")
}

// GetLatestByTests gets latest execution results by test names
func (r *PostgresRepository) GetLatestByTests(ctx context.Context, testNames []string) ([]testkube.Execution, error) {
	return nil, errors.New("GetLatestByTests not implemented")
}

// GetExecutions gets executions using a filter, use filter with no data for all
func (r *PostgresRepository) GetExecutions(ctx context.Context, filter result.Filter) ([]testkube.Execution, error) {
	// Support old reconciler
	return nil, nil
}

// GetExecutionTotals gets the statistics on number of executions using a filter, but without paging
func (r *PostgresRepository) GetExecutionTotals(ctx context.Context, paging bool, filter ...result.Filter) (testkube.ExecutionsTotals, error) {
	return testkube.ExecutionsTotals{}, errors.New("GetExecutionTotals not implemented")
}

// Insert inserts new execution result
func (r *PostgresRepository) Insert(ctx context.Context, result testkube.Execution) error {
	return errors.New("Insert not implemented")
}

// Update updates execution result
func (r *PostgresRepository) Update(ctx context.Context, result testkube.Execution) error {
	return errors.New("Update not implemented")
}

// UpdateResult updates result in execution
func (r *PostgresRepository) UpdateResult(ctx context.Context, id string, execution testkube.Execution) error {
	return errors.New("UpdateResult not implemented")
}

// StartExecution updates execution start time
func (r *PostgresRepository) StartExecution(ctx context.Context, id string, startTime time.Time) error {
	return errors.New("StartExecution not implemented")
}

// EndExecution updates execution end time
func (r *PostgresRepository) EndExecution(ctx context.Context, execution testkube.Execution) error {
	return errors.New("EndExecution not implemented")
}

// GetLabels get all available labels
func (r *PostgresRepository) GetLabels(ctx context.Context) (map[string][]string, error) {
	return nil, errors.New("GetLabels not implemented")
}

// DeleteByTest deletes execution results by test
func (r *PostgresRepository) DeleteByTest(ctx context.Context, testName string) error {
	return errors.New("DeleteByTest not implemented")
}

// DeleteByTestSuite deletes execution results by test suite
func (r *PostgresRepository) DeleteByTestSuite(ctx context.Context, testSuiteName string) error {
	return errors.New("DeleteByTestSuite not implemented")
}

// DeleteAll deletes execution results
func (r *PostgresRepository) DeleteAll(ctx context.Context) error {
	return errors.New("DeleteAll not implemented")
}

// DeleteByTests deletes execution results by tests
func (r *PostgresRepository) DeleteByTests(ctx context.Context, testNames []string) error {
	return errors.New("DeleteByTests not implemented")
}

// DeleteByTestSuites deletes execution results by test suites
func (r *PostgresRepository) DeleteByTestSuites(ctx context.Context, testSuiteNames []string) error {
	return errors.New("DeleteByTestSuites not implemented")
}

// DeleteForAllTestSuites deletes execution results for all test suites
func (r *PostgresRepository) DeleteForAllTestSuites(ctx context.Context) error {
	return errors.New("DeleteForAllTestSuites not implemented")
}

// GetTestMetrics returns metrics for test
func (r *PostgresRepository) GetTestMetrics(ctx context.Context, name string, limit, last int) (testkube.ExecutionsMetrics, error) {
	return testkube.ExecutionsMetrics{}, errors.New("GetTestMetrics not implemented")
}

// Count returns executions count
func (r *PostgresRepository) Count(ctx context.Context, filter result.Filter) (int64, error) {
	return 0, errors.New("Count not implemented")
}
