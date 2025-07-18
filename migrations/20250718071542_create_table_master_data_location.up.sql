CREATE TABLE IF NOT EXISTS provinces (
    province_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    province_name VARCHAR(100) NOT NULL,
    province_code VARCHAR(10) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    PRIMARY KEY (province_id)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS regencies (
    regency_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    province_id VARCHAR(36) NOT NULL,
    regency_name VARCHAR(100) NOT NULL,
    regency_code VARCHAR(10) NOT NULL,
    regency_type ENUM('kabupaten', 'kota') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    PRIMARY KEY (regency_id),
    CONSTRAINT fk_regencies_provinces FOREIGN KEY (province_id) REFERENCES provinces(province_id)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS districts (
    district_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    regency_id VARCHAR(36) NOT NULL,
    district_name VARCHAR(100) NOT NULL,
    district_code VARCHAR(10) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    PRIMARY KEY (district_id),
    CONSTRAINT fk_districts_regencies FOREIGN KEY (regency_id) REFERENCES regencies(regency_id)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS villages (
    village_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    district_id VARCHAR(36) NOT NULL,
    village_name VARCHAR(100) NOT NULL,
    village_code VARCHAR(10) NOT NULL,
    village_type ENUM('desa', 'kelurahan') NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    PRIMARY KEY (village_id),
    CONSTRAINT fk_villages_districts FOREIGN KEY (district_id) REFERENCES districts(district_id)
) ENGINE = InnoDB;