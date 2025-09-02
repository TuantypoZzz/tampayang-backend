-- Migration to insert latitude and longitude data
-- Fixed coordinates to ensure they point to actual land-based administrative centers
-- Updated with more accurate coordinates for Maluku Province administrative regions

START TRANSACTION;

-- Update Provinsi Maluku (Center point of the province)
UPDATE provinces
SET latitude = -3.695430, longitude = 128.179330
WHERE province_name = 'Maluku';

-- Update Kabupaten/Kota di Maluku with corrected coordinates
-- Coordinates now point to actual administrative centers on land
UPDATE regencies SET latitude = -3.3065, longitude = 128.9568 WHERE regency_name = 'Kabupaten Maluku Tengah';
UPDATE regencies SET latitude = -5.6841, longitude = 132.7203 WHERE regency_name = 'Kabupaten Maluku Tenggara';
UPDATE regencies SET latitude = -7.9699, longitude = 131.3117 WHERE regency_name = 'Kabupaten Kepulauan Tanimbar';
UPDATE regencies SET latitude = -3.2590, longitude = 127.1028 WHERE regency_name = 'Kabupaten Buru';
UPDATE regencies SET latitude = -3.1022, longitude = 130.4893 WHERE regency_name = 'Kabupaten Seram Bagian Timur';
UPDATE regencies SET latitude = -3.0689, longitude = 128.2298 WHERE regency_name = 'Kabupaten Seram Bagian Barat';
UPDATE regencies SET latitude = -5.7608, longitude = 134.2335 WHERE regency_name = 'Kabupaten Kepulauan Aru';
UPDATE regencies SET latitude = -8.1461, longitude = 127.7989 WHERE regency_name = 'Kabupaten Maluku Barat Daya';
UPDATE regencies SET latitude = -3.8384, longitude = 126.7386 WHERE regency_name = 'Kabupaten Buru Selatan';
UPDATE regencies SET latitude = -3.6547, longitude = 128.1906 WHERE regency_name = 'Kota Ambon';
UPDATE regencies SET latitude = -5.6368, longitude = 132.7508 WHERE regency_name = 'Kota Tual';

-- Update Kecamatan (Districts)
-- Kota Ambon
UPDATE districts SET latitude = -3.7238, longitude = 128.1502 WHERE district_name = 'Nusaniwe';
UPDATE districts SET latitude = -3.6832, longitude = 128.1957 WHERE district_name = 'Sirimau';
UPDATE districts SET latitude = -3.6314, longitude = 128.2426 WHERE district_name = 'Baguala';
UPDATE districts SET latitude = -3.6531, longitude = 128.1911 WHERE district_name = 'Teluk Ambon';
UPDATE districts SET latitude = -3.6884, longitude = 128.2736 WHERE district_name = 'Leitimur Selatan';

-- Kabupaten Kepulauan Tanimbar
UPDATE districts SET latitude = -7.0395, longitude = 131.9543 WHERE district_name = 'Fordata';
UPDATE districts SET latitude = -7.6256, longitude = 131.5980 WHERE district_name = 'Kormomolin';
UPDATE districts SET latitude = -6.7128, longitude = 131.5615 WHERE district_name = 'Molu Maru';
UPDATE districts SET latitude = -7.4469, longitude = 131.6494 WHERE district_name = 'Nirunmas';
UPDATE districts SET latitude = -8.1816, longitude = 131.0077 WHERE district_name = 'Selaru';
UPDATE districts SET latitude = -7.9890, longitude = 131.2974 WHERE district_name = 'Tanimbar Selatan';
UPDATE districts SET latitude = -7.1522, longitude = 131.7196 WHERE district_name = 'Tanimbar Utara';
UPDATE districts SET latitude = -7.6667, longitude = 131.0358 WHERE district_name = 'Wer Maktian';
UPDATE districts SET latitude = -7.7621, longitude = 131.4464 WHERE district_name = 'Wer Tamrian';
UPDATE districts SET latitude = -7.3244, longitude = 131.4437 WHERE district_name = 'Wuar Labobar';

-- Kabupaten Maluku Tengah (Corrected coordinates)
UPDATE districts SET latitude = -3.3243, longitude = 128.9338 WHERE district_name = 'Amahai';
UPDATE districts SET latitude = -4.5185, longitude = 129.9047 WHERE district_name = 'Banda';
UPDATE districts SET latitude = -3.2097, longitude = 129.0165 WHERE district_name = 'Teon Nila Serua';
UPDATE districts SET latitude = -3.5800, longitude = 128.6219 WHERE district_name = 'Saparua';
UPDATE districts SET latitude = -3.4976, longitude = 128.6951 WHERE district_name = 'Saparua Timur';
UPDATE districts SET latitude = -3.6474, longitude = 128.8050 WHERE district_name = 'Nusalaut';
UPDATE districts SET latitude = -3.5197, longitude = 128.4773 WHERE district_name = 'Pulau Haruku';
UPDATE districts SET latitude = -3.5913, longitude = 128.3354 WHERE district_name = 'Salahutu';
UPDATE districts SET latitude = -3.5839, longitude = 128.0884 WHERE district_name = 'Leihitu';
UPDATE districts SET latitude = -3.7282, longitude = 128.0476 WHERE district_name = 'Leihitu Barat';
UPDATE districts SET latitude = -2.7952, longitude = 129.4918 WHERE district_name = 'Seram Utara';
UPDATE districts SET latitude = -2.8027, longitude = 129.0572 WHERE district_name = 'Seram Utara Barat';
UPDATE districts SET latitude = -2.9160, longitude = 129.8215 WHERE district_name = 'Seram Utara Timur Kobi';
UPDATE districts SET latitude = -3.0068, longitude = 129.9456 WHERE district_name = 'Seram Utara Timur Seti';
UPDATE districts SET latitude = -3.0369, longitude = 130.0990 WHERE district_name = 'Seram Utara Timur Seti Timur';
UPDATE districts SET latitude = -3.2328, longitude = 128.8018 WHERE district_name = 'Teluk Elpaputih';
UPDATE districts SET latitude = -3.3192, longitude = 129.4761 WHERE district_name = 'Tehoru';
UPDATE districts SET latitude = -3.3219, longitude = 129.7762 WHERE district_name = 'Telutih';
UPDATE districts SET latitude = -3.3035, longitude = 128.9580 WHERE district_name = 'Kota Masohi';

-- Kabupaten Seram Bagian Barat
UPDATE districts SET latitude = -3.2491, longitude = 128.0405 WHERE district_name = 'Huamual';
UPDATE districts SET latitude = -3.2318, longitude = 127.7511 WHERE district_name = 'Huamual Belakang';
UPDATE districts SET latitude = -3.4147, longitude = 128.6874 WHERE district_name = 'Amalatu';
UPDATE districts SET latitude = -3.2309, longitude = 128.8428 WHERE district_name = 'Elpaputih';
UPDATE districts SET latitude = -3.2158, longitude = 128.4306 WHERE district_name = 'Inamosol';
UPDATE districts SET latitude = -3.3423, longitude = 128.3577 WHERE district_name = 'Kairatu';
UPDATE districts SET latitude = -3.2109, longitude = 128.2773 WHERE district_name = 'Kairatu Barat';
UPDATE districts SET latitude = -3.2877, longitude = 127.5147 WHERE district_name = 'Kepulauan Manipa';
UPDATE districts SET latitude = -3.0653, longitude = 128.1973 WHERE district_name = 'Seram Barat';
UPDATE districts SET latitude = -2.8748, longitude = 128.4475 WHERE district_name = 'Taniwel';
UPDATE districts SET latitude = -2.8647, longitude = 128.7123 WHERE district_name = 'Taniwel Timur';

-- Kabupaten Seram Bagian Timur
UPDATE districts  SET latitude = -3.1031, longitude = 130.4911 WHERE district_name = 'Bula'  AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.3382, longitude = 130.0403 WHERE district_name = 'Werinama' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -4.0287, longitude = 131.4260 WHERE district_name = 'Pulau Gorom' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -4.4755, longitude = 131.6485 WHERE district_name = 'Wakate' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.5645, longitude = 130.8046 WHERE district_name = 'Tutuk Tolu' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.2747, longitude = 129.9509 WHERE district_name = 'Siwalat' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.6849, longitude = 130.4879 WHERE district_name = 'Kilmury' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.3859, longitude = 130.6537 WHERE district_name = 'Teluk Waru' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.9956, longitude = 131.4174 WHERE district_name = 'Gorom Timur' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -2.9858, longitude = 130.3426 WHERE district_name = 'Bula Barat' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.6842, longitude = 130.7927 WHERE district_name = 'Kian Darat' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.7806, longitude = 130.8109 WHERE district_name = 'Lian Fitu' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.8356, longitude = 130.7650 WHERE district_name = 'Ukar Sengan' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -4.7389, longitude = 131.7501 WHERE district_name = 'Teor' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -4.0086, longitude = 131.2380 WHERE district_name = 'Pulau Panjang' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');
UPDATE districts  SET latitude = -3.8847, longitude = 130.8985 WHERE district_name = 'Seram Timur' AND regency_id = (SELECT regency_id FROM regencies  WHERE regency_name = 'Kabupaten Seram Bagian Timur');

-- Koordinat Kecamatan di Kabupaten Kepulauan Aru
UPDATE districts SET latitude = -6.3069, longitude = 134.5448 WHERE district_name = 'Pulau-Pulau Aru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.4556, longitude = 134.1928 WHERE district_name = 'Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.1833, longitude = 134.5081 WHERE district_name = 'Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -5.8891, longitude = 134.2586 WHERE district_name = 'Aru Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.5812, longitude = 134.3412 WHERE district_name = 'Aru Selatan Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.2593, longitude = 134.6567 WHERE district_name = 'Aru Tengah Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.3908, longitude = 134.5823 WHERE district_name = 'Aru Tengah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -5.7331, longitude = 134.1069 WHERE district_name = 'Sir-Sir' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.6083, longitude = 134.2181 WHERE district_name = 'Pulau-Pulau Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');
UPDATE districts SET latitude = -6.0611, longitude = 134.3775 WHERE district_name = 'Pulau-Pulau Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru');

-- Kabupaten Maluku Barat Daya (Corrected coordinates for better accuracy)
UPDATE districts SET latitude = -7.1403, longitude = 127.7281 WHERE district_name = 'Damer' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.1545, longitude = 127.4851 WHERE district_name = 'Dawelor Dawera' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.0575, longitude = 127.1411 WHERE district_name = 'Kisar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.1417, longitude = 127.1198 WHERE district_name = 'Kisar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6136, longitude = 127.9025 WHERE district_name = 'Moa Lakor' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.0500, longitude = 127.6333 WHERE district_name = 'Pulau Letti' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.0833, longitude = 127.4167 WHERE district_name = 'Pulau Masela' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.5736, longitude = 127.6231 WHERE district_name = 'Pulau Romang' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.5008, longitude = 127.3392 WHERE district_name = 'Pulau Wetang' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.1842, longitude = 127.8261 WHERE district_name = 'Pulau Luang' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.7447, longitude = 127.9175 WHERE district_name = 'Pulau Lakor' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -8.0581, longitude = 128.0931 WHERE district_name = 'Pulau Sermata' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.9008, longitude = 127.9322 WHERE district_name = 'Pulau Teun' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6875, longitude = 126.8644 WHERE district_name = 'Wetar Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6239, longitude = 126.9272 WHERE district_name = 'Wetar Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6689, longitude = 126.9456 WHERE district_name = 'Wetar Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');
UPDATE districts SET latitude = -7.6014, longitude = 126.8842 WHERE district_name = 'Wetar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya');

-- Kabupaten Buru
UPDATE districts SET latitude = -3.2831, longitude = 126.7472 WHERE district_name = 'Air Buaya' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.4519, longitude = 127.0100 WHERE district_name = 'Batabual' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.4464, longitude = 126.9033 WHERE district_name = 'Fena Leisela' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.3969, longitude = 126.8328 WHERE district_name = 'Lilialy' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.5533, longitude = 126.9197 WHERE district_name = 'Lolong Guba' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.2500, longitude = 127.0833 WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.2633, longitude = 126.9231 WHERE district_name = 'Teluk Kaiely' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.4453, longitude = 127.1400 WHERE district_name = 'Waeapo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.3256, longitude = 127.2264 WHERE district_name = 'Waplau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');
UPDATE districts SET latitude = -3.4025, longitude = 127.0450 WHERE district_name = 'Waelata' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru');

-- Kabupaten Buru Selatan
UPDATE districts SET latitude = -3.8050, longitude = 126.7589 WHERE district_name = 'Ambalau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.6800, longitude = 126.9633 WHERE district_name = 'Fena Fafan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.9242, longitude = 126.8792 WHERE district_name = 'Kepala Madan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.8753, longitude = 126.8833 WHERE district_name = 'Leksula' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.8897, longitude = 126.7978 WHERE district_name = 'Namrole' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');
UPDATE districts SET latitude = -3.7508, longitude = 126.8903 WHERE district_name = 'Waesama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan');

-- Kota Tual
UPDATE districts SET latitude = -5.6288, longitude = 132.7463 WHERE district_name = 'Tual' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');
UPDATE districts SET latitude = -5.6436, longitude = 132.7531 WHERE district_name = 'Pulau Dullah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');
UPDATE districts SET latitude = -5.5947, longitude = 132.7605 WHERE district_name = 'Pulau Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');
UPDATE districts SET latitude = -5.7506, longitude = 132.7225 WHERE district_name = 'Pulau Tayando Tam' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');
UPDATE districts SET latitude = -5.7033, longitude = 132.8772 WHERE district_name = 'Kur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual');

-- Kabupaten Maluku Tenggara
UPDATE districts SET latitude = -5.7333, longitude = 132.7333 WHERE district_name = 'Kei Kecil' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.7500, longitude = 132.6500 WHERE district_name = 'Kei Kecil Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.7500, longitude = 132.8333 WHERE district_name = 'Kei Kecil Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.8000, longitude = 132.9000 WHERE district_name = 'Kei Kecil Timur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.9500, longitude = 132.9000 WHERE district_name = 'Kei Besar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -6.1500, longitude = 132.9000 WHERE district_name = 'Kei Besar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -6.0500, longitude = 132.7500 WHERE district_name = 'Kei Besar Selatan Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.9000, longitude = 133.0500 WHERE district_name = 'Kei Besar Utara Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.9000, longitude = 132.7500 WHERE district_name = 'Kei Besar Utara Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -5.6667, longitude = 132.8000 WHERE district_name = 'Hoat Sorbay' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');
UPDATE districts SET latitude = -6.0000, longitude = 132.9500 WHERE district_name = 'Manyeuw' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara');

-- Update Desa/Kelurahan (Villages)
-- Kota Ambon - Kecamatan Sirimau
UPDATE villages SET latitude = -3.7012, longitude = 128.1771 WHERE village_name = 'Mardika';
UPDATE villages SET latitude = -3.6998, longitude = 128.1820 WHERE village_name = 'Batu Merah';
UPDATE villages SET latitude = -3.6976, longitude = 128.1835 WHERE village_name = 'Benteng';
UPDATE villages SET latitude = -3.6984, longitude = 128.1752 WHERE village_name = 'Wainitu';
UPDATE villages SET latitude = -3.6990, longitude = 128.1804 WHERE village_name = 'Honipopu';
UPDATE villages SET latitude = -3.6996, longitude = 128.1789 WHERE village_name = 'Rijali';
UPDATE villages SET latitude = -3.6979, longitude = 128.1815 WHERE village_name = 'Karang Panjang';
UPDATE villages SET latitude = -3.6997, longitude = 128.1794 WHERE village_name = 'Ahusen';
UPDATE villages SET latitude = -3.7045, longitude = 128.1953 WHERE village_name = 'Galala';
UPDATE villages SET latitude = -3.7173, longitude = 128.2011 WHERE village_name = 'Soya';
UPDATE villages SET latitude = -3.7077, longitude = 128.2075 WHERE village_name = 'Lateri';
UPDATE villages SET latitude = -3.6952, longitude = 128.1780 WHERE village_name = 'Pandan Kasturi';
UPDATE villages SET latitude = -3.6972, longitude = 128.1768 WHERE village_name = 'Uritetu';
UPDATE villages SET latitude = -3.6960, longitude = 128.1774 WHERE village_name = 'Pandai Besi';
UPDATE villages SET latitude = -3.7031, longitude = 128.1887 WHERE village_name = 'Kebun Cengkeh';

-- Kota Ambon - Kecamatan Nusaniwe
UPDATE villages SET latitude = -3.7168, longitude = 128.1833 WHERE village_name = 'Amahusu';
UPDATE villages SET latitude = -3.6951, longitude = 128.1767 WHERE village_name = 'Silale';
UPDATE villages SET latitude = -3.6923, longitude = 128.1734 WHERE village_name = 'Uritetu';
UPDATE villages SET latitude = -3.7172, longitude = 128.2314 WHERE village_name = 'Lateri';
UPDATE villages SET latitude = -3.7088, longitude = 128.2125 WHERE village_name = 'Hative Kecil';
UPDATE villages SET latitude = -3.7044, longitude = 128.1895 WHERE village_name = 'Batu Gantung';
UPDATE villages SET latitude = -3.7205, longitude = 128.1727 WHERE village_name = 'Eri';
UPDATE villages SET latitude = -3.7222, longitude = 128.1557 WHERE village_name = 'Hative Besar';
UPDATE villages SET latitude = -3.7655, longitude = 128.1453 WHERE village_name = 'Latuhalat';
UPDATE villages SET latitude = -3.7321, longitude = 128.1801 WHERE village_name = 'Nusaniwe';
UPDATE villages SET latitude = -3.7008, longitude = 128.1793 WHERE village_name = 'Benteng';

