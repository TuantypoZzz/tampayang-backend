CREATE TABLE IF NOT EXISTS pangkat 
(
    pangkat_id          INT(11)         NOT NULL        AUTO_INCREMENT,
    pangkat_name        VARCHAR(100)    NOT NULL,
    created_date         DATETIME,
    updated_date         DATETIME,
    PRIMARY KEY (pangkat_id)
)ENGINE = InnoDB;