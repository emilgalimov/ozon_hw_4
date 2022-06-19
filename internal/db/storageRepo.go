package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"
)

type cache interface {
	Get(orderID int) ([]storage.StoreItem, error)
	Set(orderID int, items []storage.StoreItem) error
	Delete(orderID int) error
}

type StorageRepo struct {
	pool  *pgxpool.Pool
	cache cache
}

func NewStorageRepo(pool *pgxpool.Pool, cache cache) *StorageRepo {
	return &StorageRepo{
		pool:  pool,
		cache: cache,
	}
}

func (s StorageRepo) GetUnreservedByOrderID(orderID int) (items []storage.StoreItem) {

	items, err := s.cache.Get(orderID)
	if err == nil {
		return items
	}

	//language=PostgreSQL
	const sql = `SELECT order_id, product_id, is_reserved FROM storage WHERE order_id = $1`

	ctx := context.Background()
	row, _ := s.pool.Query(ctx, sql, orderID)

	for row.Next() {
		item := storage.StoreItem{}

		_ = row.Scan(
			&item.OrderID,
			&item.ProductID,
			&item.IsReserved,
		)
		items = append(items, item)
	}

	_ = s.cache.Set(orderID, items)

	return items
}

func (s StorageRepo) ReserveByOrderID(orderID int) {
	//language=PostgreSQL
	const sql = `UPDATE storage SET is_reserved = TRUE WHERE order_id = $1`

	ctx := context.Background()
	_, _ = s.pool.Query(ctx, sql, orderID)
	_ = s.cache.Delete(orderID)
}

func (s StorageRepo) UnreserveByOrderID(orderID int) {
	//language=PostgreSQL
	const sql = `UPDATE storage SET is_reserved = FALSE WHERE order_id = $1`

	ctx := context.Background()
	_, _ = s.pool.Query(ctx, sql, orderID)
	_ = s.cache.Delete(orderID)
}
