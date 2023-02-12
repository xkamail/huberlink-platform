create table users
(
    id         bigint primary key,
    name       varchar(255),
    email      varchar(255),
    password   varchar(255),
    created_at timestamptz,
    updated_at timestamptz
);

create table home
(
    id             bigint primary key,
    name           varchar(255),
    user_id        bigint references users (id),
    background_url varchar(255),
    created_at     timestamptz,
    updated_at     timestamptz
);


create table home_members
(
    id         bigint primary key,
    home_id    bigint references home (id),
    user_id    bigint references users (id),
    permission bigint default 0,
    created_at timestamptz,
    updated_at timestamptz
);

create table rooms
(
    id         bigint primary key,
    name       varchar(255),
    home_id    bigint references home (id),
    created_at timestamptz,
    updated_at timestamptz
);

create table devices
(
    id                  bigint primary key,
    name                varchar(255),
    home_id             bigint references home (id)  default null,
    room_id             bigint references rooms (id) default null,
    device_kind         smallint,
    latest_heartbeat_at timestamptz,
    created_at          timestamptz,
    updated_at          timestamptz
);

