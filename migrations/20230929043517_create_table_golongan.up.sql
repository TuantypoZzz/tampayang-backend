CREATE TABLE IF NOT EXISTS golongan 
(
    golongan_id          INT(11)         NOT NULL        AUTO_INCREMENT,
    golongan_name        VARCHAR(100)    NOT NULL,
    created_date         DATETIME        NOT NULL,
    updated_date         DATETIME,
    PRIMARY KEY (golongan_id)
)ENGINE = InnoDB;