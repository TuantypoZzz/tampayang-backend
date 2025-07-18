-- Delete Jenis Kerusakan untuk Pengaman Sungai
DELETE FROM damage_types WHERE category_id = 3;

-- Delete Jenis Kerusakan untuk Pengaman Pantai
DELETE FROM damage_types WHERE category_id = 2;

-- Delete Jenis Kerusakan untuk Irigasi
DELETE FROM damage_types WHERE category_id = 1;

-- Delete Kategori Infrastruktur
DELETE FROM infrastructure_categories WHERE code IN ('IRG', 'PNT', 'SNG');
