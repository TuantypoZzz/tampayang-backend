CREATE TABLE IF NOT EXISTS user 
(
    id              INT(11)         NOT NULL        AUTO_INCREMENT,
    name            VARCHAR(255)    NOT NULL,
    age             INT(11)         NOT NULL        DEFAULT 0,
    created_date    DATETIME        NOT NULL,
    updated_date    DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;