create table users
(
    id         bigint primary key,
    username   text   not null,
    email      text   not null,
    password   text        default '',
    discord_id bigint not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create unique index users_email_unique on users (email);