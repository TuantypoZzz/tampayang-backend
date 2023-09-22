CREATE TABLE IF NOT EXISTS example 
(
    id              INT(11)         NOT NULL        AUTO_INCREMENT,
    name            VARCHAR(255)    NOT NULL,
    created         DATE            NOT NULL,
    rating          DOUBLE          NOT NULL        DEFAULT 0.0,
    booleandesu     TINYINT(1)      NOT NULL        DEFAULT 0,
    created_date    DATETIME        NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;