-- Insert Kategori Infrastruktur
INSERT INTO infrastructure_categories (name, code, description, icon, created_at) VALUES 
('Irigasi', 'IRG', 'Infrastruktur irigasi dan pengairan', 'fas fa-tint', NOW()),
('Pengaman Pantai', 'PNT', 'Infrastruktur pengaman pantai dan pesisir', 'fas fa-umbrella-beach', NOW()),
('Pengaman Sungai', 'SNG', 'Infrastruktur pengaman sungai dan banjir', 'fas fa-water', NOW());

-- Insert Jenis Kerusakan untuk Irigasi
INSERT INTO damage_types (infrastructure_category_id, name, code, description, created_at) VALUES 
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'IRG'), 'Saluran Tersumbat', 'STS', 'Saluran irigasi tersumbat sampah atau lumpur',NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'IRG'), 'Saluran Bocor', 'SBC', 'Kebocoran pada saluran irigasi', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'IRG'), 'Pintu Air Rusak', 'PAR', 'Kerusakan pada pintu air irigasi', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'IRG'), 'Tanggul Jebol', 'TJB', 'Tanggul irigasi jebol atau rusak', NOW());

-- Insert Jenis Kerusakan untuk Pengaman Pantai
INSERT INTO damage_types (infrastructure_category_id, name, code, description, created_at) VALUES 
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'PNT'), 'Pemecah Gelombang Rusak', 'PGR', 'Kerusakan pada pemecah gelombang', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'PNT'), 'Tanggul Laut Jebol', 'TLJ', 'Tanggul pengaman pantai jebol', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'PNT'), 'Abrasi Pantai', 'APS', 'Abrasi atau pengikisan pantai', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'PNT'), 'Dermaga Rusak', 'DMR', 'Kerusakan pada dermaga atau pelabuhan kecil', NOW());

-- Insert Jenis Kerusakan untuk Pengaman Sungai
INSERT INTO damage_types (infrastructure_category_id, name, code, description, created_at) VALUES 
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'SNG'), 'Tanggul Sungai Jebol', 'TSJ', 'Tanggul pengaman sungai jebol', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'SNG'), 'Sungai Pendangkalan', 'SPD', 'Pendangkalan sungai akibat sedimentasi', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'SNG'), 'Jembatan Rusak', 'JMR', 'Kerusakan pada jembatan', NOW()),
((SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'SNG'), 'Bronjong Rusak', 'BRR', 'Kerusakan pada bronjong pengaman sungai', NOW());
