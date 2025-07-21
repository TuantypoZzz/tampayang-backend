-- Migration to insert latitude and longitude data

START TRANSACTION;

-- Update Provinsi Maluku
UPDATE provinces
SET latitude = -3.695430, longitude = 128.179330
WHERE province_name = 'Maluku';

-- Update Kabupaten/Kota di Maluku
UPDATE regencies SET latitude = -3.6917, longitude = 128.1833 WHERE regency_name = 'Kota Ambon';
UPDATE regencies SET latitude = -3.2000, longitude = 129.2500 WHERE regency_name = 'Kabupaten Maluku Tengah';
UPDATE regencies SET latitude = -3.3333, longitude = 128.5000 WHERE regency_name = 'Kabupaten Seram Bagian Barat';
UPDATE regencies SET latitude = -3.5000, longitude = 130.0000 WHERE regency_name = 'Kabupaten Seram Bagian Timur';
UPDATE regencies SET latitude = -6.0000, longitude = 134.5000 WHERE regency_name = 'Kabupaten Kepulauan Aru';
UPDATE regencies SET latitude = -7.8333, longitude = 126.5000 WHERE regency_name = 'Kabupaten Maluku Barat Daya';
UPDATE regencies SET latitude = -3.3333, longitude = 127.0000 WHERE regency_name = 'Kabupaten Buru';
UPDATE regencies SET latitude = -3.8333, longitude = 126.6667 WHERE regency_name = 'Kabupaten Buru Selatan';
UPDATE regencies SET latitude = -7.9700, longitude = 131.2990 WHERE regency_name = 'Kabupaten Kepulauan Tanimbar';
UPDATE regencies SET latitude = -5.6572, longitude = 132.7321 WHERE regency_name = 'Kota Tual';
UPDATE regencies SET latitude = -5.6500, longitude = 132.7500 WHERE regency_name = 'Kabupaten Maluku Tenggara';
UPDATE regencies SET latitude = -7.6667, longitude = 131.5000 WHERE regency_name = 'Kabupaten Maluku Tenggara Barat';

-- Update Kecamatan (Districts)
-- Kota Ambon
UPDATE districts SET latitude = -3.7073, longitude = 128.1602 WHERE district_name = 'Nusaniwe';
UPDATE districts SET latitude = -3.6950, longitude = 128.1794 WHERE district_name = 'Sirimau';
UPDATE districts SET latitude = -3.6667, longitude = 128.1667 WHERE district_name = 'Teluk Ambon';
UPDATE districts SET latitude = -3.6333, longitude = 128.2333 WHERE district_name = 'Baguala';
UPDATE districts SET latitude = -3.6500, longitude = 128.2167 WHERE district_name = 'Teluk Ambon Baguala';
UPDATE districts SET latitude = -3.7333, longitude = 128.2167 WHERE district_name = 'Leitimur Selatan';

-- Kabupaten Kepulauan Tanimbar
UPDATE districts SET latitude = -7.9700, longitude = 131.2990 WHERE district_name = 'Saumlaki';
UPDATE districts SET latitude = -7.9833, longitude = 131.3000 WHERE district_name = 'Tanimbar Selatan';
UPDATE districts SET latitude = -7.4167, longitude = 131.5000 WHERE district_name = 'Tanimbar Utara';
UPDATE districts SET latitude = -8.2000, longitude = 130.8333 WHERE district_name = 'Selaru';
UPDATE districts SET latitude = -7.6667, longitude = 131.1667 WHERE district_name = 'Wer Tamrian';
UPDATE districts SET latitude = -7.5833, longitude = 131.5000 WHERE district_name = 'Wer Maktian';
UPDATE districts SET latitude = -7.5000, longitude = 131.8333 WHERE district_name = 'Nirunmas';
UPDATE districts SET latitude = -7.1667, longitude = 131.5833 WHERE district_name = 'Molu Maru';
UPDATE districts SET latitude = -7.3333, longitude = 131.0000 WHERE district_name = 'Wuar Labobar';
UPDATE districts SET latitude = -7.6667, longitude = 131.6667 WHERE district_name = 'Kormomolin';
UPDATE districts SET latitude = -7.0833, longitude = 131.7500 WHERE district_name = 'Nuswotar';
UPDATE districts SET latitude = -7.6667, longitude = 131.1667 WHERE district_name = 'Wertamrian';

