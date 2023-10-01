CREATE TABLE IF NOT EXISTS seksi 
(
    seksi_id             INT(11)         NOT NULL        AUTO_INCREMENT,
    seksi_name           VARCHAR(100)    NOT NULL,
    created_date         DATETIME,
    updated_date         DATETIME,
    PRIMARY KEY (seksi_id)
)ENGINE = InnoDB;