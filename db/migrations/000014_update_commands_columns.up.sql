alter table device_ir_remote_commands
    add column remark text default null;
alter table device_ir_remote_commands
    drop column frequency;