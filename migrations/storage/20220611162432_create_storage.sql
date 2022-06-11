-- +goose Up
-- +goose StatementBegin
CREATE TABLE storage(
    order_id int,
    product_id int,
    isReserved boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE storage;
-- +goose StatementEnd
