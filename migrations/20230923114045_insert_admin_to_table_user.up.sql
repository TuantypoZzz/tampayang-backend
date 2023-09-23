INSERT INTO user (user_name, user_email, created_date)
SELECT 'Admin', 'admin@it.com', NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM user WHERE user_name = 'Admin' OR user_email = 'admin@it.com'
);