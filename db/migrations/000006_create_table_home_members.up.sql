create table home_members
(
    id         bigint primary key,
    home_id    bigint references home (id),
    user_id    bigint references users (id),
    permission bigint      default 0,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
create unique index home_members_home_id_user_id_unique on home_members (home_id, user_id);