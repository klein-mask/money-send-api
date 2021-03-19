create table if not exists users (
    id bigint primary key,
    name varchar(20) not null unique,
    password varchar(20) not null,
    balance bigint default 0,
    is_balance_receivable boolean default false,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
