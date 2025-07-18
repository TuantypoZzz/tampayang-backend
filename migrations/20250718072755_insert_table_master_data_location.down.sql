START TRANSACTION;

-- Hapus Desa/Kelurahan yang terkait dengan Provinsi Maluku
DELETE FROM villages WHERE district_id IN (
    SELECT district_id FROM districts WHERE regency_id IN (
        SELECT regency_id FROM regencies WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')
    )
);

-- Hapus Kecamatan (Districts) yang terkait dengan Provinsi Maluku
DELETE FROM districts WHERE regency_id IN (
    SELECT regency_id FROM regencies WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')
);

-- Hapus Kabupaten/Kota yang terkait dengan Provinsi Maluku
DELETE FROM regencies WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku');

-- Hapus Provinsi Maluku
DELETE FROM provinces WHERE province_name = 'Maluku';

COMMIT;