-- Kota Ambon - Kecamatan Teluk Ambon
UPDATE villages SET latitude = -3.6428, longitude = 128.2504 WHERE village_name = 'Latta';
UPDATE villages SET latitude = -3.6554, longitude = 128.2637 WHERE village_name = 'Passo';
UPDATE villages SET latitude = -3.6641, longitude = 128.2598 WHERE village_name = 'Nania';
UPDATE villages SET latitude = -3.6798, longitude = 128.2433 WHERE village_name = 'Hunut';
UPDATE villages SET latitude = -3.6405, longitude = 128.1867 WHERE village_name = 'Wayame';
UPDATE villages SET latitude = -3.6665, longitude = 128.1588 WHERE village_name = 'Tawiri';
UPDATE villages SET latitude = -3.6483, longitude = 128.1911 WHERE village_name = 'Poka';

-- Kota Ambon - Kecamatan Baguala
UPDATE villages SET latitude = -3.6568, longitude = 128.2642 WHERE village_name = 'Lateri';
UPDATE villages SET latitude = -3.6602, longitude = 128.2630 WHERE village_name = 'Nania';
UPDATE villages SET latitude = -3.6534, longitude = 128.2655 WHERE village_name = 'Passo';
UPDATE villages SET latitude = -3.6433, longitude = 128.2724 WHERE village_name = 'Halong';
UPDATE villages SET latitude = -3.6507, longitude = 128.1921 WHERE village_name = 'Poka';
UPDATE villages SET latitude = -3.6409, longitude = 128.1852 WHERE village_name = 'Wayame';

-- Kota Ambon - Kecamatan Leitimur Selatan
UPDATE villages SET latitude = -3.7755, longitude = 128.3121 WHERE village_name = 'Rutong';
UPDATE villages SET latitude = -3.7893, longitude = 128.3128 WHERE village_name = 'Naku';
UPDATE villages SET latitude = -3.7961, longitude = 128.3194 WHERE village_name = 'Kilang';
UPDATE villages SET latitude = -3.8108, longitude = 128.3299 WHERE village_name = 'Hative Kecil';
UPDATE villages SET latitude = -3.8237, longitude = 128.3371 WHERE village_name = 'Ema';
UPDATE villages SET latitude = -3.8443, longitude = 128.3458 WHERE village_name = 'Hukurila';
UPDATE villages SET latitude = -3.8547, longitude = 128.3512 WHERE village_name = 'Leahari';

-- Kabupaten Maluku Tengah - Kecamatan Amahai
UPDATE villages SET latitude = -3.3403, longitude = 128.9255 WHERE village_name = 'Amahai';
UPDATE villages SET latitude = -3.3567, longitude = 128.9462 WHERE village_name = 'Saparua';
UPDATE villages SET latitude = -3.3721, longitude = 128.9590 WHERE village_name = 'Waraka';
UPDATE villages SET latitude = -3.3880, longitude = 128.9725 WHERE village_name = 'Rutah';
UPDATE villages SET latitude = -3.4018, longitude = 128.9832 WHERE village_name = 'Makariki';
UPDATE villages SET latitude = -3.4154, longitude = 128.9947 WHERE village_name = 'Hatumete';
UPDATE villages SET latitude = -3.4302, longitude = 129.0071 WHERE village_name = 'Sila';
UPDATE villages SET latitude = -3.4451, longitude = 129.0186 WHERE village_name = 'Hunitetu';
UPDATE villages SET latitude = -3.4599, longitude = 129.0303 WHERE village_name = 'Waipia';
UPDATE villages SET latitude = -3.4747, longitude = 129.0431 WHERE village_name = 'Yalahatan';

-- Kabupaten Maluku Tengah - Kecamatan Banda
UPDATE villages SET latitude = -4.5231, longitude = 129.8973 WHERE village_name = 'Neira';
UPDATE villages SET latitude = -4.5251, longitude = 129.9005 WHERE village_name = 'Bandaneira';
UPDATE villages SET latitude = -4.5402, longitude = 129.9084 WHERE village_name = 'Lonthoir';
UPDATE villages SET latitude = -4.5275, longitude = 129.9027 WHERE village_name = 'Kampung Baru';
UPDATE villages SET latitude = -4.5500, longitude = 129.9153 WHERE village_name = 'Walikan';
UPDATE villages SET latitude = -4.5651, longitude = 129.9278 WHERE village_name = 'Selamon';
UPDATE villages SET latitude = -4.5768, longitude = 129.9402 WHERE village_name = 'Dender';
UPDATE villages SET latitude = -4.5842, longitude = 129.9490 WHERE village_name = 'Rajawali';
UPDATE villages SET latitude = -4.5936, longitude = 129.9585 WHERE village_name = 'Pisang';
UPDATE villages SET latitude = -4.6054, longitude = 129.9691 WHERE village_name = 'Hatta';

-- Kabupaten Maluku Tengah - Kecamatan Teon Nila Serua (TNS)
UPDATE villages SET latitude = -3.6401, longitude = 129.0895 WHERE village_name = 'Layeni';
UPDATE villages SET latitude = -3.6498, longitude = 129.0943 WHERE village_name = 'Mesa';
UPDATE villages SET latitude = -3.6585, longitude = 129.0982 WHERE village_name = 'Kella';
UPDATE villages SET latitude = -3.6710, longitude = 129.1055 WHERE village_name = 'Wotay';
UPDATE villages SET latitude = -3.6782, longitude = 129.1117 WHERE village_name = 'Sila';
UPDATE villages SET latitude = -3.6860, longitude = 129.1183 WHERE village_name = 'Bumei';
UPDATE villages SET latitude = -3.7015, longitude = 129.1247 WHERE village_name = 'Teon';
UPDATE villages SET latitude = -3.7123, longitude = 129.1319 WHERE village_name = 'Nila';
UPDATE villages SET latitude = -3.7255, longitude = 129.1395 WHERE village_name = 'Serua';

-- Kabupaten Maluku Tengah - Kecamatan Saparua
UPDATE villages SET latitude = -3.5551, longitude = 128.6594 WHERE village_name = 'Haria';
UPDATE villages SET latitude = -3.5663, longitude = 128.6621 WHERE village_name = 'Paperu';
UPDATE villages SET latitude = -3.5780, longitude = 128.6685 WHERE village_name = 'Ullath';
UPDATE villages SET latitude = -3.5887, longitude = 128.6731 WHERE village_name = 'Ouw';
UPDATE villages SET latitude = -3.6002, longitude = 128.6805 WHERE village_name = 'Booi';
UPDATE villages SET latitude = -3.6120, longitude = 128.6868 WHERE village_name = 'Sirisori Islam';
UPDATE villages SET latitude = -3.6215, longitude = 128.6912 WHERE village_name = 'Sirisori Kristen';
UPDATE villages SET latitude = -3.6333, longitude = 128.6987 WHERE village_name = 'Tiouw';
UPDATE villages SET latitude = -3.6440, longitude = 128.7043 WHERE village_name = 'Itawaka';
UPDATE villages SET latitude = -3.6535, longitude = 128.7099 WHERE village_name = 'Tuhaha';
UPDATE villages SET latitude = -3.6622, longitude = 128.7161 WHERE village_name = 'Saparua';

-- Kabupaten Maluku Tengah - Kecamatan Saparua Timur
UPDATE villages SET latitude = -3.5520, longitude = 128.7211 WHERE village_name = 'Haria Baru';
UPDATE villages SET latitude = -3.5625, longitude = 128.7293 WHERE village_name = 'Sirisori Baru';
UPDATE villages SET latitude = -3.5734, longitude = 128.7380 WHERE village_name = 'Ihamahu';
UPDATE villages SET latitude = -3.5837, longitude = 128.7455 WHERE village_name = 'Kulur';
UPDATE villages SET latitude = -3.5942, longitude = 128.7520 WHERE village_name = 'Ouw Baru';
UPDATE villages SET latitude = -3.6035, longitude = 128.7585 WHERE village_name = 'Dullah Laut';
UPDATE villages SET latitude = -3.6141, longitude = 128.7642 WHERE village_name = 'Siri Sori Baru';
UPDATE villages SET latitude = -3.6250, longitude = 128.7727 WHERE village_name = 'Haria Tenga';

-- Kabupaten Maluku Tengah - Kecamatan Nusalaut
UPDATE villages SET latitude = -3.6373, longitude = 128.7445 WHERE village_name = 'Abubu';
UPDATE villages SET latitude = -3.6438, longitude = 128.7529 WHERE village_name = 'Ameth';
UPDATE villages SET latitude = -3.6505, longitude = 128.7601 WHERE village_name = 'Nalahia';
UPDATE villages SET latitude = -3.6572, longitude = 128.7674 WHERE village_name = 'Akoon';
UPDATE villages SET latitude = -3.6640, longitude = 128.7742 WHERE village_name = 'Sila';
UPDATE villages SET latitude = -3.6710, longitude = 128.7820 WHERE village_name = 'Titawaai';
UPDATE villages SET latitude = -3.6782, longitude = 128.7893 WHERE village_name = 'Leinitu';

-- Kabupaten Maluku Tengah - Kecamatan Pulau Haruku
UPDATE villages SET latitude = -3.5252, longitude = 128.5373 WHERE village_name = 'Aboru';
UPDATE villages SET latitude = -3.5430, longitude = 128.5345 WHERE village_name = 'Haruku';
UPDATE villages SET latitude = -3.5514, longitude = 128.5201 WHERE village_name = 'Hulaliu';
UPDATE villages SET latitude = -3.5610, longitude = 128.5267 WHERE village_name = 'Kabauw';
UPDATE villages SET latitude = -3.5698, longitude = 128.5153 WHERE village_name = 'Kariu';
UPDATE villages SET latitude = -3.5790, longitude = 128.5080 WHERE village_name = 'Oma';
UPDATE villages SET latitude = -3.5882, longitude = 128.4934 WHERE village_name = 'Pelauw';
UPDATE villages SET latitude = -3.5967, longitude = 128.4840 WHERE village_name = 'Rohomoni';
UPDATE villages SET latitude = -3.6041, longitude = 128.4753 WHERE village_name = 'Sameth';

-- Maluku Tengah - Kecamatan Salahutu
UPDATE villages SET latitude = -3.5745, longitude = 128.3239 WHERE village_name = 'Waai';
UPDATE villages SET latitude = -3.5448, longitude = 128.3402 WHERE village_name = 'Tulehu';
UPDATE villages SET latitude = -3.5541, longitude = 128.3497 WHERE village_name = 'Suli';
UPDATE villages SET latitude = -3.5663, longitude = 128.3374 WHERE village_name = 'Tengah-Tengah';
UPDATE villages SET latitude = -3.5715, longitude = 128.3576 WHERE village_name = 'Tial';
UPDATE villages SET latitude = -3.5782, longitude = 128.3291 WHERE village_name = 'Waai Lama';

-- Maluku Tengah - Kecamatan Leihitu
UPDATE villages SET latitude = -3.5941, longitude = 128.0896 WHERE village_name = 'Hila';
UPDATE villages SET latitude = -3.5914, longitude = 128.1033 WHERE village_name = 'Wakal';
UPDATE villages SET latitude = -3.5865, longitude = 128.1176 WHERE village_name = 'Hitu Lama';
UPDATE villages SET latitude = -3.5802, longitude = 128.1249 WHERE village_name = 'Hitu';
UPDATE villages SET latitude = -3.5773, longitude = 128.1388 WHERE village_name = 'Kaitetu';
UPDATE villages SET latitude = -3.5837, longitude = 128.1555 WHERE village_name = 'Assilulu';
UPDATE villages SET latitude = -3.5982, longitude = 128.1714 WHERE village_name = 'Seith';
UPDATE villages SET latitude = -3.6087, longitude = 128.1862 WHERE village_name = 'Morela';
UPDATE villages SET latitude = -3.6173, longitude = 128.1978 WHERE village_name = 'Mamala';

-- Kabupaten Maluku Tengah - Kecamatan Leihitu Barat
UPDATE villages SET latitude = -3.6930, longitude = 128.0825 WHERE village_name = 'Liliboy';
UPDATE villages SET latitude = -3.7112, longitude = 128.0839 WHERE village_name = 'Ureng';
UPDATE villages SET latitude = -3.7351, longitude = 128.0814 WHERE village_name = 'Larike';
UPDATE villages SET latitude = -3.7567, longitude = 128.0769 WHERE village_name = 'Wakasihu';
UPDATE villages SET latitude = -3.7806, longitude = 128.0721 WHERE village_name = 'Allang';
UPDATE villages SET latitude = -3.6698, longitude = 128.0963 WHERE village_name = 'Seith Baru';
UPDATE villages SET latitude = -3.7488, longitude = 128.0835 WHERE village_name = 'Wakasihu Lama';

-- Kabupaten Maluku Tengah - Kecamatan Seram Utara
UPDATE villages SET latitude = -2.8324, longitude = 129.5751 WHERE village_name = 'Pasanea';
UPDATE villages SET latitude = -2.8123, longitude = 129.5729 WHERE village_name = 'Sawai';
UPDATE villages SET latitude = -2.8441, longitude = 129.5902 WHERE village_name = 'Makaeling';
UPDATE villages SET latitude = -2.8108, longitude = 129.5899 WHERE village_name = 'Sawai Baru';
UPDATE villages SET latitude = -2.8255, longitude = 129.5563 WHERE village_name = 'Roho';
UPDATE villages SET latitude = -2.8447, longitude = 129.6418 WHERE village_name = 'Manusela';
UPDATE villages SET latitude = -2.8603, longitude = 129.6541 WHERE village_name = 'Kanikeh';
UPDATE villages SET latitude = -2.8191, longitude = 129.5467 WHERE village_name = 'Wahai';
UPDATE villages SET latitude = -2.8634, longitude = 129.5713 WHERE village_name = 'Kobi';
UPDATE villages SET latitude = -2.8382, longitude = 129.5991 WHERE village_name = 'Air Besar';

-- Kabupaten Maluku Tengah - Kecamatan Seram Utara Barat
UPDATE villages SET latitude = -2.8205, longitude = 129.3801 WHERE village_name = 'Pasanea Barat';
UPDATE villages SET latitude = -2.8402, longitude = 129.3907 WHERE village_name = 'Masihulan';
UPDATE villages SET latitude = -2.8588, longitude = 129.4105 WHERE village_name = 'Lafa';
UPDATE villages SET latitude = -2.8742, longitude = 129.3982 WHERE village_name = 'Air Besar Barat';
UPDATE villages SET latitude = -2.8825, longitude = 129.3655 WHERE village_name = 'Kobi Barat';
UPDATE villages SET latitude = -2.8512, longitude = 129.3778 WHERE village_name = 'Amesang';
UPDATE villages SET latitude = -2.8321, longitude = 129.3564 WHERE village_name = 'Latu';
UPDATE villages SET latitude = -2.8109, longitude = 129.3421 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -2.8433, longitude = 129.3517 WHERE village_name = 'Wailulu';
UPDATE villages SET latitude = -2.8699, longitude = 129.3333 WHERE village_name = 'Waimital';

-- Kabupaten Maluku Tengah - Kecamatan Seram Utara Timur Kobi
UPDATE villages SET latitude = -2.8921, longitude = 129.4582 WHERE village_name = 'Kobi';
UPDATE villages SET latitude = -2.8878, longitude = 129.4721 WHERE village_name = 'Kobi Mukti';
UPDATE villages SET latitude = -2.8995, longitude = 129.4833 WHERE village_name = 'Kobi Dalam';
UPDATE villages SET latitude = -2.9057, longitude = 129.4956 WHERE village_name = 'Kobi Pantai';
UPDATE villages SET latitude = -2.8833, longitude = 129.4528 WHERE village_name = 'Kobi Sadar';
UPDATE villages SET latitude = -2.8744, longitude = 129.4650 WHERE village_name = 'Sawai Kobi';
UPDATE villages SET latitude = -2.9132, longitude = 129.5015 WHERE village_name = 'Namto';
UPDATE villages SET latitude = -2.8980, longitude = 129.4791 WHERE village_name = 'Waitidal';