-- Kabupaten Maluku Tengah
UPDATE districts SET latitude = -3.5833, longitude = 128.3167 WHERE district_name = 'Salahutu';
UPDATE districts SET latitude = -3.5000, longitude = 128.0833 WHERE district_name = 'Leihitu';
UPDATE districts SET latitude = -3.5500, longitude = 127.9000 WHERE district_name = 'Leihitu Barat';
UPDATE districts SET latitude = -3.3167, longitude = 128.9167 WHERE district_name = 'Kota Masohi';
UPDATE districts SET latitude = -3.3333, longitude = 128.9167 WHERE district_name = 'Amahai';
UPDATE districts SET latitude = -6.5000, longitude = 129.5000 WHERE district_name = 'Teon Nila Serua';
UPDATE districts SET latitude = -3.5000, longitude = 129.5000 WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');
UPDATE districts SET latitude = -3.3167, longitude = 127.0833 WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');
UPDATE districts SET latitude = -3.1000, longitude = 130.4833 WHERE district_name = 'Bula' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');
UPDATE districts SET latitude = -3.2500, longitude = 126.8333 WHERE district_name = 'Waplau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');
UPDATE districts SET latitude = -4.5167, longitude = 129.9000 WHERE district_name = 'Kepulauan Banda';
UPDATE districts SET latitude = -4.5333, longitude = 129.9000 WHERE district_name = 'Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');
UPDATE districts SET latitude = -3.6667, longitude = 128.7833 WHERE district_name = 'Nusalaut';
UPDATE districts SET latitude = -3.5667, longitude = 128.6500 WHERE district_name = 'Saparua';
UPDATE districts SET latitude = -3.5667, longitude = 128.4833 WHERE district_name = 'Pulau Haruku';
UPDATE districts SET latitude = -3.5000, longitude = 128.0833 WHERE district_name = 'Seit Kaitetu';
UPDATE districts SET latitude = -3.4000, longitude = 129.0000 WHERE district_name = 'Teluk Dalam';
UPDATE districts SET latitude = -6.5500, longitude = 129.6000 WHERE district_name = 'TNS Timur';
UPDATE districts SET latitude = -3.4167, longitude = 128.7500 WHERE district_name = 'Elpaputih' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah');

-- Kabupaten Seram Bagian Barat
UPDATE districts SET latitude = -3.3833, longitude = 128.3333 WHERE district_name = 'Kairatu';
UPDATE districts SET latitude = -3.4000, longitude = 128.2500 WHERE district_name = 'Kairatu Barat';
UPDATE districts SET latitude = -2.9167, longitude = 129.5000 WHERE district_name = 'Seram Utara';
UPDATE districts SET latitude = -3.1667, longitude = 128.8333 WHERE district_name = 'Taniwel';
UPDATE districts SET latitude = -3.3333, longitude = 127.9167 WHERE district_name = 'Huamual';
UPDATE districts SET latitude = -3.2500, longitude = 127.5000 WHERE district_name = 'Huamual Belakang';
UPDATE districts SET latitude = -3.4500, longitude = 128.4167 WHERE district_name = 'Amalatu';
UPDATE districts SET latitude = -3.3000, longitude = 128.5000 WHERE district_name = 'Inamosol';
UPDATE districts SET latitude = -3.4167, longitude = 128.7500 WHERE district_name = 'Elpaputih' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat');
UPDATE districts SET latitude = -3.3333, longitude = 127.9167 WHERE district_name = 'Huamual Utara';
UPDATE districts SET latitude = -3.2500, longitude = 127.5833 WHERE district_name = 'Kepulauan Manipa';
UPDATE districts SET latitude = -3.2500, longitude = 128.1833 WHERE district_name = 'Piru';

