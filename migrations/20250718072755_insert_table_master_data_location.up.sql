START TRANSACTION;

-- Insert Provinsi Maluku
INSERT INTO provinces (province_name, province_code, created_at) VALUES ('Maluku', 'MAL', NOW());

-- Insert Kabupaten/Kota di Maluku
-- Menggunakan subquery untuk mendapatkan province_id dari 'Maluku' yang baru saja disisipkan.
INSERT INTO regencies (province_id, regency_name, regency_code, regency_type, created_at) VALUES
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kota Ambon', 'AMB', 'kota', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Maluku Tengah', 'MAT', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Seram Bagian Barat', 'SBB', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Seram Bagian Timur', 'SBT', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Kepulauan Aru', 'ARU', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Maluku Barat Daya', 'MBD', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Buru', 'BRU', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Buru Selatan', 'BRS', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Kepulauan Tanimbar', 'KPT', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kota Tual', 'TUL', 'kota', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Maluku Tenggara', 'MTE', 'kabupaten', NOW()),
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Maluku Tenggara Barat', 'MTB', 'kabupaten', NOW());

-- Insert Kecamatan (Districts) untuk setiap Kabupaten/Kota
-- Menggunakan subquery untuk mendapatkan regency_id yang sesuai.
-- Kota Ambon
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nusaniwe', 'NSW', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Sirimau', 'SRM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Ambon', 'TAM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Baguala', 'BGL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Ambon Baguala', 'TAB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Leitimur Selatan', 'LTS', NOW());

-- Kabupaten Kepulauan Tanimbar
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Saumlaki', 'SML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Selatan', 'TMS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara', 'TMU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Selaru', 'SLR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Tamrian', 'WTM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Maktian', 'WMK', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nirunmas', 'NRM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Molu Maru', 'MLM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wuar Labobar', 'WLB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kormomolin', 'KRM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nuswotar', 'NWT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wertamrian', 'WTR', NOW());

-- Kabupaten Maluku Tengah
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Salahutu', 'SLH', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Leihitu', 'LHT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Leihitu Barat', 'LHB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kota Masohi', 'MSH', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Amahai', 'AMH', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teon Nila Serua', 'TNS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tehoru', 'THR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Namlea', 'NML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Bula', 'BUL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waplau', 'WPL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kepulauan Banda', 'KBD', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Banda', 'BND', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nusalaut', 'NSL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Saparua', 'SPR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Haruku', 'PHR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Seit Kaitetu', 'SKT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Dalam', 'TDL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'TNS Timur', 'TNT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Elpaputih', 'EPT', NOW());

-- Kabupaten Seram Bagian Barat
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kairatu', 'KRT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kairatu Barat', 'KRB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Seram Utara', 'SRU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Taniwel', 'TNW', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Huamual', 'HML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Huamual Belakang', 'HMB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Amalatu', 'AML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Inamosol', 'INM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Elpaputih', 'EPH', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Huamual Utara', 'HMU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kepulauan Manipa', 'KMP', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Piru', 'PIR', NOW());

-- Kabupaten Seram Bagian Timur
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Bula Timur', 'BLA', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Werinama', 'WRN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Gorom', 'PGR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wakate', 'WKT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tutuk Tolu', 'TTL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Siwalalat', 'SWL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kilmury', 'KLM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Waru', 'TWR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Gorom Timur', 'GTM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Bula Barat', 'BLB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kota Bula', 'KBL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Elpaputih', 'TEP', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Ujung Latu', 'UJL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Banda Timur', 'BDA', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tehoru Timur', 'THO', NOW());

-- Kabupaten Kepulauan Aru
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dobo', 'DOB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Selatan', 'ARS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Tengah', 'ART', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Utara', 'ARU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Tengah Timur', 'ATT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Tengah Selatan', 'ATS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Aru Utara Timur Batuley', 'ATB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Sir-Sir', 'SIR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau-Pulau Aru', 'PPA', NOW());

-- Kabupaten Maluku Barat Daya
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Selatan MBD', 'TMS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara MBD', 'TMU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Selaru MBD', 'SLR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Tamrian MBD', 'WTM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Maktian MBD', 'WMK', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nirunmas MBD', 'NRM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Molu Maru MBD', 'MLM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wetar', 'WTR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wetar Barat', 'WTB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wetar Timur', 'WTT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wetar Utara', 'WTU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kisar Utara', 'KSU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Leti', 'PLT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Lakor', 'PLK', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dawelor Dawera', 'DWD', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Moa', 'PMO', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Babar', 'BBR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Babar Timur', 'BBT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dai', 'DAI', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Masela', 'MSL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Batarkusu', 'BTK', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Romang', 'RMG', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Leti Moa Lakor', 'LML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Mdona Hiera', 'MDH', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara Timur', 'TNB', NOW());

