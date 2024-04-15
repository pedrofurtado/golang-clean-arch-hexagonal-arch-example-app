-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
  identifier INTEGER,
  full_name VARCHAR(255),
  state_name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
