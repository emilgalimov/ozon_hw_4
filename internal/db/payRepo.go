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
	const sql = `INSERT INTO storage(order_id, product_id, is_reserved) VALUES ($1, $2, $3)`

	ctx := context.Background()

	r.pool.QueryRow(ctx, sql, pay.OrderID, pay.Sum, pay.CreatedAt)
}