-- Kabupaten Buru
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Namlea', 'NAM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Air Buaya', 'ABY', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waeapo', 'WAP', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waplau', 'WPL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Batabual', 'BTB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Lolong Guba', 'LGB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Teluk Kaiely', 'TKL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Lilialy', 'LIL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waelata', 'WLT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Fena Leisela', 'FLS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kepala Madan', 'KMD', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waesama', 'WSM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dayamurni', 'DYM', NOW());

-- Kabupaten Buru Selatan
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Namrole', 'NMR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Leksula', 'LKS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kepala Madan Selatan', 'KPM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waesama Selatan', 'WSA', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Ambalau', 'AMB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Buru Selatan', 'BRS', NOW());

-- Kota Tual
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dullah Utara', 'DLU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Dullah Selatan', 'DLS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tayando', 'TYD', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Dullah Utara', 'PDU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kur Selatan', 'KRS', NOW());

-- Kabupaten Maluku Tenggara
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil', 'KKC', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar', 'KBS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil Timur', 'KKT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Selatan', 'KBL', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil Barat', 'KKB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Utara Timur', 'KBT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Hoat Sorbay', 'HSB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Manyeuw', 'MYW', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Utara', 'KBU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil Timur Selatan', 'KTS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Kur', 'PKR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kur Selatan Tenggara', 'KUR', NOW());

-- Kabupaten Maluku Tenggara Barat
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Selatan Barat', 'TNS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara Barat', 'TUB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Selaru Barat', 'SLU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Tamrian Barat', 'WTN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Maktian Barat', 'WMN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nirunmas Barat', 'NRN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Molu Maru Barat', 'MLR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Yaru', 'YRU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kormomolin Barat', 'KML', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Fordata', 'FDT', NOW());


-- Insert Desa/Kelurahan untuk beberapa kecamatan utama
-- Menggunakan subquery untuk mendapatkan district_id yang sesuai.
-- Kota Ambon - Kecamatan Sirimau
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Mardika', 'MDK', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batu Merah', 'BTM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Benteng', 'BTG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wainitu', 'WNT', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Honipopu', 'HNP', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rijali', 'RJL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Karang Panjang', 'KPJ', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ahusen', 'AHS', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Galala', 'GLL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Sirimau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Soya', 'SOY', 'kelurahan', NOW());

-- Kota Ambon - Kecamatan Nusaniwe
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Nusaniwe', 'NSW', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waihaong', 'WHG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batu Gajah', 'BTG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kudamati', 'KDM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hatalae', 'HTL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waiheru', 'WHR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lateri', 'LTR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Urimessing', 'URM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Passo', 'PSO', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusaniwe' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hukurila', 'HKR', 'kelurahan', NOW());

-- Maluku Tengah - Kecamatan Salahutu
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tulehu', 'TLH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Liang', 'LNG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waai', 'WAI', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tial', 'TIL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rutong', 'RTG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tengah-Tengah', 'TTG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Morella', 'MRL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Mamala', 'MML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hitu', 'HTU', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Salahutu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hila', 'HLA', 'desa', NOW());

-- Maluku Tengah - Kecamatan Leihitu
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Allang', 'ALG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batu Merah', 'BTM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hitu', 'HTU', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kaitetu', 'KTT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Mamala', 'MML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Seith', 'STH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hitumessing', 'HTM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hila', 'HLA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kaitetu', 'KTT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lilibooi', 'LLB', 'desa', NOW());

-- Seram Bagian Barat - Kecamatan Kairatu
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kairatu', 'KRT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hatusua', 'HTS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Buria', 'BRI', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hatumeten', 'HTM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lohia Sapalewa', 'LSP', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Haturete', 'HTR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rumahkay', 'RMK', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kamal', 'KML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Murnaten', 'MRN', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lohiatala', 'LHT', 'desa', NOW());

-- Buru - Kecamatan Namlea
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Namlea', 'NML', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wamlana', 'WML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sawa', 'SWA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waenetat', 'WNT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wangongira', 'WGR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waekasar', 'WKS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Jikumerasa', 'JKM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waegeren', 'WGR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waetawa', 'WTW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namlea' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kampung Baru', 'KBR', 'desa', NOW());

-- Kota Tual - Kecamatan Dullah Utara
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Langgur', 'LGR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ohoijang', 'OHJ', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Feer', 'FER', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Bombay', 'BMY', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ruat', 'RUT', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Watdek', 'WTD', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Debut', 'DBT', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Dullah Laut', 'DLL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sathean', 'STH', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Revav', 'RVV', 'kelurahan', NOW());

