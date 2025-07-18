-- Delete Desa/Kelurahan
DELETE FROM villages WHERE districts_id IN (1, 2, 7, 8, 21, 47, 60);

-- Delete Kecamatan (Districts)
DELETE FROM districts WHERE regencies_id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11);

-- Delete Kabupaten/Kota
DELETE FROM regencies WHERE province_id = 1;

-- Delete Provinsi Maluku
DELETE FROM provinces WHERE province_code = 'MAL';
