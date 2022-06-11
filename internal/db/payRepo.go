package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/pay"
)

type PayRepo struct {
	pool *pgxpool.Pool
}

func NewPayRepo(pool *pgxpool.Pool) *PayRepo {
	return &PayRepo{
		pool: pool,
	}
}

func (r PayRepo) Save(pay pay.Payment) {
	//language=PostgreSQL
	const sql = `INSERT INTO payments(order_id, sum, created_at) VALUES ($1, $2, $3)`

	ctx := context.Background()

	_, _ = r.pool.Query(ctx, sql, pay.OrderID, pay.Sum, pay.CreatedAt)
}
