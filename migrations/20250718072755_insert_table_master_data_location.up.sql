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
((SELECT province_id FROM provinces WHERE province_name = 'Maluku'), 'Kabupaten Maluku Tenggara', 'MAT', 'kabupaten', NOW()),
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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wertamrian', 'WTR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Kepulauan Tanimbar' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara', 'TNU', NOW());

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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Bula', 'BLA', NOW()),
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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Banda', 'BDA', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Seram Bagian Timur' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tehoru', 'THO', NOW());

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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Selatan', 'TMS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara', 'TMU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Selaru', 'SLR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Tamrian', 'WTM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Maktian', 'WMK', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nirunmas', 'NRM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Molu Maru', 'MLM', NOW()),
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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Barat Daya' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara', 'TNB', NOW());

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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kepala Madan', 'KPM', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Buru Selatan' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Waesama', 'WSA', NOW()),
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
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Selatan', 'KBS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil Barat', 'KKB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Utara Timur', 'KBT', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Hoat Sorbay', 'HSB', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Manyeuw', 'MYW', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Besar Utara', 'KBU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kei Kecil Timur Selatan', 'KTS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Pulau Kur', 'PKR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kur Selatan', 'KUR', NOW());

-- Kabupaten Maluku Tenggara Barat
INSERT INTO districts (regency_id, district_name, district_code, created_at) VALUES
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Selatan', 'TNS', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Tanimbar Utara', 'TNU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Selaru', 'SLU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Tamrian', 'WTN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Wer Maktian', 'WMN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Nirunmas', 'NRN', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Molu Maru', 'MLR', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Yaru', 'YRU', NOW()),
((SELECT regency_id FROM regencies WHERE regency_name = 'Kabupaten Maluku Tenggara Barat' AND province_id = (SELECT province_id FROM provinces WHERE province_name = 'Maluku')), 'Kormomolin', 'KML', NOW()),
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

COMMIT;
