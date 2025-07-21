-- Migration Down: Mengembalikan nilai latitude dan longitude menjadi NULL

START TRANSACTION;

-- Mengatur NULL untuk tabel provinces
UPDATE provinces
SET latitude = NULL, longitude = NULL
WHERE province_name = 'Maluku';

-- Mengatur NULL untuk tabel regencies
-- Ini akan mengatur ulang semua kabupaten/kota di provinsi Maluku
UPDATE regencies
SET latitude = NULL, longitude = NULL
WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku');

-- Mengatur NULL untuk tabel districts
-- Ini akan mengatur ulang semua kecamatan yang datanya telah diisi
UPDATE districts
SET latitude = NULL, longitude = NULL
WHERE regency_id IN (SELECT regency_id FROM regencies WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'));

-- Mengatur NULL untuk tabel villages
-- Ini akan mengatur ulang semua desa/kelurahan yang datanya telah diisi
UPDATE villages
SET latitude = NULL, longitude = NULL
WHERE district_id IN (
    SELECT district_id FROM districts WHERE regency_id IN (
        SELECT regency_id FROM regencies WHERE province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')
    )
);

COMMIT;
