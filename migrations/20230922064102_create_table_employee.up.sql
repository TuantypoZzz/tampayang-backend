CREATE TABLE IF NOT EXISTS employee
(
    id              INT(11)         NOT NULL        AUTO_INCREMENT,
    name            VARCHAR(255)    NOT NULL,
    nip             VARCHAR(150)    NOT NULL,
    bidang          VARCHAR(100)    NOT NULL,
    seksi           VARCHAR(100)    NOT NULL,
    unit_kerja      VARCHAR(100)    NOT NULL,
    gender          int(11)         NOT NULL,
    birth_place     VARCHAR(255)    NOT NULL,
    birth_date      DATETIME        NOT NULL,
    phone           VARCHAR(100)    NOT NULL,
    email           VARCHAR(100)    NOT NULL        DEFAULT 0,
    created_date    DATETIME        NOT NULL,
    updated_date    DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;