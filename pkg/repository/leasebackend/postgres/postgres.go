package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kubeshop/testkube/pkg/database/postgres/sqlc"
	"github.com/kubeshop/testkube/pkg/repository/leasebackend"
)

const (
	documentType = "lease"
)

type PostgresLeaseBackend struct {
	db      sqlc.DatabaseInterface
	queries sqlc.LeaseBackendQueriesInterface
}

type PostgresLeaseBackendOpt func(*PostgresLeaseBackend)

func NewPostgresLeaseBackend(db *pgxpool.Pool, opts ...PostgresLeaseBackendOpt) *PostgresLeaseBackend {
	b := &PostgresLeaseBackend{
		db:      &PgxPoolWrapper{Pool: db},
		queries: sqlc.New(db),
	}

	for _, opt := range opts {
		opt(b)
	}

	return b
}

func WithQueriesInterface(queries sqlc.LeaseBackendQueriesInterface) PostgresLeaseBackendOpt {
	return func(b *PostgresLeaseBackend) {
		b.queries = queries
	}
}

func WithDatabaseInterface(db sqlc.DatabaseInterface) PostgresLeaseBackendOpt {
	return func(b *PostgresLeaseBackend) {
		b.db = db
	}
}

// PgxPoolWrapper wraps pgxpool.Pool to implement DatabaseInterface
type PgxPoolWrapper struct {
	*pgxpool.Pool
}

func (w *PgxPoolWrapper) Begin(ctx context.Context) (pgx.Tx, error) {
	return w.Pool.Begin(ctx)
}

// TryAcquire tries to acquire a lease for the given identifier and cluster ID
func (b *PostgresLeaseBackend) TryAcquire(ctx context.Context, id, clusterID string) (bool, error) {
	leaseID := newLeaseID(clusterID)

	currentLease, err := b.findOrInsertCurrentLease(ctx, leaseID, id, clusterID)
	if err != nil {
		return false, err
	}

	acquired, renewable := leaseStatus(currentLease, id, clusterID)
	switch {
	case acquired:
		return true, nil
	case !renewable:
		return false, nil
	}

	acquiredAt := currentLease.AcquiredAt
	if currentLease.Identifier != id {
		acquiredAt = toPgTimestamp(time.Now())
	}

	newLease, err := b.tryUpdateLease(ctx, leaseID, id, clusterID, acquiredAt.Time)
	if err != nil {
		return false, err
	}

	acquired, _ = leaseStatus(newLease, id, clusterID)
	return acquired, nil
}

func (b *PostgresLeaseBackend) findOrInsertCurrentLease(ctx context.Context, leaseID, id, clusterID string) (*sqlc.Lease, error) {
	lease, err := b.queries.FindLeaseById(ctx, leaseID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Lease doesn't exist, try to insert it
			newLease, insertErr := b.insertLease(ctx, leaseID, id, clusterID)
			if insertErr != nil {
				return nil, fmt.Errorf("error inserting lease: %w", insertErr)
			}
			return newLease, nil
		}
		return nil, fmt.Errorf("error finding lease: %w", err)
	}

	return &sqlc.Lease{
		Identifier: lease.Identifier,
		ClusterID:  lease.ClusterID,
		AcquiredAt: lease.AcquiredAt,
		RenewedAt:  lease.RenewedAt,
	}, nil
}

func (b *PostgresLeaseBackend) insertLease(ctx context.Context, leaseID, id, clusterID string) (*sqlc.Lease, error) {
	current := time.Now()

	result, err := b.queries.InsertLease(ctx, sqlc.InsertLeaseParams{
		ID:         leaseID,
		Identifier: id,
		ClusterID:  clusterID,
		AcquiredAt: toPgTimestamp(current),
		RenewedAt:  toPgTimestamp(current),
	})
	if err != nil {
		return nil, fmt.Errorf("error inserting lease: %w", err)
	}

	return &sqlc.Lease{
		Identifier: result.Identifier,
		ClusterID:  result.ClusterID,
		AcquiredAt: result.AcquiredAt,
		RenewedAt:  result.RenewedAt,
	}, nil
}

func (b *PostgresLeaseBackend) tryUpdateLease(ctx context.Context, leaseID, id, clusterID string, acquiredAt time.Time) (*sqlc.Lease, error) {
	renewedAt := time.Now()

	result, err := b.queries.UpdateLease(ctx, sqlc.UpdateLeaseParams{
		ID:         leaseID,
		Identifier: id,
		ClusterID:  clusterID,
		AcquiredAt: toPgTimestamp(acquiredAt),
		RenewedAt:  toPgTimestamp(renewedAt),
	})
	if err != nil {
		return nil, fmt.Errorf("error updating lease: %w", err)
	}

	return &sqlc.Lease{
		Identifier: result.Identifier,
		ClusterID:  result.ClusterID,
		AcquiredAt: result.AcquiredAt,
		RenewedAt:  result.RenewedAt,
	}, nil
}

// Helper functions
func leaseStatus(lease *sqlc.Lease, id, clusterID string) (acquired bool, renewable bool) {
	if lease == nil {
		return false, false
	}

	maxLeaseDurationStaleness := time.Now().Add(-leasebackend.DefaultMaxLeaseDuration)
	isLeaseExpired := lease.RenewedAt.Time.Before(maxLeaseDurationStaleness)
	isMyLease := lease.Identifier == id && lease.ClusterID == clusterID

	switch {
	case isLeaseExpired:
		acquired = false
		renewable = true
	case isMyLease:
		acquired = true
		renewable = false
	default:
		acquired = false
		renewable = false
	}
	return
}

func newLeaseID(clusterID string) string {
	return fmt.Sprintf("%s-%s", documentType, clusterID)
}

// Type conversion helpers
func toPgTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}