-- Kabupaten Seram Bagian Timur
UPDATE districts SET latitude = -3.1000, longitude = 130.4833 WHERE district_name = 'Bula' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts SET latitude = -3.6667, longitude = 130.0000 WHERE district_name = 'Werinama';
UPDATE districts SET latitude = -4.0000, longitude = 131.5000 WHERE district_name = 'Pulau Gorom';
UPDATE districts SET latitude = -4.2500, longitude = 131.0000 WHERE district_name = 'Wakate';
UPDATE districts SET latitude = -3.5000, longitude = 130.5000 WHERE district_name = 'Tutuk Tolu';
UPDATE districts SET latitude = -3.7500, longitude = 130.2500 WHERE district_name = 'Siwalalat';
UPDATE districts SET latitude = -3.8333, longitude = 130.5000 WHERE district_name = 'Kilmury';
UPDATE districts SET latitude = -3.2500, longitude = 130.7500 WHERE district_name = 'Teluk Waru';
UPDATE districts SET latitude = -4.0500, longitude = 131.6000 WHERE district_name = 'Gorom Timur';
UPDATE districts SET latitude = -3.1500, longitude = 130.4000 WHERE district_name = 'Bula Barat';
UPDATE districts SET latitude = -3.1000, longitude = 130.4833 WHERE district_name = 'Kota Bula';
UPDATE districts SET latitude = -3.5000, longitude = 130.0000 WHERE district_name = 'Teluk Elpaputih';
UPDATE districts SET latitude = -3.5000, longitude = 130.0000 WHERE district_name = 'Ujung Latu';
UPDATE districts SET latitude = -3.8000, longitude = 130.8000 WHERE district_name = 'Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts SET latitude = -3.5000, longitude = 129.5000 WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur');

-- Kabupaten Kepulauan Aru
UPDATE districts SET latitude = -5.7667, longitude = 134.2167 WHERE district_name = 'Dobo';
UPDATE districts SET latitude = -6.5000, longitude = 134.5000 WHERE district_name = 'Aru Selatan';
UPDATE districts SET latitude = -6.0000, longitude = 134.5000 WHERE district_name = 'Aru Tengah';
UPDATE districts SET latitude = -5.5000, longitude = 134.5000 WHERE district_name = 'Aru Utara';
UPDATE districts SET latitude = -6.0000, longitude = 134.7500 WHERE district_name = 'Aru Tengah Timur';
UPDATE districts SET latitude = -6.2500, longitude = 134.5000 WHERE district_name = 'Aru Tengah Selatan';
UPDATE districts SET latitude = -5.4000, longitude = 134.8000 WHERE district_name = 'Aru Utara Timur Batuley';
UPDATE districts SET latitude = -5.9000, longitude = 134.1000 WHERE district_name = 'Sir-Sir';
UPDATE districts SET latitude = -6.1667, longitude = 134.3333 WHERE district_name = 'Pulau-Pulau Aru';

