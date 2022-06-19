-- +goose Up
-- +goose StatementBegin
CREATE TABLE storage_even(
    order_id int NOT NULL
        CONSTRAINT order_id CHECK (mod(order_id, 2) = 0),
    product_id int,
    is_reserved boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE storage_even;
-- +goose StatementEnd
