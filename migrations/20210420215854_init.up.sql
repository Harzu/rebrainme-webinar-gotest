CREATE TABLE IF NOT EXISTS clients (
    email      text primary key,
    full_name  text not null,
    city       text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone
);