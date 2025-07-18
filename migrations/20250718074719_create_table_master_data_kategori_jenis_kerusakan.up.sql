CREATE TABLE IF NOT EXISTS infrastructure_categories (
    infrastructure_category_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE,
    description TEXT,
    icon VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (infrastructure_category_id)
) ENGINE = InnoDB;

-- Tabel Jenis Kerusakan
CREATE TABLE IF NOT EXISTS damage_types (
    damage_type_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    infrastructure_category_id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (damage_type_id),
    CONSTRAINT fk_damage_types_infrastructure_categories FOREIGN KEY (infrastructure_category_id) REFERENCES infrastructure_categories(infrastructure_category_id),
    INDEX idx_infrastructure_category_id (infrastructure_category_id)
) ENGINE = InnoDB;
