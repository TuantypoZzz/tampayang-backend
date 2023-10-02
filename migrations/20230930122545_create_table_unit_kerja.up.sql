CREATE TABLE IF NOT EXISTS unit_kerja 
(
    uk_id                INT(11)         NOT NULL        AUTO_INCREMENT,
    uk_name              VARCHAR(100)    NOT NULL,
    created_date         DATETIME,
    updated_date         DATETIME,
    PRIMARY KEY (uk_id)
)ENGINE = InnoDB;