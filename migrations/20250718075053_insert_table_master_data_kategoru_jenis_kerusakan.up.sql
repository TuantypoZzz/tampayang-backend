-- Insert Kategori Infrastruktur
INSERT INTO infrastructure_categories (name, code, description, icon) VALUES 
('Irigasi', 'IRG', 'Infrastruktur irigasi dan pengairan', 'fas fa-tint'),
('Pengaman Pantai', 'PNT', 'Infrastruktur pengaman pantai dan pesisir', 'fas fa-umbrella-beach'),
('Pengaman Sungai', 'SNG', 'Infrastruktur pengaman sungai dan banjir', 'fas fa-water');

-- Insert Jenis Kerusakan untuk Irigasi
INSERT INTO damage_types (category_id, name, code, description) VALUES 
(1, 'Saluran Tersumbat', 'STS', 'Saluran irigasi tersumbat sampah atau lumpur'),
(1, 'Saluran Bocor', 'SBC', 'Kebocoran pada saluran irigasi'),
(1, 'Pintu Air Rusak', 'PAR', 'Kerusakan pada pintu air irigasi'),
(1, 'Tanggul Jebol', 'TJB', 'Tanggul irigasi jebol atau rusak');

-- Insert Jenis Kerusakan untuk Pengaman Pantai
INSERT INTO damage_types (category_id, name, code, description) VALUES 
(2, 'Pemecah Gelombang Rusak', 'PGR', 'Kerusakan pada pemecah gelombang'),
(2, 'Tanggul Laut Jebol', 'TLJ', 'Tanggul pengaman pantai jebol'),
(2, 'Abrasi Pantai', 'APS', 'Abrasi atau pengikisan pantai'),
(2, 'Dermaga Rusak', 'DMR', 'Kerusakan pada dermaga atau pelabuhan kecil');

-- Insert Jenis Kerusakan untuk Pengaman Sungai
INSERT INTO damage_types (category_id, name, code, description) VALUES 
(3, 'Tanggul Sungai Jebol', 'TSJ', 'Tanggul pengaman sungai jebol'),
(3, 'Sungai Pendangkalan', 'SPD', 'Pendangkalan sungai akibat sedimentasi'),
(3, 'Jembatan Rusak', 'JMR', 'Kerusakan pada jembatan'),
(3, 'Bronjong Rusak', 'BRR', 'Kerusakan pada bronjong pengaman sungai');-- Update created_at for infrastructure_categories
UPDATE infrastructure_categories SET created_at = NOW();

-- Update created_at for damage_types
UPDATE damage_types SET created_at = NOW();