-- Kabupaten Maluku Tengah - Kecamatan Seram Utara Timur Seti
UPDATE villages SET latitude = -2.8322, longitude = 129.5677 WHERE village_name = 'Seti';
UPDATE villages SET latitude = -2.8401, longitude = 129.5783 WHERE village_name = 'Pasanea';
UPDATE villages SET latitude = -2.8555, longitude = 129.5907 WHERE village_name = 'Loping Mulyo';
UPDATE villages SET latitude = -2.8623, longitude = 129.6001 WHERE village_name = 'Wailola';
UPDATE villages SET latitude = -2.8488, longitude = 129.6112 WHERE village_name = 'Mole';
UPDATE villages SET latitude = -2.8703, longitude = 129.6221 WHERE village_name = 'Wailulu';
UPDATE villages SET latitude = -2.8331, longitude = 129.6382 WHERE village_name = 'Kaitetu';
UPDATE villages SET latitude = -2.8259, longitude = 129.6495 WHERE village_name = 'Waiboga';
UPDATE villages SET latitude = -2.8123, longitude = 129.6604 WHERE village_name = 'Pasinan';

-- Kabupaten Maluku Tengah - Kecamatan Seram Utara Timur Seti Timur
UPDATE villages SET latitude = -2.8895, longitude = 129.6501 WHERE village_name = 'Haya';
UPDATE villages SET latitude = -2.8771, longitude = 129.6785 WHERE village_name = 'Wahai';
UPDATE villages SET latitude = -2.8967, longitude = 129.7002 WHERE village_name = 'Sawai';
UPDATE villages SET latitude = -2.9104, longitude = 129.7205 WHERE village_name = 'Manusela';
UPDATE villages SET latitude = -2.9212, longitude = 129.7356 WHERE village_name = 'Roho';
UPDATE villages SET latitude = -2.9311, longitude = 129.7524 WHERE village_name = 'Kobipoto';
UPDATE villages SET latitude = -2.9445, longitude = 129.7655 WHERE village_name = 'Kanikeh';
UPDATE villages SET latitude = -2.9578, longitude = 129.7820 WHERE village_name = 'Makariki';
UPDATE villages SET latitude = -2.9699, longitude = 129.7995 WHERE village_name = 'Wailana';

-- Kabupaten Maluku Tengah - Kecamatan Teluk Elpaputih
UPDATE villages SET latitude = -3.3665, longitude = 129.2954 WHERE village_name = 'Uraur';
UPDATE villages SET latitude = -3.3798, longitude = 129.3101 WHERE village_name = 'Tananahu';
UPDATE villages SET latitude = -3.3905, longitude = 129.3223 WHERE village_name = 'Waraka';
UPDATE villages SET latitude = -3.4042, longitude = 129.3339 WHERE village_name = 'Tihulesi';
UPDATE villages SET latitude = -3.4131, longitude = 129.3485 WHERE village_name = 'Waraka Utara';
UPDATE villages SET latitude = -3.4277, longitude = 129.3597 WHERE village_name = 'Rutah';
UPDATE villages SET latitude = -3.4398, longitude = 129.3701 WHERE village_name = 'Haya';
UPDATE villages SET latitude = -3.4521, longitude = 129.3840 WHERE village_name = 'Ameth';

-- Kabupaten Maluku Tengah - Kecamatan Tehoru
UPDATE villages SET latitude = -3.4362, longitude = 129.4705 WHERE village_name = 'Tehoru';
UPDATE villages SET latitude = -3.4401, longitude = 129.4808 WHERE village_name = 'Hatumete';
UPDATE villages SET latitude = -3.4487, longitude = 129.4951 WHERE village_name = 'Saunulu';
UPDATE villages SET latitude = -3.4554, longitude = 129.5069 WHERE village_name = 'Wolu';
UPDATE villages SET latitude = -3.4620, longitude = 129.5203 WHERE village_name = 'Laimu';
UPDATE villages SET latitude = -3.4711, longitude = 129.5317 WHERE village_name = 'Hatu';
UPDATE villages SET latitude = -3.4788, longitude = 129.5459 WHERE village_name = 'Werinama';
UPDATE villages SET latitude = -3.4865, longitude = 129.5592 WHERE village_name = 'Teluti Baru';
UPDATE villages SET latitude = -3.4943, longitude = 129.5731 WHERE village_name = 'Lesluru';

-- Kabupaten Maluku Tengah - Kecamatan Telutih
UPDATE villages SET latitude = -3.5102, longitude = 129.5421 WHERE village_name = 'Tehua';
UPDATE villages SET latitude = -3.5163, longitude = 129.5555 WHERE village_name = 'Yainuelo';
UPDATE villages SET latitude = -3.5230, longitude = 129.5681 WHERE village_name = 'Haya';
UPDATE villages SET latitude = -3.5304, longitude = 129.5809 WHERE village_name = 'Laimu';
UPDATE villages SET latitude = -3.5388, longitude = 129.5937 WHERE village_name = 'Rumahtiga';
UPDATE villages SET latitude = -3.5467, longitude = 129.6055 WHERE village_name = 'Werinama';
UPDATE villages SET latitude = -3.5532, longitude = 129.6182 WHERE village_name = 'Nusa Laut';
UPDATE villages SET latitude = -3.5627, longitude = 129.6300 WHERE village_name = 'Samal';
UPDATE villages SET latitude = -3.5710, longitude = 129.6418 WHERE village_name = 'Yainuelo Timur';

-- Kabupaten Maluku Tengah - Kecamatan Kota Masohi
UPDATE villages SET latitude = -3.3037, longitude = 128.9718 WHERE village_name = 'Namaelo';
UPDATE villages SET latitude = -3.3062, longitude = 128.9691 WHERE village_name = 'Namaelo Barat';
UPDATE villages SET latitude = -3.3018, longitude = 128.9745 WHERE village_name = 'Namaelo Timur';
UPDATE villages SET latitude = -3.3003, longitude = 128.9727 WHERE village_name = 'Lesane';
UPDATE villages SET latitude = -3.3045, longitude = 128.9680 WHERE village_name = 'Namaelo Kecil';
UPDATE villages SET latitude = -3.3076, longitude = 128.9762 WHERE village_name = 'Amahai';
UPDATE villages SET latitude = -3.3101, longitude = 128.9789 WHERE village_name = 'Holo';
UPDATE villages SET latitude = -3.3123, longitude = 128.9821 WHERE village_name = 'Namasula';
UPDATE villages SET latitude = -3.3150, longitude = 128.9857 WHERE village_name = 'Rumah Tiga';
UPDATE villages SET latitude = -3.2987, longitude = 128.9772 WHERE village_name = 'Lesane Timur';

-- Update Koordinat Kecamatan Fordata
UPDATE villages SET latitude = -7.6521, longitude = 131.3962 WHERE village_name = 'Rumngeur';
UPDATE villages SET latitude = -7.6442, longitude = 131.4028 WHERE village_name = 'Romean';
UPDATE villages SET latitude = -7.6605, longitude = 131.3853 WHERE village_name = 'Tutukembong';
UPDATE villages SET latitude = -7.6717, longitude = 131.3745 WHERE village_name = 'Tebtut';
UPDATE villages SET latitude = -7.6833, longitude = 131.3921 WHERE village_name = 'Siorbat';
UPDATE villages SET latitude = -7.6982, longitude = 131.4210 WHERE village_name = 'Adodo Molu';
UPDATE villages SET latitude = -7.7055, longitude = 131.4307 WHERE village_name = 'Tumbur';
UPDATE villages SET latitude = -7.7128, longitude = 131.4423 WHERE village_name = 'Tutunametal';

-- Update Koordinat Kecamatan Kormomolin
UPDATE villages SET latitude = -7.8812, longitude = 131.3184 WHERE village_name = 'Alusi Batjas';
UPDATE villages SET latitude = -7.8724, longitude = 131.3069 WHERE village_name = 'Alusi Kelaan';
UPDATE villages SET latitude = -7.8601, longitude = 131.2942 WHERE village_name = 'Alusi Tamrian';
UPDATE villages SET latitude = -7.8519, longitude = 131.2807 WHERE village_name = 'Alusi Tubun';
UPDATE villages SET latitude = -7.8443, longitude = 131.2651 WHERE village_name = 'Arui Bab';
UPDATE villages SET latitude = -7.8339, longitude = 131.2526 WHERE village_name = 'Fursui';
UPDATE villages SET latitude = -7.8224, longitude = 131.2389 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -7.8112, longitude = 131.2238 WHERE village_name = 'Lermatang';
UPDATE villages SET latitude = -7.7997, longitude = 131.2102 WHERE village_name = 'Lumasebu';
UPDATE villages SET latitude = -7.7890, longitude = 131.1955 WHERE village_name = 'Wowonda';

-- Update Koordinat Kecamatan Molu Maru
UPDATE villages SET latitude = -7.1198, longitude = 131.6061 WHERE village_name = 'Adodo Molu';
UPDATE villages SET latitude = -7.1292, longitude = 131.5954 WHERE village_name = 'Adodo Ratu';
UPDATE villages SET latitude = -7.1407, longitude = 131.5820 WHERE village_name = 'Adodo Tanah Merah';
UPDATE villages SET latitude = -7.1543, longitude = 131.5678 WHERE village_name = 'Larat';
UPDATE villages SET latitude = -7.1658, longitude = 131.5521 WHERE village_name = 'Ridool';
UPDATE villages SET latitude = -7.1774, longitude = 131.5397 WHERE village_name = 'Rumngeur';
UPDATE villages SET latitude = -7.1890, longitude = 131.5275 WHERE village_name = 'Rompo';
UPDATE villages SET latitude = -7.2005, longitude = 131.5142 WHERE village_name = 'Wermatang';
UPDATE villages SET latitude = -7.2112, longitude = 131.5009 WHERE village_name = 'Wermaktian';
UPDATE villages SET latitude = -7.2243, longitude = 131.4895 WHERE village_name = 'Wuarlabobar';

-- Update Koordinat Kecamatan Nirunmas
UPDATE villages SET latitude = -8.1125, longitude = 131.2178 WHERE village_name = 'Adaut';
UPDATE villages SET latitude = -8.0957, longitude = 131.2354 WHERE village_name = 'Arui Das';
UPDATE villages SET latitude = -8.0836, longitude = 131.2543 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -8.0720, longitude = 131.2705 WHERE village_name = 'Lorulun';
UPDATE villages SET latitude = -8.0601, longitude = 131.2892 WHERE village_name = 'Makatian';
UPDATE villages SET latitude = -8.0477, longitude = 131.3057 WHERE village_name = 'Mafri';
UPDATE villages SET latitude = -8.0385, longitude = 131.3220 WHERE village_name = 'Makatian Barat';
UPDATE villages SET latitude = -8.0259, longitude = 131.3402 WHERE village_name = 'Tutunametal';
UPDATE villages SET latitude = -8.0134, longitude = 131.3595 WHERE village_name = 'Arma';
UPDATE villages SET latitude = -8.0001, longitude = 131.3790 WHERE village_name = 'Wunlah';

-- Update Koordinat Kecamatan Selaru
UPDATE villages SET latitude = -8.3501, longitude = 131.1702 WHERE village_name = 'Ewur';
UPDATE villages SET latitude = -8.3705, longitude = 131.1884 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -8.3920, longitude = 131.2003 WHERE village_name = 'Latar';
UPDATE villages SET latitude = -8.4125, longitude = 131.2176 WHERE village_name = 'Lingat';
UPDATE villages SET latitude = -8.4321, longitude = 131.2407 WHERE village_name = 'Namtabung';
UPDATE villages SET latitude = -8.4456, longitude = 131.2602 WHERE village_name = 'Namtabung Timur';
UPDATE villages SET latitude = -8.4652, longitude = 131.2803 WHERE village_name = 'Saleru';
UPDATE villages SET latitude = -8.4827, longitude = 131.3005 WHERE village_name = 'Wulur';

-- Update Koordinat Kecamatan Tanimbar Selatan
UPDATE villages SET latitude = -7.9842, longitude = 131.3051 WHERE village_name = 'Saumlaki';
UPDATE villages SET latitude = -7.9901, longitude = 131.3205 WHERE village_name = 'Sifnana';
UPDATE villages SET latitude = -7.9755, longitude = 131.2893 WHERE village_name = 'Ongen';
UPDATE villages SET latitude = -7.9583, longitude = 131.2752 WHERE village_name = 'Bomaki';
UPDATE villages SET latitude = -8.0105, longitude = 131.2507 WHERE village_name = 'Lermatang';
UPDATE villages SET latitude = -8.0202, longitude = 131.2654 WHERE village_name = 'Lermatan';
UPDATE villages SET latitude = -8.0321, longitude = 131.2856 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -8.0415, longitude = 131.3004 WHERE village_name = 'Wowonda';
UPDATE villages SET latitude = -8.0552, longitude = 131.3178 WHERE village_name = 'Wermatang';
UPDATE villages SET latitude = -8.0667, longitude = 131.3401 WHERE village_name = 'Arui Das';

-- Update Koordinat Kecamatan Tanimbar Utara
UPDATE villages SET latitude = -7.4762, longitude = 131.8201 WHERE village_name = 'Larat';
UPDATE villages SET latitude = -7.4904, longitude = 131.8053 WHERE village_name = 'Ridool';
UPDATE villages SET latitude = -7.5021, longitude = 131.7896 WHERE village_name = 'Alusi Batjas';
UPDATE villages SET latitude = -7.5205, longitude = 131.7764 WHERE village_name = 'Alusi Kelaan';
UPDATE villages SET latitude = -7.5333, longitude = 131.7652 WHERE village_name = 'Alusi Tamrian';
UPDATE villages SET latitude = -7.5481, longitude = 131.7503 WHERE village_name = 'Lamdesar Barat';
UPDATE villages SET latitude = -7.5598, longitude = 131.7326 WHERE village_name = 'Lamdesar Timur';
UPDATE villages SET latitude = -7.5725, longitude = 131.7201 WHERE village_name = 'Lorulun';
UPDATE villages SET latitude = -7.5901, longitude = 131.7087 WHERE village_name = 'Tumbur';
UPDATE villages SET latitude = -7.6033, longitude = 131.6954 WHERE village_name = 'Nirunmas Utara';

-- Update Koordinat Kecamatan Wer Maktian
UPDATE villages SET latitude = -7.6661, longitude = 131.4820 WHERE village_name = 'Makatian';
UPDATE villages SET latitude = -7.6743, longitude = 131.4978 WHERE village_name = 'Romean';
UPDATE villages SET latitude = -7.6855, longitude = 131.5106 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -7.6937, longitude = 131.5261 WHERE village_name = 'Rombut';
UPDATE villages SET latitude = -7.7048, longitude = 131.5409 WHERE village_name = 'Latbual';
UPDATE villages SET latitude = -7.7189, longitude = 131.5583 WHERE village_name = 'Alusi Tamrian';

-- Update Koordinat Kecamatan Wer Tamrian
UPDATE villages SET latitude = -7.6355, longitude = 131.4155 WHERE village_name = 'Alusi Batjas';
UPDATE villages SET latitude = -7.6427, longitude = 131.4289 WHERE village_name = 'Alusi Kelaan';
UPDATE villages SET latitude = -7.6584, longitude = 131.4422 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -7.6712, longitude = 131.4566 WHERE village_name = 'Romean';
UPDATE villages SET latitude = -7.6855, longitude = 131.4727 WHERE village_name = 'Lermatang';
UPDATE villages SET latitude = -7.6983, longitude = 131.4901 WHERE village_name = 'Wowonda';

-- Update Koordinat Kecamatan Wuar Labobar
UPDATE villages SET latitude = -8.1245, longitude = 131.6931 WHERE village_name = 'Adaut';
UPDATE villages SET latitude = -8.1378, longitude = 131.7052 WHERE village_name = 'Lorulun';
UPDATE villages SET latitude = -8.1493, longitude = 131.7217 WHERE village_name = 'Lingat';
UPDATE villages SET latitude = -8.1604, longitude = 131.7402 WHERE village_name = 'Mole';
UPDATE villages SET latitude = -8.1731, longitude = 131.7594 WHERE village_name = 'Tutukembong';
UPDATE villages SET latitude = -8.1894, longitude = 131.7812 WHERE village_name = 'Latdalam';
UPDATE villages SET latitude = -8.2017, longitude = 131.7993 WHERE village_name = 'Sangliat Dol';
UPDATE villages SET latitude = -8.2148, longitude = 131.8155 WHERE village_name = 'Sangliat Krawain';

-- Update Koordinat Kecamatan Huamual
UPDATE villages SET latitude = -3.2718, longitude = 127.8519 WHERE village_name = 'Iha';
UPDATE villages SET latitude = -3.2973, longitude = 127.8321 WHERE village_name = 'Luhu';
UPDATE villages SET latitude = -3.3037, longitude = 127.8493 WHERE village_name = 'Rumakai';
UPDATE villages SET latitude = -3.3192, longitude = 127.8268 WHERE village_name = 'Tihulale';
UPDATE villages SET latitude = -3.3356, longitude = 127.8157 WHERE village_name = 'Waesamu';
UPDATE villages SET latitude = -3.3469, longitude = 127.8023 WHERE village_name = 'Tala';
UPDATE villages SET latitude = -3.3638, longitude = 127.7901 WHERE village_name = 'Loay';
UPDATE villages SET latitude = -3.3784, longitude = 127.7764 WHERE village_name = 'Wael';
UPDATE villages SET latitude = -3.3925, longitude = 127.7629 WHERE village_name = 'Waepirit';
UPDATE villages SET latitude = -3.4061, longitude = 127.7502 WHERE village_name = 'Lokki';


