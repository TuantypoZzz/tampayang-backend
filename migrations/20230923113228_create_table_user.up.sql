CREATE TABLE IF NOT EXISTS user 
(
    user_id              INT(11)         NOT NULL        AUTO_INCREMENT,
    user_name            VARCHAR(255)    NOT NULL,
    user_email           VARCHAR(255)    NOT NULL,
    user_password        VARCHAR(255),
    user_role            VARCHAR(10)     NOT NULL        DEFAULT 'user',
    created_date         DATETIME        NOT NULL,
    updated_date         DATETIME,
    PRIMARY KEY (user_id)
)ENGINE = InnoDB;