-- change column kind to smallint
alter table device_ir_remote_virtual_keys
    alter column kind set data type text;