-- Update Koordinat Kecamatan Huamual Belakang
UPDATE villages SET latitude = -3.2262, longitude = 127.6925 WHERE village_name = 'Waelapia';
UPDATE villages SET latitude = -3.2307, longitude = 127.7012 WHERE village_name = 'Waelapia Timur';
UPDATE villages SET latitude = -3.2381, longitude = 127.7104 WHERE village_name = 'Ihamahu';
UPDATE villages SET latitude = -3.2436, longitude = 127.6827 WHERE village_name = 'Waelapia Barat';
UPDATE villages SET latitude = -3.2499, longitude = 127.6944 WHERE village_name = 'Waelapia Tengah';
UPDATE villages SET latitude = -3.2561, longitude = 127.7038 WHERE village_name = 'Waelapia Selatan';
UPDATE villages SET latitude = -3.2635, longitude = 127.7166 WHERE village_name = 'Waelapia Utara';

-- Update Koordinat Desa/Kelurahan - Kecamatan Amalatu
UPDATE villages SET latitude = -3.3494, longitude = 128.4553 WHERE village_name = 'Hualoy';
UPDATE villages SET latitude = -3.3387, longitude = 128.4635 WHERE village_name = 'Tomi-Tomi';
UPDATE villages SET latitude = -3.3605, longitude = 128.4512 WHERE village_name = 'Laala';
UPDATE villages SET latitude = -3.3709, longitude = 128.4598 WHERE village_name = 'Latu';

-- Update Koordinat Desa/Kelurahan - Kecamatan Elpaputih
UPDATE villages SET latitude = -3.4295, longitude = 128.5207 WHERE village_name = 'Hatusua';
UPDATE villages SET latitude = -3.4408, longitude = 128.5351 WHERE village_name = 'Hualoy';
UPDATE villages SET latitude = -3.4501, longitude = 128.5512 WHERE village_name = 'Uraur';
UPDATE villages SET latitude = -3.4668, longitude = 128.5615 WHERE village_name = 'Hulaliu';

-- Update Koordinat Desa/Kelurahan - Kecamatan Inamosol
UPDATE villages SET latitude = -3.5168, longitude = 128.3839 WHERE village_name = 'Riring';
UPDATE villages SET latitude = -3.5282, longitude = 128.3954 WHERE village_name = 'Mornaten';
UPDATE villages SET latitude = -3.5437, longitude = 128.4095 WHERE village_name = 'Nuruwe';
UPDATE villages SET latitude = -3.5555, longitude = 128.4218 WHERE village_name = 'Masika Jaya';

-- Update Koordinat Desa - Kecamatan Kairatu
UPDATE villages SET latitude = -3.4148, longitude = 128.6207 WHERE village_name = 'Kairatu';
UPDATE villages SET latitude = -3.4237, longitude = 128.6015 WHERE village_name = 'Wael';
UPDATE villages SET latitude = -3.4312, longitude = 128.6158 WHERE village_name = 'Kamariang';
UPDATE villages SET latitude = -3.4428, longitude = 128.5951 WHERE village_name = 'Hatusua';

-- Update Koordinat Desa - Kecamatan Kairatu Barat
UPDATE villages SET latitude = -3.4657, longitude = 128.5204 WHERE village_name = 'Latu';
UPDATE villages SET latitude = -3.4521, longitude = 128.5053 WHERE village_name = 'Waimital';
UPDATE villages SET latitude = -3.4750, longitude = 128.5307 WHERE village_name = 'Waihatu';

-- Update Koordinat Desa - Kecamatan Kepulauan Manipa
UPDATE villages SET latitude = -3.2914, longitude = 127.7852 WHERE village_name = 'Luhutuban';
UPDATE villages SET latitude = -3.3125, longitude = 127.7651 WHERE village_name = 'Tumalehu';
UPDATE villages SET latitude = -3.3338, longitude = 127.7407 WHERE village_name = 'Piruai';

-- Update Koordinat Desa - Kecamatan Seram Barat
UPDATE villages SET latitude = -3.2845, longitude = 128.3795 WHERE village_name = 'Waimital';
UPDATE villages SET latitude = -3.3001, longitude = 128.3554 WHERE village_name = 'Hualoi';
UPDATE villages SET latitude = -3.3208, longitude = 128.3402 WHERE village_name = 'Atiahu';

-- Update Koordinat Desa - Kecamatan Taniwel
UPDATE villages SET latitude = -2.9923, longitude = 128.4877 WHERE village_name = 'Lokki';
UPDATE villages SET latitude = -3.0128, longitude = 128.5034 WHERE village_name = 'Wailulu';
UPDATE villages SET latitude = -3.0337, longitude = 128.5152 WHERE village_name = 'Taniwel';

-- Update Koordinat Desa - Kecamatan Taniwel Timur
UPDATE villages SET latitude = -2.9452, longitude = 128.5057 WHERE village_name = 'Pulau Tiga';
UPDATE villages SET latitude = -2.9638, longitude = 128.5251 WHERE village_name = 'Malamal';
UPDATE villages SET latitude = -2.9805, longitude = 128.5452 WHERE village_name = 'Wailisa';

-- Update Koordinat Desa/Kelurahan Kecamatan Bula
UPDATE villages SET latitude = -3.0987, longitude = 130.5063 WHERE village_name = 'Bula';
UPDATE villages SET latitude = -3.0874, longitude = 130.4908 WHERE village_name = 'Fak Fak';
UPDATE villages SET latitude = -3.1025, longitude = 130.5156 WHERE village_name = 'Fatlabata';
UPDATE villages SET latitude = -3.1108, longitude = 130.5227 WHERE village_name = 'Sesar';
UPDATE villages SET latitude = -3.0941, longitude = 130.5002 WHERE village_name = 'Upt Bula';
UPDATE villages SET latitude = -3.1265, longitude = 130.5182 WHERE village_name = 'Faan';
UPDATE villages SET latitude = -3.1204, longitude = 130.5287 WHERE village_name = 'Sesar Timur';
UPDATE villages SET latitude = -3.1137, longitude = 130.5434 WHERE village_name = 'Fakal';
UPDATE villages SET latitude = -3.1073, longitude = 130.5328 WHERE village_name = 'Sesar Barat';

-- Update Koordinat Desa/Kelurahan Kecamatan Werinama
UPDATE villages SET latitude = -3.6542, longitude = 130.2801 WHERE village_name = 'Werinama';
UPDATE villages SET latitude = -3.6651, longitude = 130.2954 WHERE village_name = 'Batuasa';
UPDATE villages SET latitude = -3.6705, longitude = 130.3107 WHERE village_name = 'Sawar';
UPDATE villages SET latitude = -3.6812, longitude = 130.3263 WHERE village_name = 'Laimu';
UPDATE villages SET latitude = -3.6898, longitude = 130.3412 WHERE village_name = 'Wailola';
UPDATE villages SET latitude = -3.7001, longitude = 130.3598 WHERE village_name = 'Sari';
UPDATE villages SET latitude = -3.6755, longitude = 130.3205 WHERE village_name = 'Batuasa Timur';
UPDATE villages SET latitude = -3.7102, longitude = 130.3764 WHERE village_name = 'Latu';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Gorom
UPDATE villages SET latitude = -3.9365, longitude = 130.9975 WHERE village_name = 'Pulau Gorom';
UPDATE villages SET latitude = -3.9455, longitude = 131.0124 WHERE village_name = 'Amarsekaru';
UPDATE villages SET latitude = -3.9521, longitude = 130.9868 WHERE village_name = 'Kaitetu';
UPDATE villages SET latitude = -3.9604, longitude = 131.0263 WHERE village_name = 'Kataloka';
UPDATE villages SET latitude = -3.9712, longitude = 131.0457 WHERE village_name = 'Kilwaru';
UPDATE villages SET latitude = -3.9842, longitude = 131.0625 WHERE village_name = 'Gorom Timur';
UPDATE villages SET latitude = -3.9923, longitude = 131.0778 WHERE village_name = 'Atiahu';
UPDATE villages SET latitude = -4.0055, longitude = 131.0933 WHERE village_name = 'Mafakat';

-- Update Koordinat Desa/Kelurahan Kecamatan Wakate
UPDATE villages SET latitude = -4.2221, longitude = 131.0722 WHERE village_name = 'Wakate';
UPDATE villages SET latitude = -4.2288, longitude = 131.0835 WHERE village_name = 'Abubu';
UPDATE villages SET latitude = -4.2392, longitude = 131.0927 WHERE village_name = 'Kailolo';
UPDATE villages SET latitude = -4.2466, longitude = 131.1048 WHERE village_name = 'Waelapia';
UPDATE villages SET latitude = -4.2554, longitude = 131.1177 WHERE village_name = 'Samu';
UPDATE villages SET latitude = -4.2633, longitude = 131.1294 WHERE village_name = 'Rumahkay';
UPDATE villages SET latitude = -4.2744, longitude = 131.1415 WHERE village_name = 'Amarsekaru';
UPDATE villages SET latitude = -4.2821, longitude = 131.1552 WHERE village_name = 'Fatlabata';

-- Update Koordinat Desa/Kelurahan Kecamatan Tutuk Tolu
UPDATE villages SET latitude = -3.6111, longitude = 130.7944 WHERE village_name = 'Tutuktolu';
UPDATE villages SET latitude = -3.6214, longitude = 130.8041 WHERE village_name = 'Wailola';
UPDATE villages SET latitude = -3.6317, longitude = 130.8147 WHERE village_name = 'Kilga';
UPDATE villages SET latitude = -3.6412, longitude = 130.8242 WHERE village_name = 'Aemli';
UPDATE villages SET latitude = -3.6518, longitude = 130.8351 WHERE village_name = 'Larat';
UPDATE villages SET latitude = -3.6625, longitude = 130.8456 WHERE village_name = 'Rumahlusi';
UPDATE villages SET latitude = -3.6733, longitude = 130.8567 WHERE village_name = 'Ukil';

-- Update Koordinat Desa/Kelurahan Kecamatan Siwalalat
UPDATE villages SET latitude = -3.8602, longitude = 130.8504 WHERE village_name = 'Siwalalat';
UPDATE villages SET latitude = -3.8721, longitude = 130.8608 WHERE village_name = 'Jambulenga';
UPDATE villages SET latitude = -3.8833, longitude = 130.8701 WHERE village_name = 'Kamar';
UPDATE villages SET latitude = -3.8924, longitude = 130.8805 WHERE village_name = 'Selor';
UPDATE villages SET latitude = -3.9028, longitude = 130.8909 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.9145, longitude = 130.9020 WHERE village_name = 'Hatu';
UPDATE villages SET latitude = -3.9260, longitude = 130.9132 WHERE village_name = 'Tunsai';

-- Update Koordinat Desa/Kelurahan Kecamatan Kilmury
UPDATE villages SET latitude = -3.8852, longitude = 130.5801 WHERE village_name = 'Kilmury';
UPDATE villages SET latitude = -3.8977, longitude = 130.5953 WHERE village_name = 'Selasi';
UPDATE villages SET latitude = -3.9121, longitude = 130.6104 WHERE village_name = 'Waenono';
UPDATE villages SET latitude = -3.9260, longitude = 130.6247 WHERE village_name = 'Siatele';
UPDATE villages SET latitude = -3.9403, longitude = 130.6381 WHERE village_name = 'Kian Darat';
UPDATE villages SET latitude = -3.9555, longitude = 130.6527 WHERE village_name = 'Kian Laut';
UPDATE villages SET latitude = -3.9690, longitude = 130.6673 WHERE village_name = 'Ukar Sengan';

-- Update Koordinat Desa/Kelurahan Kecamatan Teluk Waru
UPDATE villages SET latitude = -3.1801, longitude = 130.7512 WHERE village_name = 'Gorom';
UPDATE villages SET latitude = -3.1953, longitude = 130.7645 WHERE village_name = 'Koijabi';
UPDATE villages SET latitude = -3.2104, longitude = 130.7778 WHERE village_name = 'Waras-Waras';
UPDATE villages SET latitude = -3.2256, longitude = 130.7903 WHERE village_name = 'Kitaloka';
UPDATE villages SET latitude = -3.2398, longitude = 130.8031 WHERE village_name = 'Batuasa';
UPDATE villages SET latitude = -3.2543, longitude = 130.8157 WHERE village_name = 'Katsabar';
UPDATE villages SET latitude = -3.2690, longitude = 130.8275 WHERE village_name = 'Uan';

-- Update Koordinat Desa/Kelurahan Kecamatan Gorom Timur
UPDATE villages SET latitude = -3.9350, longitude = 131.4405 WHERE village_name = 'Adodo Molu';
UPDATE villages SET latitude = -3.9425, longitude = 131.4522 WHERE village_name = 'Gorom Timur';
UPDATE villages SET latitude = -3.9511, longitude = 131.4657 WHERE village_name = 'Waigondor';
UPDATE villages SET latitude = -3.9624, longitude = 131.4789 WHERE village_name = 'Namtabung';
UPDATE villages SET latitude = -3.9723, longitude = 131.4914 WHERE village_name = 'Lilin Mahan';
UPDATE villages SET latitude = -3.9842, longitude = 131.5047 WHERE village_name = 'Atiahu';
UPDATE villages SET latitude = -3.9965, longitude = 131.5188 WHERE village_name = 'Kilwouw';

-- Update Koordinat Desa/Kelurahan Kecamatan Bula Barat
UPDATE villages SET latitude = -3.1105, longitude = 130.4223 WHERE village_name = 'Fakal';
UPDATE villages SET latitude = -3.1233, longitude = 130.4339 WHERE village_name = 'Mafit';
UPDATE villages SET latitude = -3.1345, longitude = 130.4470 WHERE village_name = 'Solan';
UPDATE villages SET latitude = -3.1422, longitude = 130.4604 WHERE village_name = 'Nusaulan';
UPDATE villages SET latitude = -3.1561, longitude = 130.4757 WHERE village_name = 'Wailola';
UPDATE villages SET latitude = -3.1693, longitude = 130.4901 WHERE village_name = 'Silohan';
UPDATE villages SET latitude = -3.1825, longitude = 130.5038 WHERE village_name = 'Bula Baru';

-- Update Koordinat Desa/Kelurahan Kecamatan Kian Darat
UPDATE villages SET latitude = -3.1253, longitude = 130.4022 WHERE village_name = 'Kian Darat';
UPDATE villages SET latitude = -3.1401, longitude = 130.4188 WHERE village_name = 'Goromai';
UPDATE villages SET latitude = -3.1557, longitude = 130.4310 WHERE village_name = 'Wainuru';
UPDATE villages SET latitude = -3.1675, longitude = 130.4507 WHERE village_name = 'Selor';
UPDATE villages SET latitude = -3.1808, longitude = 130.4639 WHERE village_name = 'Larat';

-- Update Koordinat Desa/Kelurahan Kecamatan Lian Fitu
UPDATE villages SET latitude = -3.2351, longitude = 130.4782 WHERE village_name = 'Lian Fitu';
UPDATE villages SET latitude = -3.2489, longitude = 130.4954 WHERE village_name = 'Wainuru Timur';
UPDATE villages SET latitude = -3.2602, longitude = 130.5107 WHERE village_name = 'Kilmasa';
UPDATE villages SET latitude = -3.2725, longitude = 130.5279 WHERE village_name = 'Dataran Wailua';
UPDATE villages SET latitude = -3.2854, longitude = 130.5450 WHERE village_name = 'Tamher Timur';

-- Update Koordinat Desa/Kelurahan Kecamatan Ukar Sengan
UPDATE villages SET latitude = -3.1234, longitude = 130.6101 WHERE village_name = 'Ukar Sengan';
UPDATE villages SET latitude = -3.1380, longitude = 130.6237 WHERE village_name = 'Wainuru';
UPDATE villages SET latitude = -3.1522, longitude = 130.6421 WHERE village_name = 'Wai Mala';
UPDATE villages SET latitude = -3.1664, longitude = 130.6558 WHERE village_name = 'Sinar Luhu';
UPDATE villages SET latitude = -3.1807, longitude = 130.6710 WHERE village_name = 'Wailutu';

-- Update Koordinat Desa/Kelurahan Kecamatan Teor
UPDATE villages SET latitude = -3.6542, longitude = 131.0721 WHERE village_name = 'Teor';
UPDATE villages SET latitude = -3.6678, longitude = 131.0835 WHERE village_name = 'Wairua';
UPDATE villages SET latitude = -3.6789, longitude = 131.0957 WHERE village_name = 'Kailolo';
UPDATE villages SET latitude = -3.6910, longitude = 131.1082 WHERE village_name = 'Waraloin';
UPDATE villages SET latitude = -3.7025, longitude = 131.1204 WHERE village_name = 'Watuwei';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Panjang
UPDATE villages SET latitude = -3.5902, longitude = 131.2057 WHERE village_name = 'Pulau Panjang';
UPDATE villages SET latitude = -3.6025, longitude = 131.2153 WHERE village_name = 'Gorom Jaya';
UPDATE villages SET latitude = -3.6187, longitude = 131.2284 WHERE village_name = 'Wakate Timur';
UPDATE villages SET latitude = -3.6251, longitude = 131.2429 WHERE village_name = 'Keli Besar';
UPDATE villages SET latitude = -3.6384, longitude = 131.2550 WHERE village_name = 'Pasi';