-- Kabupaten Maluku Barat Daya
UPDATE districts SET latitude = -7.9833, longitude = 131.3000 WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.4167, longitude = 131.5000 WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.2000, longitude = 130.8333 WHERE district_name = 'Selaru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6667, longitude = 131.1667 WHERE district_name = 'Wer Tamrian' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.5833, longitude = 131.5000 WHERE district_name = 'Wer Maktian' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.5000, longitude = 131.8333 WHERE district_name = 'Nirunmas' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.1667, longitude = 131.5833 WHERE district_name = 'Molu Maru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.8833, longitude = 126.3333 WHERE district_name = 'Wetar';
UPDATE districts SET latitude = -7.8000, longitude = 126.1667 WHERE district_name = 'Wetar Barat';
UPDATE districts SET latitude = -7.7500, longitude = 126.7500 WHERE district_name = 'Wetar Timur';
UPDATE districts SET latitude = -7.6333, longitude = 126.4500 WHERE district_name = 'Wetar Utara';
UPDATE districts SET latitude = -8.0667, longitude = 126.2833 WHERE district_name = 'Kisar Utara';
UPDATE districts SET latitude = -8.2500, longitude = 127.3333 WHERE district_name = 'Pulau Leti';
UPDATE districts SET latitude = -8.2167, longitude = 127.5833 WHERE district_name = 'Pulau Lakor';
UPDATE districts SET latitude = -8.1333, longitude = 128.6667 WHERE district_name = 'Dawelor Dawera';
UPDATE districts SET latitude = -8.2500, longitude = 127.4167 WHERE district_name = 'Pulau Moa';
UPDATE districts SET latitude = -7.8333, longitude = 129.7500 WHERE district_name = 'Babar';
UPDATE districts SET latitude = -7.9167, longitude = 129.9167 WHERE district_name = 'Babar Timur';
UPDATE districts SET latitude = -7.5000, longitude = 129.5000 WHERE district_name = 'Dai';
UPDATE districts SET latitude = -8.3333, longitude = 129.8333 WHERE district_name = 'Masela';
UPDATE districts SET latitude = -7.8333, longitude = 129.7500 WHERE district_name = 'Batarkusu';
UPDATE districts SET latitude = -7.5833, longitude = 127.4167 WHERE district_name = 'Romang';
UPDATE districts SET latitude = -8.2167, longitude = 127.4167 WHERE district_name = 'Leti Moa Lakor';
UPDATE districts SET latitude = -8.1000, longitude = 127.2000 WHERE district_name = 'Mdona Hiera';
UPDATE districts SET latitude = -7.4167, longitude = 131.5000 WHERE district_name = 'Tanimbar Utara' AND district_code = 'TNB';

-- Kabupaten Buru
UPDATE districts SET latitude = -3.3167, longitude = 127.0833 WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.1000, longitude = 126.6667 WHERE district_name = 'Air Buaya';
UPDATE districts SET latitude = -3.3667, longitude = 127.0500 WHERE district_name = 'Waeapo';
UPDATE districts SET latitude = -3.2500, longitude = 126.8333 WHERE district_name = 'Waplau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.4167, longitude = 126.5000 WHERE district_name = 'Batabual';
UPDATE districts SET latitude = -3.2000, longitude = 126.9000 WHERE district_name = 'Lolong Guba';
UPDATE districts SET latitude = -3.3833, longitude = 127.1333 WHERE district_name = 'Teluk Kaiely';
UPDATE districts SET latitude = -3.3000, longitude = 127.0000 WHERE district_name = 'Lilialy';
UPDATE districts SET latitude = -3.3500, longitude = 126.9500 WHERE district_name = 'Waelata';
UPDATE districts SET latitude = -3.0833, longitude = 126.3333 WHERE district_name = 'Fena Leisela';
UPDATE districts SET latitude = -3.7500, longitude = 126.5833 WHERE district_name = 'Kepala Madan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.6667, longitude = 126.6667 WHERE district_name = 'Waesama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.3667, longitude = 127.0500 WHERE district_name = 'Dayamurni';

-- Kabupaten Buru Selatan
UPDATE districts SET latitude = -3.8500, longitude = 126.6500 WHERE district_name = 'Namrole';
UPDATE districts SET latitude = -3.9167, longitude = 126.3333 WHERE district_name = 'Leksula';
UPDATE districts SET latitude = -3.7500, longitude = 126.5833 WHERE district_name = 'Kepala Madan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.6667, longitude = 126.6667 WHERE district_name = 'Waesama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.9833, longitude = 127.0000 WHERE district_name = 'Ambalau';
UPDATE districts SET latitude = -3.8333, longitude = 126.6667 WHERE district_name = 'Buru Selatan';

