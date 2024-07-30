-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL Unique,
    email text NOT NULL Unique,
    password VARCHAR(255) NOT NULL,
    description text NOT NULL,
    avatar_id VARCHAR(255) NOT NULL,
    role TEXT[] NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_login_idx ON users (login);
CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS users_login_idx;
DROP INDEX IF EXISTS users_email_idx;
-- +goose StatementEnd