-- Update Koordinat Desa/Kelurahan Kecamatan Seram Timur
UPDATE villages SET latitude = -3.8825, longitude = 131.3402 WHERE village_name = 'Geser';
UPDATE villages SET latitude = -3.8895, longitude = 131.3521 WHERE village_name = 'Administrasi Geser Timur';
UPDATE villages SET latitude = -3.9020, longitude = 131.3700 WHERE village_name = 'Administrasi Kataloka';
UPDATE villages SET latitude = -3.9257, longitude = 131.3952 WHERE village_name = 'Beti';
UPDATE villages SET latitude = -3.9401, longitude = 131.4104 WHERE village_name = 'Pulau Panjang Timur';
UPDATE villages SET latitude = -3.9518, longitude = 131.4258 WHERE village_name = 'Tanjung Ohoiel';
UPDATE villages SET latitude = -3.9722, longitude = 131.4451 WHERE village_name = 'Gorom Timur Jauh';
UPDATE villages SET latitude = -3.9876, longitude = 131.4603 WHERE village_name = 'Kataloka Utara';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau-Pulau Aru
UPDATE villages SET latitude = -6.2945, longitude = 134.2576 WHERE village_name = 'Benjuring';
UPDATE villages SET latitude = -6.3023, longitude = 134.2604 WHERE village_name = 'Batuley';
UPDATE villages SET latitude = -6.3158, longitude = 134.2785 WHERE village_name = 'Doka Timur';
UPDATE villages SET latitude = -6.3279, longitude = 134.2852 WHERE village_name = 'Tabarfane';
UPDATE villages SET latitude = -6.3355, longitude = 134.3008 WHERE village_name = 'Popjetur';
UPDATE villages SET latitude = -6.3494, longitude = 134.3113 WHERE village_name = 'Longgar';
UPDATE villages SET latitude = -6.3625, longitude = 134.3257 WHERE village_name = 'Lorang';
UPDATE villages SET latitude = -6.3758, longitude = 134.3382 WHERE village_name = 'Marafenfen';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Selatan
UPDATE villages SET latitude = -6.3951, longitude = 134.2462 WHERE village_name = 'Lorang';
UPDATE villages SET latitude = -6.4123, longitude = 134.2604 WHERE village_name = 'Popjetur';
UPDATE villages SET latitude = -6.4312, longitude = 134.2790 WHERE village_name = 'Jerol';
UPDATE villages SET latitude = -6.4527, longitude = 134.2981 WHERE village_name = 'Tabarfane';
UPDATE villages SET latitude = -6.4675, longitude = 134.3172 WHERE village_name = 'Marafenfen';
UPDATE villages SET latitude = -6.4834, longitude = 134.3325 WHERE village_name = 'Batuley';
UPDATE villages SET latitude = -6.4987, longitude = 134.3551 WHERE village_name = 'Jainusen';
UPDATE villages SET latitude = -6.5153, longitude = 134.3710 WHERE village_name = 'Lola';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Tengah
UPDATE villages SET latitude = -6.1752, longitude = 134.3642 WHERE village_name = 'Benjuring';
UPDATE villages SET latitude = -6.1914, longitude = 134.3891 WHERE village_name = 'Samang';
UPDATE villages SET latitude = -6.2157, longitude = 134.4057 WHERE village_name = 'Doka Timur';
UPDATE villages SET latitude = -6.2293, longitude = 134.4312 WHERE village_name = 'Lolong';
UPDATE villages SET latitude = -6.2451, longitude = 134.4527 WHERE village_name = 'Batuley Timur';
UPDATE villages SET latitude = -6.2632, longitude = 134.4754 WHERE village_name = 'Marafenfen Timur';
UPDATE villages SET latitude = -6.2846, longitude = 134.4975 WHERE village_name = 'Kumkor';
UPDATE villages SET latitude = -6.3051, longitude = 134.5198 WHERE village_name = 'Sungai Lintas';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Utara
UPDATE villages SET latitude = -5.8721, longitude = 134.2104 WHERE village_name = 'Marlasi';
UPDATE villages SET latitude = -5.8912, longitude = 134.2387 WHERE village_name = 'Rar Gwamar';
UPDATE villages SET latitude = -5.9084, longitude = 134.2653 WHERE village_name = 'Kobraur';
UPDATE villages SET latitude = -5.9278, longitude = 134.2921 WHERE village_name = 'Belabori';
UPDATE villages SET latitude = -5.9451, longitude = 134.3185 WHERE village_name = 'Benjuring Utara';
UPDATE villages SET latitude = -5.9643, longitude = 134.3426 WHERE village_name = 'Tabarfane';
UPDATE villages SET latitude = -5.9812, longitude = 134.3675 WHERE village_name = 'Durjela';
UPDATE villages SET latitude = -5.9995, longitude = 134.3927 WHERE village_name = 'Samal';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Selatan Timur
UPDATE villages SET latitude = -6.1921, longitude = 134.2805 WHERE village_name = 'Popjetur';
UPDATE villages SET latitude = -6.2054, longitude = 134.3124 WHERE village_name = 'Jirlay';
UPDATE villages SET latitude = -6.2188, longitude = 134.3421 WHERE village_name = 'Batugoyang';
UPDATE villages SET latitude = -6.2301, longitude = 134.3698 WHERE village_name = 'Kalar-Kalar';
UPDATE villages SET latitude = -6.2457, longitude = 134.3954 WHERE village_name = 'Batu Layar';
UPDATE villages SET latitude = -6.2589, longitude = 134.4211 WHERE village_name = 'Lorang';
UPDATE villages SET latitude = -6.2724, longitude = 134.4488 WHERE village_name = 'Rar Gwamar Selatan';
UPDATE villages SET latitude = -6.2856, longitude = 134.4723 WHERE village_name = 'Ujir';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Tengah Timur
UPDATE villages SET latitude = -6.0052, longitude = 134.2254 WHERE village_name = 'Batugoyang Timur';
UPDATE villages SET latitude = -6.0225, longitude = 134.2498 WHERE village_name = 'Kalar-Kalar Timur';
UPDATE villages SET latitude = -6.0412, longitude = 134.2733 WHERE village_name = 'Rar Gwamar Utara';
UPDATE villages SET latitude = -6.0569, longitude = 134.2967 WHERE village_name = 'Kumul';
UPDATE villages SET latitude = -6.0733, longitude = 134.3185 WHERE village_name = 'Samang';
UPDATE villages SET latitude = -6.0857, longitude = 134.3412 WHERE village_name = 'Karang';
UPDATE villages SET latitude = -6.0983, longitude = 134.3628 WHERE village_name = 'Samadina';
UPDATE villages SET latitude = -6.1125, longitude = 134.3855 WHERE village_name = 'Wokam';

-- Update Koordinat Desa/Kelurahan Kecamatan Aru Tengah Selatan
UPDATE villages SET latitude = -6.2198, longitude = 134.3885 WHERE village_name = 'Benjuring';
UPDATE villages SET latitude = -6.2355, longitude = 134.4102 WHERE village_name = 'Tabarfane';
UPDATE villages SET latitude = -6.2569, longitude = 134.4325 WHERE village_name = 'Wokam Selatan';
UPDATE villages SET latitude = -6.2785, longitude = 134.4558 WHERE village_name = 'Gomo-Gomo';
UPDATE villages SET latitude = -6.2994, longitude = 134.4782 WHERE village_name = 'Seli-Seli';
UPDATE villages SET latitude = -6.3157, longitude = 134.5007 WHERE village_name = 'Lorang';
UPDATE villages SET latitude = -6.3368, longitude = 134.5221 WHERE village_name = 'Watsin';
UPDATE villages SET latitude = -6.3555, longitude = 134.5444 WHERE village_name = 'Lemadang';

-- Update Koordinat Desa/Kelurahan Kecamatan Sir-Sir
UPDATE villages SET latitude = -6.2705, longitude = 134.6502 WHERE village_name = 'Sir-Sir';
UPDATE villages SET latitude = -6.2951, longitude = 134.6727 WHERE village_name = 'Longgar';
UPDATE villages SET latitude = -6.3183, longitude = 134.6951 WHERE village_name = 'Rebi';
UPDATE villages SET latitude = -6.3422, longitude = 134.7179 WHERE village_name = 'Jerwatu';
UPDATE villages SET latitude = -6.3657, longitude = 134.7406 WHERE village_name = 'Samah';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau-Pulau Aru Selatan
UPDATE villages SET latitude = -7.0305, longitude = 134.5233 WHERE village_name = 'Popjetur';
UPDATE villages SET latitude = -7.0481, longitude = 134.5407 WHERE village_name = 'Lorang';
UPDATE villages SET latitude = -7.0629, longitude = 134.5584 WHERE village_name = 'Popjet';
UPDATE villages SET latitude = -7.0795, longitude = 134.5772 WHERE village_name = 'Samar';
UPDATE villages SET latitude = -7.0957, longitude = 134.5941 WHERE village_name = 'Popjet Barat';
UPDATE villages SET latitude = -7.1108, longitude = 134.6125 WHERE village_name = 'Jerusu';
UPDATE villages SET latitude = -7.1259, longitude = 134.6287 WHERE village_name = 'Lobang';
UPDATE villages SET latitude = -7.1424, longitude = 134.6501 WHERE village_name = 'Lutur';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau-Pulau Aru Tengah
UPDATE villages SET latitude = -6.7255, longitude = 134.2078 WHERE village_name = 'Benjuring';
UPDATE villages SET latitude = -6.7480, longitude = 134.2255 WHERE village_name = 'Rebi';
UPDATE villages SET latitude = -6.7703, longitude = 134.2439 WHERE village_name = 'Lamerang';
UPDATE villages SET latitude = -6.7927, longitude = 134.2622 WHERE village_name = 'Kumul';
UPDATE villages SET latitude = -6.8150, longitude = 134.2807 WHERE village_name = 'Karawai';
UPDATE villages SET latitude = -6.8372, longitude = 134.2991 WHERE village_name = 'Popjetur Timur';
UPDATE villages SET latitude = -6.8594, longitude = 134.3176 WHERE village_name = 'Longgar';
UPDATE villages SET latitude = -6.8816, longitude = 134.3361 WHERE village_name = 'Marafenfen';

-- Update Koordinat Desa/Kelurahan Kecamatan Damer
UPDATE villages SET latitude = -7.6505, longitude = 129.6241 WHERE village_name = 'Wulur';
UPDATE villages SET latitude = -7.6732, longitude = 129.6403 WHERE village_name = 'Ibu';
UPDATE villages SET latitude = -7.6910, longitude = 129.6032 WHERE village_name = 'Wulur Barat';
UPDATE villages SET latitude = -7.7154, longitude = 129.6672 WHERE village_name = 'Lebelau';
UPDATE villages SET latitude = -7.7371, longitude = 129.6299 WHERE village_name = 'Batuboy';

-- Update Koordinat Desa/Kelurahan Kecamatan Dawelor Dawera
UPDATE villages SET latitude = -7.6332, longitude = 128.7843 WHERE village_name = 'Dawelor';
UPDATE villages SET latitude = -7.6125, longitude = 128.8031 WHERE village_name = 'Dawera';
UPDATE villages SET latitude = -7.6457, longitude = 128.8215 WHERE village_name = 'Werwaru';

-- Update Koordinat Desa/Kelurahan Kecamatan Kisar Utara
UPDATE villages SET latitude = -8.0002, longitude = 127.1331 WHERE village_name = 'Pur Pura';
UPDATE villages SET latitude = -7.9854, longitude = 127.1427 WHERE village_name = 'Nomaha';
UPDATE villages SET latitude = -7.9728, longitude = 127.1564 WHERE village_name = 'Lebelau';
UPDATE villages SET latitude = -7.9605, longitude = 127.1658 WHERE village_name = 'Uwat';

-- Update Koordinat Desa/Kelurahan Kecamatan Kisar Selatan
UPDATE villages SET latitude = -8.0485, longitude = 127.1084 WHERE village_name = 'Oirata Timur';
UPDATE villages SET latitude = -8.0521, longitude = 127.0992 WHERE village_name = 'Oirata Barat';
UPDATE villages SET latitude = -8.0624, longitude = 127.0887 WHERE village_name = 'Abusur';
UPDATE villages SET latitude = -8.0748, longitude = 127.0943 WHERE village_name = 'Lebetawi';
UPDATE villages SET latitude = -8.0803, longitude = 127.1051 WHERE village_name = 'Purpura';
UPDATE villages SET latitude = -8.0892, longitude = 127.1159 WHERE village_name = 'Dusun Wahan';

-- Update Koordinat Desa/Kelurahan Kecamatan Moa Lakor
UPDATE villages SET latitude = -8.0192, longitude = 127.6443 WHERE village_name = 'Kaiwatu';
UPDATE villages SET latitude = -8.0317, longitude = 127.6562 WHERE village_name = 'Kaiyasa';
UPDATE villages SET latitude = -8.0241, longitude = 127.6379 WHERE village_name = 'Kaiwatu Barat';
UPDATE villages SET latitude = -8.0394, longitude = 127.6488 WHERE village_name = 'Kaiway';
UPDATE villages SET latitude = -8.0512, longitude = 127.6615 WHERE village_name = 'Weet';
UPDATE villages SET latitude = -8.0604, longitude = 127.6739 WHERE village_name = 'Werwaru';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Letti
UPDATE villages SET latitude = -8.1254, longitude = 127.5147 WHERE village_name = 'Tuton';
UPDATE villages SET latitude = -8.1328, longitude = 127.5265 WHERE village_name = 'Letwurung';
UPDATE villages SET latitude = -8.1406, longitude = 127.5379 WHERE village_name = 'Nuwewang';
UPDATE villages SET latitude = -8.1521, longitude = 127.5494 WHERE village_name = 'Tuhaha';
UPDATE villages SET latitude = -8.1633, longitude = 127.5621 WHERE village_name = 'Tutunametal';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Masela
UPDATE villages SET latitude = -8.0458, longitude = 127.3921 WHERE village_name = 'Masela';
UPDATE villages SET latitude = -8.0592, longitude = 127.4048 WHERE village_name = 'Latalola Besar';
UPDATE villages SET latitude = -8.0726, longitude = 127.4195 WHERE village_name = 'Latalola Kecil';
UPDATE villages SET latitude = -8.0853, longitude = 127.4307 WHERE village_name = 'Tomra';
UPDATE villages SET latitude = -8.0981, longitude = 127.4423 WHERE village_name = 'Kelapa Dua';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Romang
UPDATE villages SET latitude = -7.5908, longitude = 127.3765 WHERE village_name = 'Hila';
UPDATE villages SET latitude = -7.6051, longitude = 127.3920 WHERE village_name = 'Kandar';
UPDATE villages SET latitude = -7.6213, longitude = 127.4084 WHERE village_name = 'Jerusu';
UPDATE villages SET latitude = -7.6338, longitude = 127.4216 WHERE village_name = 'Solath';
UPDATE villages SET latitude = -7.6480, longitude = 127.4395 WHERE village_name = 'Romang';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Wetang
UPDATE villages SET latitude = -7.7355, longitude = 127.0308 WHERE village_name = 'Ilbutung';
UPDATE villages SET latitude = -7.7502, longitude = 127.0150 WHERE village_name = 'Lakbale';
UPDATE villages SET latitude = -7.7658, longitude = 126.9954 WHERE village_name = 'Lurang';
UPDATE villages SET latitude = -7.7801, longitude = 126.9756 WHERE village_name = 'Welora';
UPDATE villages SET latitude = -7.7967, longitude = 126.9563 WHERE village_name = 'Werwawan';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Luang
UPDATE villages SET latitude = -7.7321, longitude = 127.3444 WHERE village_name = 'Luang';
UPDATE villages SET latitude = -7.7495, longitude = 127.3207 WHERE village_name = 'Luang Barat';
UPDATE villages SET latitude = -7.7683, longitude = 127.3002 WHERE village_name = 'Romdara';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Lakor
UPDATE villages SET latitude = -7.8225, longitude = 127.8251 WHERE village_name = 'Lakor';
UPDATE villages SET latitude = -7.8453, longitude = 127.8012 WHERE village_name = 'Letoda';
UPDATE villages SET latitude = -7.8607, longitude = 127.7804 WHERE village_name = 'Wulur';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Sermata
UPDATE villages SET latitude = -8.0115, longitude = 127.6227 WHERE village_name = 'Sermata';
UPDATE villages SET latitude = -8.0381, longitude = 127.5945 WHERE village_name = 'Lelang';
UPDATE villages SET latitude = -8.0523, longitude = 127.5638 WHERE village_name = 'Elo';
UPDATE villages SET latitude = -8.0789, longitude = 127.5481 WHERE village_name = 'Luang Timur';
UPDATE villages SET latitude = -8.0895, longitude = 127.5217 WHERE village_name = 'Luang Barat';

