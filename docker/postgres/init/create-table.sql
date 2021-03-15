create table if not exists users (
    id bigint primary key,
    name varchar(20) not null unique,
    balance bigint not null default 0,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

/* 
create table if not exists moneys (
    id bigint primary key,
    balance bigint default 0,
    user_id bigint not null unique,
    history_id varchar(20),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
*/