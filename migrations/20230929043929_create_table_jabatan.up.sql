CREATE TABLE IF NOT EXISTS jabatan 
(
    jabatan_id      INT(11)         NOT NULL AUTO_INCREMENT,
    jabatan_name    VARCHAR(100)    NOT NULL,
    jabatan_status  ENUM('active', 'inactive') DEFAULT 'inactive',
    created_date    DATETIME,
    updated_date    DATETIME,
    PRIMARY KEY (jabatan_id)
) ENGINE = InnoDB;
