-- Tabel Riwayat Status Laporan
CREATE TABLE IF NOT EXISTS report_status_history (
    report_status_history_id VARCHAR(36) NOT NULL DEFAULT (UUID()),
    report_id VARCHAR(36) NOT NULL,
    previous_status ENUM('baru', 'verifikasi', 'proses', 'selesai', 'ditolak') NULL,
    new_status ENUM('baru', 'verifikasi', 'proses', 'selesai', 'ditolak') NOT NULL,
    notes TEXT NULL,
    updated_by INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (report_status_history_id),
    FOREIGN KEY (report_id) REFERENCES reports(report_id) ON DELETE CASCADE,
    FOREIGN KEY (updated_by) REFERENCES users(user_id) ON DELETE SET NULL,
    INDEX idx_report_id (report_id),
    INDEX idx_created_at (created_at)
);