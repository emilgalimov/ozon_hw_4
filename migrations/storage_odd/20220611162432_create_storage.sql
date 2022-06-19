-- +goose Up
-- +goose StatementBegin
CREATE TABLE storage_odd(
    order_id int NOT NULL
        CONSTRAINT order_id CHECK (mod(order_id, 2) = 1),
    product_id int,
    is_reserved boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE storage;
-- +goose StatementEnd
