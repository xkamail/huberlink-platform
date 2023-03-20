-- change column kind to smallint
alter table device_ir_remote_virtual_keys
    drop column kind,
    add column kind smallint;
