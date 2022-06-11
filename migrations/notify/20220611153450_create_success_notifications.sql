-- +goose Up
-- +goose StatementBegin
CREATE TABLE success_notifications(
    order_id int,
    message varchar,
    created_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE success_notifications;
-- +goose StatementEnd
