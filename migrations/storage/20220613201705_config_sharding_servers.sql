-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION postgres_fdw;

CREATE SERVER storage_even
FOREIGN DATA WRAPPER postgres_fdw
OPTIONS (host 'postgres-1_even', port '5432', dbname 'storage_even');

CREATE USER MAPPING FOR user
SERVER storage_even
OPTIONS (user 'user', password 'pass');

CREATE SERVER storage_odd
    FOREIGN DATA WRAPPER postgres_fdw
    OPTIONS (host 'postgres-1_odd', port '5432', dbname 'storage_odd');

CREATE USER MAPPING FOR user
    SERVER storage_odd
    OPTIONS (user 'user', password 'pass');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP USER MAPPING FOR postgres SERVER storage_even;
DROP SERVER storage_even;
DROP USER MAPPING FOR postgres SERVER storage_odd;
DROP SERVER storage_odd;
-- +goose StatementEnd