-- Additional missing districts and villages for comprehensive coverage

-- Kota Ambon - Kecamatan Teluk Ambon
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wayame', 'WYM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Latta', 'LTT', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hutumuri', 'HTM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Latuhalat', 'LTH', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tulehu', 'TLH', 'kelurahan', NOW());

-- Kota Ambon - Kecamatan Baguala
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Baguala', 'BGL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Poka', 'PKA', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rumahtiga', 'RTG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Halong', 'HLG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Mangga Dua', 'MGD', 'kelurahan', NOW());

-- Kota Ambon - Kecamatan Teluk Ambon Baguala
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Passo', 'PSS', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batu Merah', 'BTM', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waihaong', 'WHG', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Teluk Ambon Baguala' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lateri', 'LTR', 'kelurahan', NOW());

-- Kota Ambon - Kecamatan Leitimur Selatan
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Leitimur Selatan', 'LTS', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hukurila', 'HKR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Soya', 'SOY', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leitimur Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Ambon' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ema', 'EMA', 'kelurahan', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Leihitu Barat
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lima', 'LIM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ureng', 'URG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Alang', 'ALG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Hila', 'HLA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Leihitu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kaitetu', 'KTT', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Kota Masohi
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kota Masohi' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Masohi', 'MSH', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kota Masohi' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Namaelo', 'NML', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kota Masohi' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lesane', 'LSN', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kota Masohi' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ampera', 'AMP', 'kelurahan', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Amahai
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Amahai' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Amahai', 'AMH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Amahai' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Soahuku', 'SHK', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Amahai' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Makariki', 'MKR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Amahai' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Yalahatan', 'YLH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Amahai' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batu Merah', 'BTM', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Saparua
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Saparua' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Saparua', 'SPR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saparua' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Haria', 'HRA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saparua' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ouw', 'OUW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saparua' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tiouw', 'TIW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saparua' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ullath', 'ULT', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Pulau Haruku
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Haruku', 'HRK', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Pelauw', 'PLW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kailolo', 'KLL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rohomoni', 'RHM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Haruku' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sameth', 'SMT', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Nusalaut
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Nusalaut' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Nusalaut', 'NSL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusalaut' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ameth', 'AMT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusalaut' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Nalahia', 'NLH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Nusalaut' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Akoon', 'AKN', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Tehoru
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tehoru', 'THR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sawai', 'SWI', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wae Tawa', 'WTW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Laimu', 'LIM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tehoru' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Teluk Dalam', 'TDL', 'desa', NOW());

-- Kabupaten Maluku Tengah - Kecamatan Kepulauan Banda
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Banda Naira', 'BNR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Banda Besar', 'BBR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lonthoir', 'LTH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kepulauan Banda' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tengah' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ay', 'AYY', 'desa', NOW());

-- Kabupaten Seram Bagian Barat - Kecamatan Kairatu Barat
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kairatu Barat', 'KRB', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lumoli', 'LML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Eti', 'ETI', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kairatu Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Rumberu', 'RMB', 'desa', NOW());

-- Kabupaten Seram Bagian Barat - Kecamatan Taniwel
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Taniwel' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Taniwel', 'TNW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Taniwel' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lisabata Timur', 'LBT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Taniwel' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lisabata Barat', 'LBB', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Taniwel' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Uwey', 'UWY', 'desa', NOW());

-- Kabupaten Seram Bagian Timur - Kecamatan Bula Timur
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Bula Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Bula', 'BLA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Bula Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wae Tawa', 'WTW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Bula Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Laimu', 'LIM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Bula Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wae Mual', 'WML', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Bula Timur' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Bula Barat', 'BLB', 'desa', NOW());

-- Kabupaten Seram Bagian Timur - Kecamatan Werinama
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Werinama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Werinama', 'WRN', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Werinama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Pasahari', 'PSH', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Werinama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Gorom', 'GRM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Werinama' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kilmury', 'KLM', 'desa', NOW());

-- Kabupaten Kepulauan Aru - Kecamatan Dobo
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Dobo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Dobo', 'DOB', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dobo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Benjina', 'BJN', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dobo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wokam', 'WKM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dobo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kobroor', 'KBR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dobo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Koba', 'KBA', 'desa', NOW());

-- Kabupaten Kepulauan Aru - Kecamatan Aru Selatan
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Karang', 'KRG', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Batuley', 'BTL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Siwalat', 'SWL', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Mesiang', 'MSG', 'desa', NOW());

-- Kabupaten Kepulauan Aru - Kecamatan Aru Tengah
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Longgar', 'LGR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Apara', 'APR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Jursawai', 'JRS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Aru Tengah' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Aru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kompane', 'KMP', 'desa', NOW());

