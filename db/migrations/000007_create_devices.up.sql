create table devices
(
    id                  bigint primary key,
    name                text not null,
    icon                text not null,
    model               text        default '',
    kind                smallint,
    home_id             bigint references home (id),
    user_id             bigint references users (id),
    token               text not null,
    ip_address          text,
    location            text,
    latest_heartbeat_at timestamptz default null,
    created_at          timestamptz default now(),
    updated_at          timestamptz default now()
);

create unique index devices_home_id_name_unique on devices (home_id, name);