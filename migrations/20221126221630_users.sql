-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    user_id uuid PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    company_name TEXT NOT NULL,
    role TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
