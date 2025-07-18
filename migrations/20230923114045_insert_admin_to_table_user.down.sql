-- Step 1: Create a temporary table to store rows to be deleted
CREATE TEMPORARY TABLE temp_users AS
SELECT user_id
FROM users
WHERE user_name = 'Admin' AND user_email = 'admin@it.com';

-- Step 2: Delete rows from the main table using the temporary table
DELETE FROM users WHERE user_id IN (SELECT user_id FROM temp_users);

-- Step 3: Drop the temporary table
DROP TEMPORARY TABLE temp_users;