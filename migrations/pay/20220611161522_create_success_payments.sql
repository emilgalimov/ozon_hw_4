-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments
(
    order_id   int,
    sum        real,
    created_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
