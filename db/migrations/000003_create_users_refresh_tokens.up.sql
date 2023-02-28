create table users_refresh_tokens
(
    id         bigint primary key,
    user_id    bigint      not null,
    token      varchar(255),
    expired_at timestamptz not null,
    issued_at  timestamptz not null,
    created_at timestamptz default now(),
    foreign key (user_id) references users (id)
);
