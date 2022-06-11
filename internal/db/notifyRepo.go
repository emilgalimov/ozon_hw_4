package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-4/internal/model/notify"
)

type NotifyRepo struct {
	pool *pgxpool.Pool
}

func NewNotifyRepo(pool *pgxpool.Pool) *NotifyRepo {
	return &NotifyRepo{
		pool: pool,
	}
}

func (r NotifyRepo) Save(notification notify.SuccessNotification) {
	//language=PostgreSQL
	const sql = `INSERT INTO success_notifications(order_id, message, created_at) VALUES ($1, $2, $3)`

	ctx := context.Background()

	r.pool.QueryRow(ctx, sql, notification.OrderID, notification.Message, notification.CreatedAt)
}
