-- Tabel Laporan
CREATE TABLE IF NOT EXISTS reports (
    report_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    report_number VARCHAR(20) NOT NULL UNIQUE,
    reporter_name VARCHAR(100) NOT NULL,
    reporter_phone VARCHAR(20) NOT NULL,
    reporter_email VARCHAR(100) NULL,
    
    -- Kategori & Jenis
    infrastructure_category_id VARCHAR(36) NOT NULL,
    damage_type_id VARCHAR(36) NOT NULL,
    
    -- Lokasi
    province_id VARCHAR(36) NOT NULL,
    regency_id VARCHAR(36) NOT NULL,
    district_id VARCHAR(36) NOT NULL,
    village_id VARCHAR(36) NOT NULL,
    location_detail TEXT,
    
    -- Koordinat GPS
    latitude DECIMAL(10, 8) NULL,
    longitude DECIMAL(11, 8) NULL,
    
    -- Deskripsi & Urgensi
    description TEXT NOT NULL,
    urgency_level ENUM('rendah', 'sedang', 'tinggi') NOT NULL,
    
    -- Status & Workflow
    status ENUM('baru', 'verifikasi', 'proses', 'selesai', 'ditolak') DEFAULT 'baru',
    assigned_to INT NULL,
    pic_name VARCHAR(100) NULL,
    pic_contact VARCHAR(100) NULL,
    
    -- Catatan Admin
    admin_notes TEXT NULL,
    completion_notes TEXT NULL,
    
    -- Estimasi & Tanggal
    estimated_completion DATE NULL,
    completed_at TIMESTAMP NULL,
    
    -- Metadata
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (report_id),

    -- Foreign Keys
    FOREIGN KEY (infrastructure_category_id) REFERENCES infrastructure_categories(infrastructure_category_id),
    FOREIGN KEY (damage_type_id) REFERENCES damage_types(damage_type_id),
    FOREIGN KEY (province_id) REFERENCES provinces(province_id),
    FOREIGN KEY (regency_id) REFERENCES regencies(regency_id),
    FOREIGN KEY (district_id) REFERENCES districts(district_id),
    FOREIGN KEY (village_id) REFERENCES villages(village_id),
    FOREIGN KEY (assigned_to) REFERENCES users(user_id) ON DELETE SET NULL,
    
    -- Indexes
    INDEX idx_report_number (report_number),
    INDEX idx_status (status),
    INDEX idx_infrastructure_category_id (infrastructure_category_id),
    INDEX idx_regency_id (regency_id),
    INDEX idx_created_at (created_at),
    INDEX idx_urgency_level (urgency_level),
    INDEX idx_assigned_to (assigned_to)
);