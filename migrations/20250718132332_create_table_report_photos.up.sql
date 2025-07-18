-- Tabel Foto Laporan
CREATE TABLE IF NOT EXISTS report_photos (
    report_photo_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    report_id VARCHAR(36) NOT NULL,
    filename VARCHAR(255) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_size INT NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    description TEXT NULL,
    is_main BOOLEAN DEFAULT FALSE,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (report_photo_id),
    FOREIGN KEY (report_id) REFERENCES reports(id) ON DELETE CASCADE,
    INDEX idx_report_id (report_id)
);
