-- Step 1: Create a temporary table to store rows to be deleted
CREATE TEMPORARY TABLE temp_users AS
SELECT user_id
FROM user
WHERE user_name = 'Admin' AND user_email = 'admin@it.com';

-- Step 2: Delete rows from the main table using the temporary table
DELETE FROM user WHERE user_id IN (SELECT user_id FROM temp_users);

-- Step 3: Drop the temporary table
DROP TEMPORARY TABLE temp_users;