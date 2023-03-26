create table home_scenes
(
    id              bigint primary key,
    home_id         bigint references home (id) on delete cascade,
    name            text not null,
    run             smallint    default 1,
    schedule_repeat smallint    default 0,
    schedule_time   varchar(5)  default '00:00',
    created_at      timestamptz default now(),
    updated_at      timestamptz default now()
);

create table home_scenes_actions
(
    id        bigint primary key,
    scene_id  bigint references home_scenes (id) on delete cascade,
    device_id bigint references devices (id) on delete cascade,
    action    text not null
);