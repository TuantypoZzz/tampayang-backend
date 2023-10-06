CREATE TABLE IF NOT EXISTS bidang 
(
    bidang_id            INT(11)         NOT NULL        AUTO_INCREMENT,
    bidang_name          VARCHAR(100)    NOT NULL,
    created_date         DATETIME,
    updated_date         DATETIME,
    PRIMARY KEY (bidang_id)
)ENGINE = InnoDB;