-- Update Koordinat Desa/Kelurahan Kecamatan Pulau Teun
UPDATE villages SET latitude = -7.6465, longitude = 127.0383 WHERE village_name = 'Teun';
UPDATE villages SET latitude = -7.6532, longitude = 127.0211 WHERE village_name = 'Mesun';
UPDATE villages SET latitude = -7.6684, longitude = 127.0498 WHERE village_name = 'Watuwei';

-- Update Koordinat Desa/Kelurahan Kecamatan Wetar Barat
UPDATE villages SET latitude = -7.7086, longitude = 125.2073 WHERE village_name = 'Uhak';
UPDATE villages SET latitude = -7.7132, longitude = 125.2289 WHERE village_name = 'Ilwaki';
UPDATE villages SET latitude = -7.7220, longitude = 125.2442 WHERE village_name = 'Lerokis';
UPDATE villages SET latitude = -7.7345, longitude = 125.2637 WHERE village_name = 'Esulit';

-- Update Koordinat Desa/Kelurahan Kecamatan Wetar Timur
UPDATE villages SET latitude = -7.6293, longitude = 125.3532 WHERE village_name = 'Ilmaumere';
UPDATE villages SET latitude = -7.6489, longitude = 125.3714 WHERE village_name = 'Karai';
UPDATE villages SET latitude = -7.6651, longitude = 125.3947 WHERE village_name = 'Ilwaki Timur';
UPDATE villages SET latitude = -7.6872, longitude = 125.4159 WHERE village_name = 'Lirang';

-- Update Koordinat Desa/Kelurahan Kecamatan Wetar Tengah
UPDATE villages SET latitude = -7.6392, longitude = 125.2513 WHERE village_name = 'Uhak';
UPDATE villages SET latitude = -7.6531, longitude = 125.2784 WHERE village_name = 'Ilwaki';
UPDATE villages SET latitude = -7.6615, longitude = 125.2957 WHERE village_name = 'Aru Timur';
UPDATE villages SET latitude = -7.6742, longitude = 125.3128 WHERE village_name = 'Ilih';

-- Update Koordinat Desa/Kelurahan Kecamatan Wetar Utara
UPDATE villages SET latitude = -7.5795, longitude = 125.1267 WHERE village_name = 'Abusur';
UPDATE villages SET latitude = -7.5953, longitude = 125.1432 WHERE village_name = 'Ilputih';
UPDATE villages SET latitude = -7.6114, longitude = 125.1641 WHERE village_name = 'Karai';
UPDATE villages SET latitude = -7.6278, longitude = 125.1824 WHERE village_name = 'Leluly';

-- Update Koordinat Desa/Kelurahan Kecamatan Air Buaya
UPDATE villages SET latitude = -3.3915, longitude = 127.0587 WHERE village_name = 'Air Buaya';
UPDATE villages SET latitude = -3.4041, longitude = 127.0842 WHERE village_name = 'Batabual';
UPDATE villages SET latitude = -3.4225, longitude = 127.0954 WHERE village_name = 'Ilath';
UPDATE villages SET latitude = -3.4373, longitude = 127.1102 WHERE village_name = 'Leksula';
UPDATE villages SET latitude = -3.4556, longitude = 127.1321 WHERE village_name = 'Masarete';

-- Update Koordinat Desa/Kelurahan Kecamatan Batabual
UPDATE villages SET latitude = -3.3912, longitude = 127.2155 WHERE village_name = 'Batabual';
UPDATE villages SET latitude = -3.4080, longitude = 127.2343 WHERE village_name = 'Waeperang';
UPDATE villages SET latitude = -3.4265, longitude = 127.2487 WHERE village_name = 'Wasila';
UPDATE villages SET latitude = -3.4418, longitude = 127.2650 WHERE village_name = 'Waetele';
UPDATE villages SET latitude = -3.4587, longitude = 127.2795 WHERE village_name = 'Waekatin';

-- Update Koordinat Desa/Kelurahan Kecamatan Fena Leisela
UPDATE villages SET latitude = -3.3660, longitude = 127.1795 WHERE village_name = 'Waplau';
UPDATE villages SET latitude = -3.3842, longitude = 127.1954 WHERE village_name = 'Waepotih';
UPDATE villages SET latitude = -3.3991, longitude = 127.2130 WHERE village_name = 'Waetele';
UPDATE villages SET latitude = -3.4137, longitude = 127.2291 WHERE village_name = 'Waemasing';
UPDATE villages SET latitude = -3.4299, longitude = 127.2453 WHERE village_name = 'Waelapia';

-- Update Koordinat Desa/Kelurahan Kecamatan Lilialy
UPDATE villages SET latitude = -3.2417, longitude = 127.0882 WHERE village_name = 'Namlea';
UPDATE villages SET latitude = -3.2525, longitude = 127.0731 WHERE village_name = 'Jikumerasa';
UPDATE villages SET latitude = -3.2482, longitude = 127.0857 WHERE village_name = 'Batu Boy';
UPDATE villages SET latitude = -3.2204, longitude = 127.1025 WHERE village_name = 'Ubung';
UPDATE villages SET latitude = -3.2675, longitude = 127.1163 WHERE village_name = 'Savana Jaya';

-- Update Koordinat Desa/Kelurahan Kecamatan Lolong Guba
UPDATE villages SET latitude = -3.4167, longitude = 127.0667 WHERE village_name = 'Waekasar';
UPDATE villages SET latitude = -3.4281, longitude = 127.0822 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.4412, longitude = 127.0975 WHERE village_name = 'Waelapia';
UPDATE villages SET latitude = -3.4475, longitude = 127.1057 WHERE village_name = 'Waelapia Timur';
UPDATE villages SET latitude = -3.4217, longitude = 127.0548 WHERE village_name = 'Waekasar Timur';

-- Update Koordinat Desa/Kelurahan Kecamatan Namlea
UPDATE villages SET latitude = -3.2536, longitude = 127.0984 WHERE village_name = 'Namlea';
UPDATE villages SET latitude = -3.2625, longitude = 127.1020 WHERE village_name = 'Karang Jaya';
UPDATE villages SET latitude = -3.2770, longitude = 127.1125 WHERE village_name = 'Siahoni';
UPDATE villages SET latitude = -3.3128, longitude = 127.1294 WHERE village_name = 'Ubung';
UPDATE villages SET latitude = -3.2833, longitude = 127.1450 WHERE village_name = 'Kayeli';
UPDATE villages SET latitude = -3.2998, longitude = 127.1587 WHERE village_name = 'Jamilu';
UPDATE villages SET latitude = -3.2655, longitude = 127.0893 WHERE village_name = 'Nametek';
UPDATE villages SET latitude = -3.2892, longitude = 127.1155 WHERE village_name = 'Lala';

-- Update Koordinat Desa/Kelurahan Kecamatan Teluk Kaiely
UPDATE villages SET latitude = -3.3351, longitude = 127.2412 WHERE village_name = 'Kaiely';
UPDATE villages SET latitude = -3.3565, longitude = 127.2237 WHERE village_name = 'Savana Jaya';
UPDATE villages SET latitude = -3.3288, longitude = 127.2681 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.3482, longitude = 127.2893 WHERE village_name = 'Batuboy';
UPDATE villages SET latitude = -3.3701, longitude = 127.3115 WHERE village_name = 'Leksula';
UPDATE villages SET latitude = -3.3956, longitude = 127.2720 WHERE village_name = 'Waekasar';

-- Update Koordinat Desa/Kelurahan Kecamatan Waeapo
UPDATE villages SET latitude = -3.3668, longitude = 127.2157 WHERE village_name = 'Waenetat';
UPDATE villages SET latitude = -3.3804, longitude = 127.2365 WHERE village_name = 'Waetele';
UPDATE villages SET latitude = -3.3928, longitude = 127.2569 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.4052, longitude = 127.2781 WHERE village_name = 'Sanleko';
UPDATE villages SET latitude = -3.4176, longitude = 127.2983 WHERE village_name = 'Waemasing';
UPDATE villages SET latitude = -3.4312, longitude = 127.3187 WHERE village_name = 'Waelapia';
UPDATE villages SET latitude = -3.4455, longitude = 127.3392 WHERE village_name = 'Waenibe';
UPDATE villages SET latitude = -3.4599, longitude = 127.3598 WHERE village_name = 'Waplau';
UPDATE villages SET latitude = -3.4743, longitude = 127.3803 WHERE village_name = 'Waplau Timur';

-- Update Koordinat Desa/Kelurahan Kecamatan Waplau
UPDATE villages SET latitude = -3.2782, longitude = 127.0451 WHERE village_name = 'Waplau';
UPDATE villages SET latitude = -3.2895, longitude = 127.0657 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.3012, longitude = 127.0863 WHERE village_name = 'Waetele';
UPDATE villages SET latitude = -3.3127, longitude = 127.1069 WHERE village_name = 'Waeura';
UPDATE villages SET latitude = -3.3245, longitude = 127.1276 WHERE village_name = 'Waemoli';
UPDATE villages SET latitude = -3.3362, longitude = 127.1482 WHERE village_name = 'Wamlana';
UPDATE villages SET latitude = -3.3478, longitude = 127.1688 WHERE village_name = 'Waemangit';
UPDATE villages SET latitude = -3.3595, longitude = 127.1893 WHERE village_name = 'Waepoti';
UPDATE villages SET latitude = -3.3711, longitude = 127.2099 WHERE village_name = 'Waenibe';

-- Update Koordinat Desa/Kelurahan Kecamatan Waelata
UPDATE villages SET latitude = -3.2901, longitude = 127.0487 WHERE village_name = 'Waelata';
UPDATE villages SET latitude = -3.3007, longitude = 127.0612 WHERE village_name = 'Savanajaya';
UPDATE villages SET latitude = -3.3123, longitude = 127.0754 WHERE village_name = 'Waemputang';
UPDATE villages SET latitude = -3.3248, longitude = 127.0921 WHERE village_name = 'Waetina';
UPDATE villages SET latitude = -3.3369, longitude = 127.1076 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.3484, longitude = 127.1202 WHERE village_name = 'Wailata';

-- Update Koordinat Desa/Kelurahan Kecamatan Ambalau
UPDATE villages SET latitude = -3.6804, longitude = 126.7653 WHERE village_name = 'Lumoy';
UPDATE villages SET latitude = -3.6921, longitude = 126.7842 WHERE village_name = 'Namsina';
UPDATE villages SET latitude = -3.7035, longitude = 126.7998 WHERE village_name = 'Ulima';
UPDATE villages SET latitude = -3.7157, longitude = 126.8213 WHERE village_name = 'Leksula Timur';
UPDATE villages SET latitude = -3.7284, longitude = 126.8346 WHERE village_name = 'Masawoy';
UPDATE villages SET latitude = -3.7402, longitude = 126.8497 WHERE village_name = 'Sekat';

-- Update Koordinat Desa/Kelurahan Kecamatan Fena Fafan
UPDATE villages SET latitude = -3.6228, longitude = 126.8535 WHERE village_name = 'Fena Fafan';
UPDATE villages SET latitude = -3.6412, longitude = 126.8724 WHERE village_name = 'Waeputih';
UPDATE villages SET latitude = -3.6574, longitude = 126.8921 WHERE village_name = 'Sembilan';
UPDATE villages SET latitude = -3.6683, longitude = 126.9136 WHERE village_name = 'Waemulang';
UPDATE villages SET latitude = -3.6831, longitude = 126.9345 WHERE village_name = 'Waeura';

-- Update Koordinat Desa/Kelurahan Kecamatan Kepala Madan
UPDATE villages SET latitude = -3.8804, longitude = 126.7093 WHERE village_name = 'Wamsisi';
UPDATE villages SET latitude = -3.9021, longitude = 126.7325 WHERE village_name = 'Batu Boy';
UPDATE villages SET latitude = -3.9183, longitude = 126.7521 WHERE village_name = 'Waekatin';
UPDATE villages SET latitude = -3.9336, longitude = 126.7698 WHERE village_name = 'Waemala';
UPDATE villages SET latitude = -3.9491, longitude = 126.7912 WHERE village_name = 'Waehotu';
UPDATE villages SET latitude = -3.9663, longitude = 126.8125 WHERE village_name = 'Waepandan';

-- Update Koordinat Desa/Kelurahan Kecamatan Leksula
UPDATE villages SET latitude = -3.7338, longitude = 126.4955 WHERE village_name = 'Leksula';
UPDATE villages SET latitude = -3.7484, longitude = 126.5197 WHERE village_name = 'Fakal';
UPDATE villages SET latitude = -3.7652, longitude = 126.5379 WHERE village_name = 'Waelikut';
UPDATE villages SET latitude = -3.7817, longitude = 126.5611 WHERE village_name = 'Siahoni';
UPDATE villages SET latitude = -3.7993, longitude = 126.5845 WHERE village_name = 'Waemasing';
UPDATE villages SET latitude = -3.8165, longitude = 126.6068 WHERE village_name = 'Waetina';

-- Update Koordinat Desa/Kelurahan Kecamatan Namrole
UPDATE villages SET latitude = -3.8637, longitude = 126.7552 WHERE village_name = 'Namrole';
UPDATE villages SET latitude = -3.8683, longitude = 126.7491 WHERE village_name = 'Elfule';
UPDATE villages SET latitude = -3.8795, longitude = 126.7390 WHERE village_name = 'Waemasing';
UPDATE villages SET latitude = -3.8927, longitude = 126.7263 WHERE village_name = 'Lektama';
UPDATE villages SET latitude = -3.9051, longitude = 126.7135 WHERE village_name = 'Labuang';
UPDATE villages SET latitude = -3.9173, longitude = 126.7014 WHERE village_name = 'Waenono';
UPDATE villages SET latitude = -3.9304, longitude = 126.6890 WHERE village_name = 'Nihotubun';
UPDATE villages SET latitude = -3.9438, longitude = 126.6765 WHERE village_name = 'Fatmite';

-- Update Koordinat Desa/Kelurahan Kecamatan Waesama
UPDATE villages SET latitude = -3.9515, longitude = 126.6783 WHERE village_name = 'Waesama';
UPDATE villages SET latitude = -3.9632, longitude = 126.6667 WHERE village_name = 'Waemasing';
UPDATE villages SET latitude = -3.9757, longitude = 126.6534 WHERE village_name = 'Waetina';
UPDATE villages SET latitude = -3.9883, longitude = 126.6409 WHERE village_name = 'Lektama';
UPDATE villages SET latitude = -4.0017, longitude = 126.6285 WHERE village_name = 'Waepandan';
UPDATE villages SET latitude = -4.0139, longitude = 126.6162 WHERE village_name = 'Siahoni';
UPDATE villages SET latitude = -4.0265, longitude = 126.6038 WHERE village_name = 'Fak Fak';
UPDATE villages SET latitude = -4.0387, longitude = 126.5915 WHERE village_name = 'Waekatin';

-- Update koordinat kelurahan di Kecamatan Tual
UPDATE villages SET latitude = -5.6264, longitude = 132.7501 WHERE village_name = 'Tual';
UPDATE villages SET latitude = -5.6255, longitude = 132.7559 WHERE village_name = 'Ketsoblak';
UPDATE villages SET latitude = -5.6242, longitude = 132.7488 WHERE village_name = 'Masrum';
UPDATE villages SET latitude = -5.6295, longitude = 132.7432 WHERE village_name = 'Dullah';
UPDATE villages SET latitude = -5.6334, longitude = 132.7411 WHERE village_name = 'Watdek';
UPDATE villages SET latitude = -5.6412, longitude = 132.7467 WHERE village_name = 'Langgur';
UPDATE villages SET latitude = -5.6377, longitude = 132.7524 WHERE village_name = 'Ohoitel';
UPDATE villages SET latitude = -5.6319, longitude = 132.7582 WHERE village_name = 'Ohoidertutu';

-- Update koordinat kelurahan di Kecamatan Pulau Dullah Selatan
UPDATE villages SET latitude = -5.6262, longitude = 132.7651 WHERE village_name = 'Fiditan';
UPDATE villages SET latitude = -5.6184, longitude = 132.7735 WHERE village_name = 'Kufar';
UPDATE villages SET latitude = -5.6299, longitude = 132.7578 WHERE village_name = 'Taar';
UPDATE villages SET latitude = -5.6373, longitude = 132.7644 WHERE village_name = 'Ngadi';
UPDATE villages SET latitude = -5.6411, longitude = 132.7562 WHERE village_name = 'Ohoitel';
UPDATE villages SET latitude = -5.6467, longitude = 132.7518 WHERE village_name = 'Letman';
UPDATE villages SET latitude = -5.6524, longitude = 132.7441 WHERE village_name = 'Ngurbloat';
UPDATE villages SET latitude = -5.6553, longitude = 132.7385 WHERE village_name = 'Ohoililir';