-- Kota Tual
UPDATE districts SET latitude = -5.6333, longitude = 132.7500 WHERE district_name = 'Dullah Utara';
UPDATE districts SET latitude = -5.6667, longitude = 132.7333 WHERE district_name = 'Dullah Selatan';
UPDATE districts SET latitude = -5.8333, longitude = 132.5833 WHERE district_name = 'Tayando';
UPDATE districts SET latitude = -5.6333, longitude = 132.7500 WHERE district_name = 'Pulau Dullah Utara';
UPDATE districts SET latitude = -5.9167, longitude = 132.9167 WHERE district_name = 'Kur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');

-- Kabupaten Maluku Tenggara
UPDATE districts SET latitude = -5.7500, longitude = 132.7500 WHERE district_name = 'Kei Kecil';
UPDATE districts SET latitude = -5.5000, longitude = 133.0000 WHERE district_name = 'Kei Besar';
UPDATE districts SET latitude = -5.7500, longitude = 132.8333 WHERE district_name = 'Kei Kecil Timur';
UPDATE districts SET latitude = -5.8333, longitude = 133.0000 WHERE district_name = 'Kei Besar Selatan';
UPDATE districts SET latitude = -5.7500, longitude = 132.6667 WHERE district_name = 'Kei Kecil Barat';
UPDATE districts SET latitude = -5.4167, longitude = 133.0833 WHERE district_name = 'Kei Besar Utara Timur';
UPDATE districts SET latitude = -5.6833, longitude = 132.7167 WHERE district_name = 'Hoat Sorbay';
UPDATE districts SET latitude = -5.6000, longitude = 132.7000 WHERE district_name = 'Manyeuw';
UPDATE districts SET latitude = -5.4167, longitude = 133.0000 WHERE district_name = 'Kei Besar Utara';
UPDATE districts SET latitude = -5.8333, longitude = 132.8333 WHERE district_name = 'Kei Kecil Timur Selatan';
UPDATE districts SET latitude = -5.9167, longitude = 132.9167 WHERE district_name = 'Pulau Kur';
UPDATE districts SET latitude = -5.9167, longitude = 132.9167 WHERE district_name = 'Kur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');

-- Kabupaten Maluku Tenggara Barat
UPDATE districts SET latitude = -7.9833, longitude = 131.3000 WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.4167, longitude = 131.5000 WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -8.2000, longitude = 130.8333 WHERE district_name = 'Selaru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.6667, longitude = 131.1667 WHERE district_name = 'Wer Tamrian' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.5833, longitude = 131.5000 WHERE district_name = 'Wer Maktian' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.5000, longitude = 131.8333 WHERE district_name = 'Nirunmas' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.1667, longitude = 131.5833 WHERE district_name = 'Molu Maru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.8333, longitude = 131.5833 WHERE district_name = 'Yaru';
UPDATE districts SET latitude = -7.6667, longitude = 131.6667 WHERE district_name = 'Kormomolin' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat');
UPDATE districts SET latitude = -7.2500, longitude = 131.9167 WHERE district_name = 'Fordata';

