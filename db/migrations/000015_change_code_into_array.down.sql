alter table device_ir_remote_commands
    drop column code,
    add column code text not null;