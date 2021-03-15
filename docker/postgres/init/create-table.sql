create table users (
    id bigint primary key,
    name varchar(20) not null,
    money_id bigint unique,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table moneys (
    id bigint primary key,
    name varchar(20) not null,
    history_id varchar(20),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