-- Update Desa/Kelurahan (Villages)
-- Kota Ambon - Kecamatan Sirimau
UPDATE villages SET latitude = -3.6953, longitude = 128.1817 WHERE village_name = 'Mardika';
UPDATE villages SET latitude = -3.6833, longitude = 128.1833 WHERE village_name = 'Batu Merah';
UPDATE villages SET latitude = -3.6981, longitude = 128.1758 WHERE village_name = 'Benteng';
UPDATE villages SET latitude = -3.7000, longitude = 128.1750 WHERE village_name = 'Wainitu';
UPDATE villages SET latitude = -3.6958, longitude = 128.1736 WHERE village_name = 'Honipopu';
UPDATE villages SET latitude = -3.6931, longitude = 128.1769 WHERE village_name = 'Rijali';
UPDATE villages SET latitude = -3.6894, longitude = 128.1889 WHERE village_name = 'Karang Panjang';
UPDATE villages SET latitude = -3.6917, longitude = 128.1708 WHERE village_name = 'Ahusen';
UPDATE villages SET latitude = -3.6703, longitude = 128.2031 WHERE village_name = 'Galala';
UPDATE villages SET latitude = -3.7083, longitude = 128.1917 WHERE village_name = 'Soya';

-- Kota Ambon - Kecamatan Nusaniwe
UPDATE villages SET latitude = -3.7073, longitude = 128.1602 WHERE village_name = 'Nusaniwe';
UPDATE villages SET latitude = -3.7033, longitude = 128.1692 WHERE village_name = 'Waihaong';
UPDATE villages SET latitude = -3.7111, longitude = 128.1639 WHERE village_name = 'Batu Gajah';
UPDATE villages SET latitude = -3.7069, longitude = 128.1583 WHERE village_name = 'Kudamati';
UPDATE villages SET latitude = -3.7333, longitude = 128.1500 WHERE village_name = 'Hatalae';
UPDATE villages SET latitude = -3.6500, longitude = 128.1333 WHERE village_name = 'Waiheru';
UPDATE villages SET latitude = -3.6167, longitude = 128.2667 WHERE village_name = 'Lateri';
UPDATE villages SET latitude = -3.7000, longitude = 128.1500 WHERE village_name = 'Urimessing';
UPDATE villages SET latitude = -3.6333, longitude = 128.2500 WHERE village_name = 'Passo';
UPDATE villages SET latitude = -3.7667, longitude = 128.2333 WHERE village_name = 'Hukurila';

-- Maluku Tengah - Kecamatan Salahutu
UPDATE villages SET latitude = -3.5881, longitude = 128.3244 WHERE village_name = 'Tulehu';
UPDATE villages SET latitude = -3.5167, longitude = 128.3333 WHERE village_name = 'Liang';
UPDATE villages SET latitude = -3.5667, longitude = 128.3833 WHERE village_name = 'Waai';
UPDATE villages SET latitude = -3.6000, longitude = 128.3500 WHERE village_name = 'Tial';
UPDATE villages SET latitude = -3.7500, longitude = 128.2167 WHERE village_name = 'Rutong';
UPDATE villages SET latitude = -3.5833, longitude = 128.2833 WHERE village_name = 'Tengah-Tengah';
UPDATE villages SET latitude = -3.5500, longitude = 128.2333 WHERE village_name = 'Morella';
UPDATE villages SET latitude = -3.5333, longitude = 128.2167 WHERE village_name = 'Mamala' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Salahutu');
UPDATE villages SET latitude = -3.5167, longitude = 128.1167 WHERE village_name = 'Hitu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Salahutu');
UPDATE villages SET latitude = -3.5333, longitude = 128.1000 WHERE village_name = 'Hila' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Salahutu');

-- Maluku Tengah - Kecamatan Leihitu
UPDATE villages SET latitude = -3.6333, longitude = 127.9500 WHERE village_name = 'Allang';
UPDATE villages SET latitude = -3.5167, longitude = 128.1167 WHERE village_name = 'Hitu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu');
UPDATE villages SET latitude = -3.5000, longitude = 128.1000 WHERE village_name = 'Kaitetu';
UPDATE villages SET latitude = -3.5333, longitude = 128.2167 WHERE village_name = 'Mamala' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu');
UPDATE villages SET latitude = -3.4833, longitude = 128.0667 WHERE village_name = 'Seith';
UPDATE villages SET latitude = -3.5167, longitude = 128.1167 WHERE village_name = 'Hitumessing';
UPDATE villages SET latitude = -3.5333, longitude = 128.1000 WHERE village_name = 'Hila' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu');
UPDATE villages SET latitude = -3.6667, longitude = 128.0167 WHERE village_name = 'Lilibooi';

