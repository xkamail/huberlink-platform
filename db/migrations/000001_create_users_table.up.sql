create table users
(
    id         bigint primary key,
    username   varchar(255),
    email      varchar(255),
    password   varchar(255),
    discord_id bigint not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create unique index users_email_unique on users (email);