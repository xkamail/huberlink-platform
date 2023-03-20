-- change device_ir_remote_virtual kind column from smallint to text

ALTER TABLE device_ir_remote_virtual_keys
    ALTER COLUMN kind TYPE text;