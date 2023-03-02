create table home
(
    id             bigint primary key,
    name           varchar(255),
    user_id        bigint references users (id),
    background_url varchar(255),
    created_at     timestamptz default now(),
    updated_at     timestamptz default now()
);

create unique index home_name_user_id_unique on home (name, user_id);
