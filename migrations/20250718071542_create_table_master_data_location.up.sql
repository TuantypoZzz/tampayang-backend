CREATE TABLE IF NOT EXISTS provinces 
(
    provinces_id         INT(11)         NOT NULL        AUTO_INCREMENT,
    province_name        VARCHAR(100)    NOT NULL,
    province_code        VARCHAR(10)    NOT NULL,
    created_at           DATETIME        NOT NULL,
    updated_at           DATETIME,
    PRIMARY KEY (provinces_id)
)ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS regencies 
(
    regencies_id         INT(11)         NOT NULL        AUTO_INCREMENT,
    province_id          INT(11)    NOT NULL,
    regencies_name       VARCHAR(100)    NOT NULL,
    regencies_code       VARCHAR(10)    NOT NULL,
    regencies_type       ENUM('kabupaten', 'kota') NOT NULL,
    created_at           DATETIME        NOT NULL,
    updated_at           DATETIME,
    PRIMARY KEY (regencies_id)
)ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS districts 
(
    districts_id         INT(11)         NOT NULL        AUTO_INCREMENT,
    regencies_id         INT(11)    NOT NULL,
    districts_name       VARCHAR(100)    NOT NULL,
    districts_code       VARCHAR(10)    NOT NULL,
    created_at           DATETIME        NOT NULL,
    updated_at           DATETIME,
    PRIMARY KEY (districts_id)
)ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS villages 
(
    villages_id         INT(11)         NOT NULL        AUTO_INCREMENT,
    districts_id          INT(11)    NOT NULL,
    villages_name       VARCHAR(100)    NOT NULL,
    villages_code       VARCHAR(10)    NOT NULL,
    villages_type       ENUM('desa', 'kelurahan') NOT NULL,
    created_at           DATETIME        NOT NULL,
    updated_at           DATETIME,
    PRIMARY KEY (villages_id)
)ENGINE = InnoDB;ALTER TABLE regencies
ADD CONSTRAINT fk_regencies_provinces
FOREIGN KEY (province_id) REFERENCES provinces(provinces_id);

ALTER TABLE districts
ADD CONSTRAINT fk_districts_regencies
FOREIGN KEY (regencies_id) REFERENCES regencies(regencies_id);

ALTER TABLE villages
ADD CONSTRAINT fk_villages_districts
FOREIGN KEY (districts_id) REFERENCES districts(districts_id);