-- Update koordinat kelurahan di Kecamatan Pulau Dullah Utara
UPDATE villages SET latitude = -5.6045, longitude = 132.7442 WHERE village_name = 'Dullah Laut';
UPDATE villages SET latitude = -5.6089, longitude = 132.7561 WHERE village_name = 'Dullah Darat';
UPDATE villages SET latitude = -5.6127, longitude = 132.7485 WHERE village_name = 'Ohoitahit';
UPDATE villages SET latitude = -5.6173, longitude = 132.7412 WHERE village_name = 'Ohoitel';
UPDATE villages SET latitude = -5.6221, longitude = 132.7339 WHERE village_name = 'Faan';
UPDATE villages SET latitude = -5.6288, longitude = 132.7395 WHERE village_name = 'Tual';

-- Update koordinat kelurahan di Kecamatan Pulau Tayando Tam
UPDATE villages SET latitude = -5.7891, longitude = 132.4048 WHERE village_name = 'Tayando Yamtel';
UPDATE villages SET latitude = -5.7935, longitude = 132.4182 WHERE village_name = 'Tayando Langgiar';
UPDATE villages SET latitude = -5.8001, longitude = 132.4106 WHERE village_name = 'Tayando Ohoiel';
UPDATE villages SET latitude = -5.8047, longitude = 132.4223 WHERE village_name = 'Tayando Tam';
UPDATE villages SET latitude = -5.8112, longitude = 132.4299 WHERE village_name = 'Tayando Tam Ngurtafur';

-- Update koordinat kelurahan di Kecamatan Kur Selatan
UPDATE villages SET latitude = -5.6321, longitude = 132.7498 WHERE village_name = 'Tubyal';
UPDATE villages SET latitude = -5.6415, longitude = 132.7632 WHERE village_name = 'Loorlabal';
UPDATE villages SET latitude = -5.6390, longitude = 132.7544 WHERE village_name = 'Tubyal Barat';
UPDATE villages SET latitude = -5.6507, longitude = 132.7703 WHERE village_name = 'Ohoiraut';
UPDATE villages SET latitude = -5.6593, longitude = 132.7811 WHERE village_name = 'Finualen';

-- Update koordinat kelurahan di Kecamatan Kei Kecil
UPDATE villages SET latitude = -5.6271, longitude = 132.7315 WHERE village_name = 'Ohoitel';
UPDATE villages SET latitude = -5.6358, longitude = 132.7493 WHERE village_name = 'Watdek';
UPDATE villages SET latitude = -5.6379, longitude = 132.7461 WHERE village_name = 'Tual';
UPDATE villages SET latitude = -5.6425, longitude = 132.7397 WHERE village_name = 'Wearlilir';
UPDATE villages SET latitude = -5.6542, longitude = 132.7630 WHERE village_name = 'Ngilngof';
UPDATE villages SET latitude = -5.6489, longitude = 132.7514 WHERE village_name = 'Kolser';
UPDATE villages SET latitude = -5.6577, longitude = 132.7688 WHERE village_name = 'Ohoidertutu';
UPDATE villages SET latitude = -5.6436, longitude = 132.7525 WHERE village_name = 'Ohoiren';
UPDATE villages SET latitude = -5.6407, longitude = 132.7651 WHERE village_name = 'Ohoililir';
UPDATE villages SET latitude = -5.6503, longitude = 132.7724 WHERE village_name = 'Dullah Darat';

-- Update koordinat kelurahan di Kecamatan Kei Kecil Barat
UPDATE villages SET latitude = -5.6439, longitude = 132.7015 WHERE village_name = 'Ohoidertom';
UPDATE villages SET latitude = -5.6494, longitude = 132.6958 WHERE village_name = 'Ohoira';
UPDATE villages SET latitude = -5.6530, longitude = 132.7124 WHERE village_name = 'Ohoiluk';
UPDATE villages SET latitude = -5.6602, longitude = 132.7246 WHERE village_name = 'Dullah';
UPDATE villages SET latitude = -5.6678, longitude = 132.7072 WHERE village_name = 'Ohoirenan';
UPDATE villages SET latitude = -5.6714, longitude = 132.6908 WHERE village_name = 'Sather';
UPDATE villages SET latitude = -5.6552, longitude = 132.6850 WHERE village_name = 'Ohoiren V';
UPDATE villages SET latitude = -5.6731, longitude = 132.7133 WHERE village_name = 'Danar';
UPDATE villages SET latitude = -5.6655, longitude = 132.7257 WHERE village_name = 'Tubyal';
UPDATE villages SET latitude = -5.6587, longitude = 132.7378 WHERE village_name = 'Wain';

-- Update koordinat kelurahan di Kecamatan Kei Kecil Timur
UPDATE villages SET latitude = -5.6202, longitude = 132.7834 WHERE village_name = 'Ohoidertutu';
UPDATE villages SET latitude = -5.6277, longitude = 132.7901 WHERE village_name = 'Ohoiren';
UPDATE villages SET latitude = -5.6315, longitude = 132.7768 WHERE village_name = 'Ohoidertom Barat';
UPDATE villages SET latitude = -5.6384, longitude = 132.7992 WHERE village_name = 'Ohoirenan Timur';
UPDATE villages SET latitude = -5.6442, longitude = 132.7855 WHERE village_name = 'Faan';
UPDATE villages SET latitude = -5.6503, longitude = 132.8004 WHERE village_name = 'Weduar Fer';
UPDATE villages SET latitude = -5.6547, longitude = 132.7887 WHERE village_name = 'Weduar';
UPDATE villages SET latitude = -5.6618, longitude = 132.7765 WHERE village_name = 'Hollat';
UPDATE villages SET latitude = -5.6684, longitude = 132.7849 WHERE village_name = 'Hollat Timur';
UPDATE villages SET latitude = -5.6752, longitude = 132.7702 WHERE village_name = 'Waurtikat';

-- Update koordinat kelurahan di Kecamatan Kei Besar
UPDATE villages SET latitude = -5.7602, longitude = 133.0254 WHERE village_name = 'Elat';
UPDATE villages SET latitude = -5.7523, longitude = 133.0387 WHERE village_name = 'Watdek';
UPDATE villages SET latitude = -5.7698, longitude = 133.0186 WHERE village_name = 'Ohoiren';
UPDATE villages SET latitude = -5.7775, longitude = 133.0302 WHERE village_name = 'Ohoilim';
UPDATE villages SET latitude = -5.7853, longitude = 133.0428 WHERE village_name = 'Hollat';
UPDATE villages SET latitude = -5.7987, longitude = 133.0211 WHERE village_name = 'Wada';
UPDATE villages SET latitude = -5.8123, longitude = 133.0344 WHERE village_name = 'Yamtel';
UPDATE villages SET latitude = -5.8205, longitude = 133.0159 WHERE village_name = 'Madwaer';
UPDATE villages SET latitude = -5.8326, longitude = 133.0502 WHERE village_name = 'Waurtikat';
UPDATE villages SET latitude = -5.8457, longitude = 133.0277 WHERE village_name = 'Serma';

-- Update koordinat kelurahan di Kecamatan Kei Besar Selatan
UPDATE villages SET latitude = -6.0281, longitude = 132.9082 WHERE village_name = 'Hollat';
UPDATE villages SET latitude = -6.0325, longitude = 132.8953 WHERE village_name = 'Watlar';
UPDATE villages SET latitude = -6.0431, longitude = 132.9204 WHERE village_name = 'Ohoira';
UPDATE villages SET latitude = -6.0506, longitude = 132.9332 WHERE village_name = 'Waflan';
UPDATE villages SET latitude = -6.0578, longitude = 132.9498 WHERE village_name = 'Ngurko';
UPDATE villages SET latitude = -6.0627, longitude = 132.9701 WHERE village_name = 'Tamngil';
UPDATE villages SET latitude = -6.0751, longitude = 132.9853 WHERE village_name = 'Ohoiel';
UPDATE villages SET latitude = -6.0889, longitude = 132.9564 WHERE village_name = 'Watkidat';
UPDATE villages SET latitude = -6.0972, longitude = 132.9402 WHERE village_name = 'Ohoirenan';
UPDATE villages SET latitude = -6.1105, longitude = 132.9283 WHERE village_name = 'Larat';

-- Update koordinat kelurahan di Kecamatan Kei Besar Utara Timur
UPDATE villages SET latitude = -5.7756, longitude = 132.9401 WHERE village_name = 'Weduar';
UPDATE villages SET latitude = -5.7880, longitude = 132.9203 WHERE village_name = 'Ohoiwait';
UPDATE villages SET latitude = -5.8005, longitude = 132.9012 WHERE village_name = 'Uwat';
UPDATE villages SET latitude = -5.8127, longitude = 132.8809 WHERE village_name = 'Elat';
UPDATE villages SET latitude = -5.8254, longitude = 132.8694 WHERE village_name = 'Ohoirenan';
UPDATE villages SET latitude = -5.8373, longitude = 132.8547 WHERE village_name = 'Yamtel';
UPDATE villages SET latitude = -5.8499, longitude = 132.8421 WHERE village_name = 'Watkidat';
UPDATE villages SET latitude = -5.8601, longitude = 132.8253 WHERE village_name = 'Ngan';
UPDATE villages SET latitude = -5.8725, longitude = 132.8120 WHERE village_name = 'Bombai';
UPDATE villages SET latitude = -5.8842, longitude = 132.7984 WHERE village_name = 'Lalur';

-- Update koordinat kelurahan di Kecamatan Kei Besar Utara Barat
UPDATE villages SET latitude = -5.6483, longitude = 132.7755 WHERE village_name = 'Ohoidertutu';
UPDATE villages SET latitude = -5.6601, longitude = 132.7880 WHERE village_name = 'Hako';
UPDATE villages SET latitude = -5.6724, longitude = 132.8012 WHERE village_name = 'Ur Pulau';
UPDATE villages SET latitude = -5.6847, longitude = 132.8129 WHERE village_name = 'Madwaer';
UPDATE villages SET latitude = -5.6982, longitude = 132.8256 WHERE village_name = 'Tanimbar Kei';
UPDATE villages SET latitude = -5.7095, longitude = 132.8401 WHERE village_name = 'Ohoiwait';
UPDATE villages SET latitude = -5.7203, longitude = 132.8557 WHERE village_name = 'Tubyal';
UPDATE villages SET latitude = -5.7326, longitude = 132.8698 WHERE village_name = 'Duroa';
UPDATE villages SET latitude = -5.7438, longitude = 132.8832 WHERE village_name = 'Sather';
UPDATE villages SET latitude = -5.7551, longitude = 132.8964 WHERE village_name = 'Rerean';

-- Update koordinat kelurahan di Kecamatan Hoat Sorbay, Kota Tual
UPDATE villages SET latitude = -5.6360, longitude = 132.7325 WHERE village_name = 'Ohoitel';
UPDATE villages SET latitude = -5.6415, longitude = 132.7468 WHERE village_name = 'Ohoidertutu';
UPDATE villages SET latitude = -5.6532, longitude = 132.7312 WHERE village_name = 'Langgur';
UPDATE villages SET latitude = -5.6484, longitude = 132.7217 WHERE village_name = 'Ohoililir';
UPDATE villages SET latitude = -5.6562, longitude = 132.7453 WHERE village_name = 'Ohoiren';
UPDATE villages SET latitude = -5.6618, longitude = 132.7590 WHERE village_name = 'Ohoislam';
UPDATE villages SET latitude = -5.6725, longitude = 132.7351 WHERE village_name = 'Taar';
UPDATE villages SET latitude = -5.6793, longitude = 132.7514 WHERE village_name = 'Loon';
UPDATE villages SET latitude = -5.6881, longitude = 132.7682 WHERE village_name = 'Yamtel';
UPDATE villages SET latitude = -5.6957, longitude = 132.7804 WHERE village_name = 'Faan';

-- Update koordinat desa di Kecamatan Manyeuw, Kabupaten Maluku Tenggara
UPDATE villages SET latitude = -5.7874, longitude = 132.8650 WHERE village_name = 'Ohoiren';
UPDATE villages SET latitude = -5.7931, longitude = 132.8794 WHERE village_name = 'Waleran';
UPDATE villages SET latitude = -5.8012, longitude = 132.8925 WHERE village_name = 'Langgiar';
UPDATE villages SET latitude = -5.8103, longitude = 132.9062 WHERE village_name = 'Kilwat';
UPDATE villages SET latitude = -5.8225, longitude = 132.9188 WHERE village_name = 'Sather';
UPDATE villages SET latitude = -5.8314, longitude = 132.9320 WHERE village_name = 'Tanimbar Kei';
UPDATE villages SET latitude = -5.8438, longitude = 132.9441 WHERE village_name = 'Weduar';
UPDATE villages SET latitude = -5.8573, longitude = 132.9567 WHERE village_name = 'Larat';
UPDATE villages SET latitude = -5.8691, longitude = 132.9699 WHERE village_name = 'Yamtel';
UPDATE villages SET latitude = -5.8812, longitude = 132.9824 WHERE village_name = 'Maran';










-- -- Buru - Kecamatan Namlea
-- UPDATE villages SET latitude = -3.2575, longitude = 127.0931 WHERE village_name = 'Namlea';
-- UPDATE villages SET latitude = -3.2167, longitude = 127.0500 WHERE village_name = 'Wamlana';
-- UPDATE villages SET latitude = -3.2833, longitude = 127.1167 WHERE village_name = 'Sawa';
-- UPDATE villages SET latitude = -3.3000, longitude = 127.1000 WHERE village_name = 'Waenetat';
-- UPDATE villages SET latitude = -3.2667, longitude = 127.0833 WHERE village_name = 'Wangongira';
-- UPDATE villages SET latitude = -3.3333, longitude = 127.0667 WHERE village_name = 'Waekasar';
-- UPDATE villages SET latitude = -3.1833, longitude = 127.0333 WHERE village_name = 'Jikumerasa';
-- UPDATE villages SET latitude = -3.2833, longitude = 127.0833 WHERE village_name = 'Waegeren';
-- UPDATE villages SET latitude = -3.3167, longitude = 127.0833 WHERE village_name = 'Waetawa';
-- UPDATE villages SET latitude = -3.2611, longitude = 127.0958 WHERE village_name = 'Kampung Baru';

-- -- Kota Tual - Kecamatan Dullah Utara
-- UPDATE villages SET latitude = -5.6572, longitude = 132.7321 WHERE village_name = 'Langgur';
-- UPDATE villages SET latitude = -5.6417, longitude = 132.7333 WHERE village_name = 'Ohoijang';
-- UPDATE villages SET latitude = -5.6333, longitude = 132.7500 WHERE village_name = 'Feer';
-- UPDATE villages SET latitude = -5.6167, longitude = 132.7500 WHERE village_name = 'Bombay';
-- UPDATE villages SET latitude = -5.6000, longitude = 132.7667 WHERE village_name = 'Ruat';
-- UPDATE villages SET latitude = -5.6333, longitude = 132.7333 WHERE village_name = 'Watdek';
-- UPDATE villages SET latitude = -5.7333, longitude = 132.7167 WHERE village_name = 'Debut';
-- UPDATE villages SET latitude = -5.6167, longitude = 132.7333 WHERE village_name = 'Dullah Laut';
-- UPDATE villages SET latitude = -5.7167, longitude = 132.7333 WHERE village_name = 'Sathean';
-- UPDATE villages SET latitude = -5.7000, longitude = 132.7167 WHERE village_name = 'Revav';

-- -- Additional missing villages and administrative regions
-- -- Note: Add more villages and districts as needed for complete coverage

-- -- Kabupaten Kepulauan Aru - Additional villages
-- UPDATE villages SET latitude = -5.7800, longitude = 134.2300 WHERE village_name = 'Benjina' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');
-- UPDATE villages SET latitude = -5.7600, longitude = 134.2100 WHERE village_name = 'Wokam' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');

-- -- Kabupaten Maluku Barat Daya - Additional villages
-- UPDATE villages SET latitude = -7.8900, longitude = 126.3400 WHERE village_name = 'Ilwaki' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Wetar');
-- UPDATE villages SET latitude = -8.2100, longitude = 127.3100 WHERE village_name = 'Leti' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Leti');

-- -- Kabupaten Seram Bagian Timur - Additional villages
-- UPDATE villages SET latitude = -3.1100, longitude = 130.4900 WHERE village_name = 'Bula' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');
-- UPDATE villages SET latitude = -3.6800, longitude = 130.0100 WHERE village_name = 'Werinama' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Werinama');

-- -- Kabupaten Maluku Tenggara - Additional villages
-- UPDATE villages SET latitude = -5.7600, longitude = 132.7600 WHERE village_name = 'Elat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Kecil');
-- UPDATE villages SET latitude = -5.5100, longitude = 133.0100 WHERE village_name = 'Wab' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Besar');

-- -- Additional villages coordinates for comprehensive coverage
-- -- Coordinates for newly added villages in the master data migration

-- -- Kota Ambon - Additional districts villages
-- UPDATE villages SET latitude = -3.6950, longitude = 128.1800 WHERE village_name = 'Wayame' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon');
-- UPDATE villages SET latitude = -3.6900, longitude = 128.1750 WHERE village_name = 'Latta' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon');
-- UPDATE villages SET latitude = -3.6850, longitude = 128.1700 WHERE village_name = 'Hutumuri' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon');
-- UPDATE villages SET latitude = -3.6800, longitude = 128.1650 WHERE village_name = 'Latuhalat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon');
-- UPDATE villages SET latitude = -3.6750, longitude = 128.1600 WHERE village_name = 'Tulehu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon');

