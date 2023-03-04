create table device_ir_remotes
(
    id         bigint primary key,
    device_id  bigint references devices (id),
    home_id    bigint references home (id) default null,
    created_at timestamptz                 default now(),
    updated_at timestamptz                 default now()
);

create table device_ir_remote_virtual_keys
(
    id         bigint primary key,
    remote_id  bigint references device_ir_remotes (id) on delete cascade,
    name       text not null,
    kind       text not null,
    icon       text not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table device_ir_remote_commands
(
    id         bigint primary key, -- command id
    remote_id  bigint references device_ir_remotes (id) on delete cascade,
    virtual_id bigint references device_ir_remote_virtual_keys (id) on delete cascade,
    name       text    not null,
    code       text    not null,
    frequency  integer not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);