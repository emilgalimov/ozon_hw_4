package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type StorageRepo struct {
	pool *pgxpool.Pool
}

func NewStorageRepo(pool *pgxpool.Pool) *StorageRepo {
	return &StorageRepo{
		pool: pool,
	}
}
