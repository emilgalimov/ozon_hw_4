-- +goose Up
-- +goose StatementBegin

CREATE TABLE storage
(
    order_id    int NOT NULL,
    product_id  int,
    is_reserved boolean
)
PARTITION BY HASH (order_id);

CREATE FOREIGN TABLE storage_even PARTITION OF storage
    FOR VALUES WITH (MODULUS 2,REMAINDER 0)
    SERVER storage_even;

CREATE FOREIGN TABLE storage_odd PARTITION OF storage
    FOR VALUES WITH (MODULUS 2,REMAINDER 1)
    SERVER storage_odd;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FOREIGN TABLE storage_odd;
DROP FOREIGN TABLE storage_even;
DROP TABLE storage;
-- +goose StatementEnd
