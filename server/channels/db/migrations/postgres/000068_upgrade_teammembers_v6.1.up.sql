DO $$
DECLARE 
    column_exist boolean := false;
BEGIN
SELECT count(*) != 0 INTO column_exist
    FROM information_schema.columns
    WHERE table_name = 'teammembers'
    AND table_schema = current_schema()
    AND column_name = 'roles'
    AND NOT data_type = 'varchar(256)';
IF column_exist THEN
    ALTER TABLE teammembers ALTER COLUMN roles TYPE varchar(256);
END IF;
END $$;
