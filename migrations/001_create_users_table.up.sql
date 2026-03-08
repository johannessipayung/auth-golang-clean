CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    username text unique,
    email text unique,
    password text,
    role text
);