-- Seram Bagian Barat - Kecamatan Kairatu
UPDATE villages SET latitude = -3.3833, longitude = 128.3333 WHERE village_name = 'Kairatu';
UPDATE villages SET latitude = -3.4167, longitude = 128.3167 WHERE village_name = 'Hatusua';
UPDATE villages SET latitude = -3.3667, longitude = 128.3500 WHERE village_name = 'Buria';
UPDATE villages SET latitude = -3.4000, longitude = 128.3500 WHERE village_name = 'Hatumeten';
UPDATE villages SET latitude = -3.3500, longitude = 128.4167 WHERE village_name = 'Lohia Sapalewa';
UPDATE villages SET latitude = -3.4333, longitude = 128.3000 WHERE village_name = 'Haturete';
UPDATE villages SET latitude = -3.4167, longitude = 128.4000 WHERE village_name = 'Rumahkay';
UPDATE villages SET latitude = -3.3333, longitude = 128.3000 WHERE village_name = 'Kamal';
UPDATE villages SET latitude = -3.3833, longitude = 128.3667 WHERE village_name = 'Murnaten';
UPDATE villages SET latitude = -3.3667, longitude = 128.4000 WHERE village_name = 'Lohiatala';

-- Buru - Kecamatan Namlea
UPDATE villages SET latitude = -3.2575, longitude = 127.0931 WHERE village_name = 'Namlea';
UPDATE villages SET latitude = -3.2167, longitude = 127.0500 WHERE village_name = 'Wamlana';
UPDATE villages SET latitude = -3.2833, longitude = 127.1167 WHERE village_name = 'Sawa';
UPDATE villages SET latitude = -3.3000, longitude = 127.1000 WHERE village_name = 'Waenetat';
UPDATE villages SET latitude = -3.2667, longitude = 127.0833 WHERE village_name = 'Wangongira';
UPDATE villages SET latitude = -3.3333, longitude = 127.0667 WHERE village_name = 'Waekasar';
UPDATE villages SET latitude = -3.1833, longitude = 127.0333 WHERE village_name = 'Jikumerasa';
UPDATE villages SET latitude = -3.2833, longitude = 127.0833 WHERE village_name = 'Waegeren';
UPDATE villages SET latitude = -3.3167, longitude = 127.0833 WHERE village_name = 'Waetawa';
UPDATE villages SET latitude = -3.2611, longitude = 127.0958 WHERE village_name = 'Kampung Baru';

-- Kota Tual - Kecamatan Dullah Utara
UPDATE villages SET latitude = -5.6572, longitude = 132.7321 WHERE village_name = 'Langgur';
UPDATE villages SET latitude = -5.6417, longitude = 132.7333 WHERE village_name = 'Ohoijang';
UPDATE villages SET latitude = -5.6333, longitude = 132.7500 WHERE village_name = 'Feer';
UPDATE villages SET latitude = -5.6167, longitude = 132.7500 WHERE village_name = 'Bombay';
UPDATE villages SET latitude = -5.6000, longitude = 132.7667 WHERE village_name = 'Ruat';
UPDATE villages SET latitude = -5.6333, longitude = 132.7333 WHERE village_name = 'Watdek';
UPDATE villages SET latitude = -5.7333, longitude = 132.7167 WHERE village_name = 'Debut';
UPDATE villages SET latitude = -5.6167, longitude = 132.7333 WHERE village_name = 'Dullah Laut';
UPDATE villages SET latitude = -5.7167, longitude = 132.7333 WHERE village_name = 'Sathean';
UPDATE villages SET latitude = -5.7000, longitude = 132.7167 WHERE village_name = 'Revav';
COMMIT;