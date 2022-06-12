package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/storage"
)

type StorageRepo struct {
	pool *pgxpool.Pool
}

func NewStorageRepo(pool *pgxpool.Pool) *StorageRepo {
	return &StorageRepo{
		pool: pool,
	}
}

func (s StorageRepo) GetUnreservedByOrderID(orderID int) []storage.StoreItem {
	//language=PostgreSQL
	const sql = `SELECT order_id, product_id, is_reserved FROM storage WHERE order_id = $1`

	ctx := context.Background()
	row, _ := s.pool.Query(ctx, sql, orderID)

	var items []storage.StoreItem
	for row.Next() {
		item := storage.StoreItem{}

		_ = row.Scan(
			&item.OrderID,
			&item.ProductID,
			&item.IsReserved,
		)
		items = append(items, item)
	}
	return items
}

func (s StorageRepo) ReserveByOrderID(orderID int) {
	//language=PostgreSQL
	const sql = `UPDATE storage SET is_reserved = TRUE WHERE order_id = $1`

	ctx := context.Background()
	_, _ = s.pool.Query(ctx, sql, orderID)
}

func (s StorageRepo) UnreserveByOrderID(orderID int) {
	//language=PostgreSQL
	const sql = `UPDATE storage SET is_reserved = FALSE WHERE order_id = $1`

	ctx := context.Background()
	_, _ = s.pool.Query(ctx, sql, orderID)
}
