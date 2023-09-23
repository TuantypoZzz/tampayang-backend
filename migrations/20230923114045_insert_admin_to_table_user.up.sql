INSERT INTO user (user_name, user_email, user_password, user_role, created_date)
SELECT 'administrator', 'admin@it.com', "$2a$10$tbrwfPqSZl3Y1uAqlqQx0eXkAeekk/WvEs2oJQMwjBWb/dbzs1FWe", "admin", NOW()
WHERE NOT EXISTS (
    SELECT 1 FROM user WHERE user_name = 'Admin' OR user_email = 'admin@it.com'
);