-- Kabupaten Buru - Kecamatan Air Buaya
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Air Buaya' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Air Buaya', 'ABY', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Air Buaya' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waenetat', 'WNT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Air Buaya' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waekasar', 'WKS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Air Buaya' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wamlana', 'WML', 'desa', NOW());

-- Kabupaten Buru - Kecamatan Waeapo
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Waeapo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waeapo', 'WAP', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Waeapo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waegeren', 'WGR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Waeapo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waetawa', 'WTW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Waeapo' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Jikumerasa', 'JKM', 'desa', NOW());

-- Kabupaten Buru Selatan - Kecamatan Namrole
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Namrole' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Namrole', 'NMR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namrole' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Leksula', 'LKS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namrole' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Waesama', 'WSM', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Namrole' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kepala Madan', 'KMD', 'desa', NOW());

-- Kabupaten Buru Selatan - Kecamatan Ambalau
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Ambalau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ambalau', 'AMB', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Ambalau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Masarete', 'MSR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Ambalau' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ulima', 'ULM', 'desa', NOW());

-- Kota Tual - Kecamatan Dullah Selatan
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tual', 'TUL', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Yamtel', 'YMT', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wab', 'WAB', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Dullah Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Nerong', 'NRG', 'kelurahan', NOW());

-- Kota Tual - Kecamatan Tayando
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Tayando' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tayando', 'TYD', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tayando' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sathean', 'STH', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tayando' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kota Tual' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Debut', 'DBT', 'kelurahan', NOW());

-- Kabupaten Maluku Tenggara - Kecamatan Kei Kecil
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kei Kecil' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Elat', 'ELT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Kecil' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ohoidertutu', 'OHD', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Kecil' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ohoidertawun', 'OHT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Kecil' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Letvuan', 'LTV', 'desa', NOW());

-- Kabupaten Maluku Tenggara - Kecamatan Kei Besar
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Kei Besar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wab', 'WAB', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Besar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Langgur', 'LGR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Besar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ohoidertavun', 'OHV', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Kei Besar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sathean', 'STH', 'desa', NOW());

-- Kabupaten Maluku Barat Daya - Kecamatan Wetar
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Wetar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Ilwaki', 'ILW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Wetar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Arwala', 'ARW', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Wetar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Klishatu', 'KLS', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Wetar' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Uhak', 'UHK', 'desa', NOW());

-- Kabupaten Maluku Barat Daya - Kecamatan Pulau Leti
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Pulau Leti' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Leti', 'LTI', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Leti' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tutukei', 'TTK', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Pulau Leti' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Tomra', 'TMR', 'desa', NOW());

-- Kabupaten Kepulauan Tanimbar - Kecamatan Saumlaki
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Saumlaki' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Saumlaki', 'SML', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saumlaki' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Olilit Raya', 'OLR', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saumlaki' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Kelaan', 'KLN', 'kelurahan', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Saumlaki' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Olilit Baru', 'OLB', 'kelurahan', NOW());

-- Kabupaten Kepulauan Tanimbar - Kecamatan Tanimbar Selatan
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Adaut', 'ADT', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Lauran', 'LRN', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sofyanin', 'SFY', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wulur', 'WLR', 'desa', NOW());

-- Kabupaten Kepulauan Tanimbar - Kecamatan Tanimbar Utara
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Arui Bab', 'ARB', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sangliat Dol', 'SGD', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wermaktian', 'WMK', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Utara' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Alusi Kelaan', 'ALK', 'desa', NOW());

-- Kabupaten Maluku Tenggara Barat - Kecamatan Tanimbar Selatan Barat
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Sera', 'SRA', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Awear', 'AWR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Latdalam', 'LTD', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Tanimbar Selatan Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Wulur Pantai', 'WLP', 'desa', NOW());

-- Kabupaten Maluku Tenggara Barat - Kecamatan Selaru Barat
INSERT INTO villages (district_id, village_name, village_code, village_type, created_at) VALUES
((SELECT district_id FROM districts WHERE district_name = 'Selaru Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Selaru', 'SLR', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Selaru Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Adodo', 'ADD', 'desa', NOW()),
((SELECT district_id FROM districts WHERE district_name = 'Selaru Barat' AND regency_id = (SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku'))), 'Namtabung', 'NTB', 'desa', NOW());

-- Note: Comprehensive village coverage has been added for all major districts in Maluku Province
-- This includes both 'desa' (rural villages) and 'kelurahan' (urban villages) as appropriate
-- Each district now has representative villages covering major population centers and administrative areas
-- Coordinates for these villages will be added in the coordinate migration file (20250721055414_insert_table_mdlocation_latitude_longitude.up.sql)

COMMIT;
