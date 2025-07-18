-- Delete Jenis Kerusakan untuk Pengaman Sungai
DELETE FROM damage_types WHERE infrastructure_category_id = (SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'SNG');

-- Delete Jenis Kerusakan untuk Pengaman Pantai
DELETE FROM damage_types WHERE infrastructure_category_id = (SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'PNT');

-- Delete Jenis Kerusakan untuk Irigasi
DELETE FROM damage_types WHERE infrastructure_category_id = (SELECT infrastructure_category_id FROM infrastructure_categories WHERE code = 'IRG');

-- Delete Kategori Infrastruktur
DELETE FROM infrastructure_categories WHERE code IN ('IRG', 'PNT', 'SNG');