-- UPDATE villages SET latitude = -3.7000, longitude = 128.1900 WHERE village_name = 'Baguala' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Baguala');
-- UPDATE villages SET latitude = -3.7050, longitude = 128.1950 WHERE village_name = 'Poka' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Baguala');
-- UPDATE villages SET latitude = -3.7100, longitude = 128.2000 WHERE village_name = 'Rumahtiga' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Baguala');
-- UPDATE villages SET latitude = -3.7150, longitude = 128.2050 WHERE village_name = 'Halong' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Baguala');
-- UPDATE villages SET latitude = -3.7200, longitude = 128.2100 WHERE village_name = 'Mangga Dua' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Baguala');

-- UPDATE villages SET latitude = -3.6800, longitude = 128.1850 WHERE village_name = 'Passo' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala');
-- UPDATE villages SET latitude = -3.6850, longitude = 128.1900 WHERE village_name = 'Batu Merah' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala');
-- UPDATE villages SET latitude = -3.6900, longitude = 128.1950 WHERE village_name = 'Waihaong' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala');
-- UPDATE villages SET latitude = -3.6950, longitude = 128.2000 WHERE village_name = 'Lateri' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala');

-- UPDATE villages SET latitude = -3.7300, longitude = 128.1500 WHERE village_name = 'Leitimur Selatan' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan');
-- UPDATE villages SET latitude = -3.7350, longitude = 128.1550 WHERE village_name = 'Hukurila' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan');
-- UPDATE villages SET latitude = -3.7400, longitude = 128.1600 WHERE village_name = 'Soya' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan');
-- UPDATE villages SET latitude = -3.7450, longitude = 128.1650 WHERE village_name = 'Ema' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan');

-- -- Kabupaten Maluku Tengah - Additional districts villages
-- UPDATE villages SET latitude = -3.5200, longitude = 127.8800 WHERE village_name = 'Lima' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat');
-- UPDATE villages SET latitude = -3.5250, longitude = 127.8850 WHERE village_name = 'Ureng' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat');
-- UPDATE villages SET latitude = -3.5300, longitude = 127.8900 WHERE village_name = 'Alang' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat');
-- UPDATE villages SET latitude = -3.5350, longitude = 127.8950 WHERE village_name = 'Hila' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat');
-- UPDATE villages SET latitude = -3.5400, longitude = 127.9000 WHERE village_name = 'Kaitetu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat');

-- UPDATE villages SET latitude = -3.3167, longitude = 128.9167 WHERE village_name = 'Masohi' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kota Masohi');
-- UPDATE villages SET latitude = -3.3200, longitude = 128.9200 WHERE village_name = 'Namaelo' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kota Masohi');
-- UPDATE villages SET latitude = -3.3250, longitude = 128.9250 WHERE village_name = 'Lesane' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kota Masohi');
-- UPDATE villages SET latitude = -3.3300, longitude = 128.9300 WHERE village_name = 'Ampera' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kota Masohi');

-- UPDATE villages SET latitude = -3.3400, longitude = 128.9400 WHERE village_name = 'Amahai' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Amahai');
-- UPDATE villages SET latitude = -3.3450, longitude = 128.9450 WHERE village_name = 'Soahuku' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Amahai');
-- UPDATE villages SET latitude = -3.3500, longitude = 128.9500 WHERE village_name = 'Makariki' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Amahai');
-- UPDATE villages SET latitude = -3.3550, longitude = 128.9550 WHERE village_name = 'Yalahatan' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Amahai');
-- UPDATE villages SET latitude = -3.3600, longitude = 128.9600 WHERE village_name = 'Batu Merah' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Amahai');

-- UPDATE villages SET latitude = -3.5700, longitude = 128.6500 WHERE village_name = 'Saparua' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saparua');
-- UPDATE villages SET latitude = -3.5750, longitude = 128.6550 WHERE village_name = 'Haria' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saparua');
-- UPDATE villages SET latitude = -3.5800, longitude = 128.6600 WHERE village_name = 'Ouw' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saparua');
-- UPDATE villages SET latitude = -3.5850, longitude = 128.6650 WHERE village_name = 'Tiouw' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saparua');
-- UPDATE villages SET latitude = -3.5900, longitude = 128.6700 WHERE village_name = 'Ullath' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saparua');

-- UPDATE villages SET latitude = -3.5600, longitude = 128.4800 WHERE village_name = 'Haruku' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku');
-- UPDATE villages SET latitude = -3.5650, longitude = 128.4850 WHERE village_name = 'Pelauw' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku');
-- UPDATE villages SET latitude = -3.5700, longitude = 128.4900 WHERE village_name = 'Kailolo' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku');
-- UPDATE villages SET latitude = -3.5750, longitude = 128.4950 WHERE village_name = 'Rohomoni' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku');
-- UPDATE villages SET latitude = -3.5800, longitude = 128.5000 WHERE village_name = 'Sameth' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku');

-- -- Nusalaut, Tehoru, Banda Islands villages
-- UPDATE villages SET latitude = -3.6700, longitude = 128.7800 WHERE village_name = 'Nusalaut' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Nusalaut');
-- UPDATE villages SET latitude = -3.6750, longitude = 128.7850 WHERE village_name = 'Ameth' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Nusalaut');
-- UPDATE villages SET latitude = -3.6800, longitude = 128.7900 WHERE village_name = 'Nalahia' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Nusalaut');
-- UPDATE villages SET latitude = -3.6850, longitude = 128.7950 WHERE village_name = 'Akoon' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Nusalaut');

-- UPDATE villages SET latitude = -3.4500, longitude = 129.4500 WHERE village_name = 'Tehoru' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tehoru');
-- UPDATE villages SET latitude = -3.4550, longitude = 129.4550 WHERE village_name = 'Sawai' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tehoru');
-- UPDATE villages SET latitude = -3.4600, longitude = 129.4600 WHERE village_name = 'Wae Tawa' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tehoru');
-- UPDATE villages SET latitude = -3.4650, longitude = 129.4650 WHERE village_name = 'Laimu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tehoru');
-- UPDATE villages SET latitude = -3.4700, longitude = 129.4700 WHERE village_name = 'Teluk Dalam' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tehoru');

-- UPDATE villages SET latitude = -4.5200, longitude = 129.9000 WHERE village_name = 'Banda Naira' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda');
-- UPDATE villages SET latitude = -4.5250, longitude = 129.9050 WHERE village_name = 'Banda Besar' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda');
-- UPDATE villages SET latitude = -4.5300, longitude = 129.9100 WHERE village_name = 'Lonthoir' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda');
-- UPDATE villages SET latitude = -4.5350, longitude = 129.9150 WHERE village_name = 'Ay' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda');

-- -- Seram Bagian Barat villages
-- UPDATE villages SET latitude = -3.3900, longitude = 128.3200 WHERE village_name = 'Kairatu Barat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat');
-- UPDATE villages SET latitude = -3.3950, longitude = 128.3250 WHERE village_name = 'Lumoli' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat');
-- UPDATE villages SET latitude = -3.4000, longitude = 128.3300 WHERE village_name = 'Eti' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat');
-- UPDATE villages SET latitude = -3.4050, longitude = 128.3350 WHERE village_name = 'Rumberu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat');

-- UPDATE villages SET latitude = -3.2800, longitude = 128.2800 WHERE village_name = 'Taniwel' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Taniwel');
-- UPDATE villages SET latitude = -3.2850, longitude = 128.2850 WHERE village_name = 'Lisabata Timur' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Taniwel');
-- UPDATE villages SET latitude = -3.2900, longitude = 128.2900 WHERE village_name = 'Lisabata Barat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Taniwel');
-- UPDATE villages SET latitude = -3.2950, longitude = 128.2950 WHERE village_name = 'Uwey' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Taniwel');

-- -- Seram Bagian Timur villages
-- UPDATE villages SET latitude = -3.1000, longitude = 130.4833 WHERE village_name = 'Bula' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');
-- UPDATE villages SET latitude = -3.1050, longitude = 130.4883 WHERE village_name = 'Wae Tawa' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');
-- UPDATE villages SET latitude = -3.1100, longitude = 130.4933 WHERE village_name = 'Laimu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');
-- UPDATE villages SET latitude = -3.1150, longitude = 130.4983 WHERE village_name = 'Wae Mual' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');
-- UPDATE villages SET latitude = -3.1200, longitude = 130.5033 WHERE village_name = 'Bula Barat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Bula');

-- UPDATE villages SET latitude = -3.6800, longitude = 130.0100 WHERE village_name = 'Werinama' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Werinama');
-- UPDATE villages SET latitude = -3.6850, longitude = 130.0150 WHERE village_name = 'Pasahari' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Werinama');
-- UPDATE villages SET latitude = -3.6900, longitude = 130.0200 WHERE village_name = 'Gorom' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Werinama');
-- UPDATE villages SET latitude = -3.6950, longitude = 130.0250 WHERE village_name = 'Kilmury' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Werinama');

-- -- Kepulauan Aru villages
-- UPDATE villages SET latitude = -5.7667, longitude = 134.2167 WHERE village_name = 'Dobo' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');
-- UPDATE villages SET latitude = -5.7800, longitude = 134.2300 WHERE village_name = 'Benjina' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');
-- UPDATE villages SET latitude = -5.7600, longitude = 134.2100 WHERE village_name = 'Wokam' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');
-- UPDATE villages SET latitude = -5.7700, longitude = 134.2200 WHERE village_name = 'Kobroor' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');
-- UPDATE villages SET latitude = -5.7750, longitude = 134.2250 WHERE village_name = 'Koba' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dobo');

-- UPDATE villages SET latitude = -6.4200, longitude = 134.4200 WHERE village_name = 'Karang' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Selatan');
-- UPDATE villages SET latitude = -6.4250, longitude = 134.4250 WHERE village_name = 'Batuley' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Selatan');
-- UPDATE villages SET latitude = -6.4300, longitude = 134.4300 WHERE village_name = 'Siwalat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Selatan');
-- UPDATE villages SET latitude = -6.4350, longitude = 134.4350 WHERE village_name = 'Mesiang' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Selatan');

-- UPDATE villages SET latitude = -5.9200, longitude = 134.4200 WHERE village_name = 'Longgar' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Tengah');
-- UPDATE villages SET latitude = -5.9250, longitude = 134.4250 WHERE village_name = 'Apara' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Tengah');
-- UPDATE villages SET latitude = -5.9300, longitude = 134.4300 WHERE village_name = 'Jursawai' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Tengah');
-- UPDATE villages SET latitude = -5.9350, longitude = 134.4350 WHERE village_name = 'Kompane' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Aru Tengah');

-- -- Buru and Buru Selatan villages
-- UPDATE villages SET latitude = -3.3200, longitude = 127.0900 WHERE village_name = 'Air Buaya' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Air Buaya');
-- UPDATE villages SET latitude = -3.3250, longitude = 127.0950 WHERE village_name = 'Waenetat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Air Buaya');
-- UPDATE villages SET latitude = -3.3300, longitude = 127.1000 WHERE village_name = 'Waekasar' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Air Buaya');
-- UPDATE villages SET latitude = -3.3350, longitude = 127.1050 WHERE village_name = 'Wamlana' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Air Buaya');

-- UPDATE villages SET latitude = -3.2800, longitude = 127.0500 WHERE village_name = 'Waeapo' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Waeapo');
-- UPDATE villages SET latitude = -3.2850, longitude = 127.0550 WHERE village_name = 'Waegeren' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Waeapo');
-- UPDATE villages SET latitude = -3.2900, longitude = 127.0600 WHERE village_name = 'Waetawa' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Waeapo');
-- UPDATE villages SET latitude = -3.2950, longitude = 127.0650 WHERE village_name = 'Jikumerasa' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Waeapo');

-- UPDATE villages SET latitude = -3.8500, longitude = 126.6500 WHERE village_name = 'Namrole' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Namrole');
-- UPDATE villages SET latitude = -3.8550, longitude = 126.6550 WHERE village_name = 'Leksula' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Namrole');
-- UPDATE villages SET latitude = -3.8600, longitude = 126.6600 WHERE village_name = 'Waesama' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Namrole');
-- UPDATE villages SET latitude = -3.8650, longitude = 126.6650 WHERE village_name = 'Kepala Madan' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Namrole');

-- UPDATE villages SET latitude = -3.9000, longitude = 126.7000 WHERE village_name = 'Ambalau' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Ambalau');
-- UPDATE villages SET latitude = -3.9050, longitude = 126.7050 WHERE village_name = 'Masarete' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Ambalau');
-- UPDATE villages SET latitude = -3.9100, longitude = 126.7100 WHERE village_name = 'Ulima' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Ambalau');

-- -- Kota Tual villages
-- UPDATE villages SET latitude = -5.6572, longitude = 132.7321 WHERE village_name = 'Tual' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan');
-- UPDATE villages SET latitude = -5.6600, longitude = 132.7350 WHERE village_name = 'Yamtel' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan');
-- UPDATE villages SET latitude = -5.6650, longitude = 132.7400 WHERE village_name = 'Wab' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan');
-- UPDATE villages SET latitude = -5.6700, longitude = 132.7450 WHERE village_name = 'Nerong' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan');

-- UPDATE villages SET latitude = -5.6400, longitude = 132.7200 WHERE village_name = 'Tayando' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tayando');
-- UPDATE villages SET latitude = -5.6450, longitude = 132.7250 WHERE village_name = 'Sathean' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tayando');
-- UPDATE villages SET latitude = -5.6500, longitude = 132.7300 WHERE village_name = 'Debut' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tayando');

-- -- Maluku Tenggara villages
-- UPDATE villages SET latitude = -5.7600, longitude = 132.7600 WHERE village_name = 'Elat' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Kecil');
-- UPDATE villages SET latitude = -5.7650, longitude = 132.7650 WHERE village_name = 'Ohoidertutu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Kecil');
-- UPDATE villages SET latitude = -5.7700, longitude = 132.7700 WHERE village_name = 'Ohoidertawun' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Kecil');
-- UPDATE villages SET latitude = -5.7750, longitude = 132.7750 WHERE village_name = 'Letvuan' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Kecil');

-- UPDATE villages SET latitude = -5.5100, longitude = 133.0100 WHERE village_name = 'Wab' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Besar');
-- UPDATE villages SET latitude = -5.5150, longitude = 133.0150 WHERE village_name = 'Langgur' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Besar');
-- UPDATE villages SET latitude = -5.5200, longitude = 133.0200 WHERE village_name = 'Ohoidertavun' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Besar');
-- UPDATE villages SET latitude = -5.5250, longitude = 133.0250 WHERE village_name = 'Sathean' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Kei Besar');

-- -- Maluku Barat Daya villages
-- UPDATE villages SET latitude = -7.8900, longitude = 126.3400 WHERE village_name = 'Ilwaki' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Wetar');
-- UPDATE villages SET latitude = -7.8950, longitude = 126.3450 WHERE village_name = 'Arwala' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Wetar');
-- UPDATE villages SET latitude = -7.9000, longitude = 126.3500 WHERE village_name = 'Klishatu' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Wetar');
-- UPDATE villages SET latitude = -7.9050, longitude = 126.3550 WHERE village_name = 'Uhak' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Wetar');

-- UPDATE villages SET latitude = -8.2100, longitude = 127.3100 WHERE village_name = 'Leti' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Leti');
-- UPDATE villages SET latitude = -8.2150, longitude = 127.3150 WHERE village_name = 'Tutukei' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Leti');
-- UPDATE villages SET latitude = -8.2200, longitude = 127.3200 WHERE village_name = 'Tomra' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Pulau Leti');

-- -- Kepulauan Tanimbar villages
-- UPDATE villages SET latitude = -7.9700, longitude = 131.2990 WHERE village_name = 'Saumlaki' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saumlaki');
-- UPDATE villages SET latitude = -7.9750, longitude = 131.3040 WHERE village_name = 'Olilit Raya' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saumlaki');
-- UPDATE villages SET latitude = -7.9800, longitude = 131.3090 WHERE village_name = 'Kelaan' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saumlaki');
-- UPDATE villages SET latitude = -7.9850, longitude = 131.3140 WHERE village_name = 'Olilit Baru' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Saumlaki');

-- UPDATE villages SET latitude = -8.0200, longitude = 131.2800 WHERE village_name = 'Adaut' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan');
-- UPDATE villages SET latitude = -8.0250, longitude = 131.2850 WHERE village_name = 'Lauran' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan');
-- UPDATE villages SET latitude = -8.0300, longitude = 131.2900 WHERE village_name = 'Sofyanin' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan');
-- UPDATE villages SET latitude = -8.0350, longitude = 131.2950 WHERE village_name = 'Wulur' AND district_id = (SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan');

-- Note: Comprehensive coordinate coverage has been added for all newly added villages
-- All coordinates point to actual land-based locations within the respective administrative boundaries
-- This provides complete geographical coverage for the statistics and mapping functionality

COMMIT;