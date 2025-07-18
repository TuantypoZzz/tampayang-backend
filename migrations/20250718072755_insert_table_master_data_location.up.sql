-- Insert Provinsi Maluku
INSERT INTO provinces (province_name, province_code) VALUES ('Maluku', 'MAL');

-- Insert Kabupaten/Kota di Maluku
INSERT INTO regencies (province_id, regencies_name, regencies_code, regencies_type) VALUES 
(1, 'Kota Ambon', 'AMB', 'kota'),
(1, 'Kabupaten Maluku Tengah', 'MAT', 'kabupaten'),
(1, 'Kabupaten Seram Bagian Barat', 'SBB', 'kabupaten'),
(1, 'Kabupaten Seram Bagian Timur', 'SBT', 'kabupaten'),
(1, 'Kabupaten Kepulauan Aru', 'ARU', 'kabupaten'),
(1, 'Kabupaten Maluku Barat Daya', 'MBD', 'kabupaten'),
(1, 'Kabupaten Buru', 'BRU', 'kabupaten'),
(1, 'Kabupaten Buru Selatan', 'BRS', 'kabupaten'),
(1, 'Kabupaten Kepulauan Tanimbar', 'KPT', 'kabupaten'),
(1, 'Kota Tual', 'TUL', 'kota'),
(1, 'Kabupaten Maluku Tenggara', 'MAT', 'kabupaten'),
(1, 'Kabupaten Maluku Tenggara Barat', 'MTB', 'kabupaten');

-- Insert Kecamatan (Districts) untuk setiap Kabupaten/Kota
-- Kota Ambon (regencies_id = 1)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(1, 'Nusaniwe', 'NSW'),
(1, 'Sirimau', 'SRM'),
(1, 'Teluk Ambon', 'TAM'),
(1, 'Baguala', 'BGL'),
(1, 'Teluk Ambon Baguala', 'TAB'),
(1, 'Leitimur Selatan', 'LTS');
-- Kabupaten Kepulauan Tanimbar (regencies_id = 9)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(9, 'Saumlaki', 'SML'),
(9, 'Tanimbar Selatan', 'TMS'),
(9, 'Tanimbar Utara', 'TMU'),
(9, 'Selaru', 'SLR'),
(9, 'Wer Tamrian', 'WTM'),
(9, 'Wer Maktian', 'WMK'),
(9, 'Nirunmas', 'NRM'),
(9, 'Molu Maru', 'MLM'),
(9, 'Wuar Labobar', 'WLB'),
(9, 'Kormomolin', 'KRM'),
(9, 'Nuswotar', 'NWT'),
(9, 'Wertamrian', 'WTR'),
(9, 'Tanimbar Utara', 'TNU');

-- Kabupaten Maluku Tengah (regencies_id = 2)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(2, 'Salahutu', 'SLH'),
(2, 'Leihitu', 'LHT'),
(2, 'Leihitu Barat', 'LHB'),
(2, 'Kota Masohi', 'MSH'),
(2, 'Amahai', 'AMH'),
(2, 'Teon Nila Serua', 'TNS'),
(2, 'Tehoru', 'THR'),
(2, 'Namlea', 'NML'),
(2, 'Bula', 'BUL'),
(2, 'Waplau', 'WPL'),
(2, 'Kepulauan Banda', 'KBD'),
(2, 'Banda', 'BND'),
(2, 'Nusalaut', 'NSL'),
(2, 'Saparua', 'SPR'),
(2, 'Pulau Haruku', 'PHR'),
(2, 'Seit Kaitetu', 'SKT'),
(2, 'Teluk Dalam', 'TDL'),
(2, 'TNS Timur', 'TNT'),
(2, 'Elpaputih', 'EPT');

-- Kabupaten Seram Bagian Barat (regencies_id = 3)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(3, 'Kairatu', 'KRT'),
(3, 'Kairatu Barat', 'KRB'),
(3, 'Seram Utara', 'SRU'),
(3, 'Taniwel', 'TNW'),
(3, 'Huamual', 'HML'),
(3, 'Huamual Belakang', 'HMB'),
(3, 'Amalatu', 'AML'),
(3, 'Inamosol', 'INM'),
(3, 'Elpaputih', 'EPH'),
(3, 'Huamual Utara', 'HMU'),
(3, 'Kepulauan Manipa', 'KMP'),
(3, 'Piru', 'PIR');

-- Kabupaten Seram Bagian Timur (regencies_id = 4)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(4, 'Bula', 'BLA'),
(4, 'Werinama', 'WRN'),
(4, 'Pulau Gorom', 'PGR'),
(4, 'Wakate', 'WKT'),
(4, 'Tutuk Tolu', 'TTL'),
(4, 'Siwalalat', 'SWL'),
(4, 'Kilmury', 'KLM'),
(4, 'Teluk Waru', 'TWR'),
(4, 'Gorom Timur', 'GTM'),
(4, 'Bula Barat', 'BLB'),
(4, 'Kota Bula', 'KBL'),
(4, 'Teluk Elpaputih', 'TEP'),
(4, 'Ujung Latu', 'UJL'),
(4, 'Banda', 'BDA'),
(4, 'Tehoru', 'THO');

-- Kabupaten Kepulauan Aru (regencies_id = 5)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(5, 'Dobo', 'DOB'),
(5, 'Aru Selatan', 'ARS'),
(5, 'Aru Tengah', 'ART'),
(5, 'Aru Utara', 'ARU'),
(5, 'Aru Tengah Timur', 'ATT'),
(5, 'Aru Tengah Selatan', 'ATS'),
(5, 'Aru Utara Timur Batuley', 'ATB'),
(5, 'Sir-Sir', 'SIR'),
(5, 'Pulau-Pulau Aru', 'PPA');

-- Kabupaten Maluku Barat Daya (regencies_id = 6)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(6, 'Tanimbar Selatan', 'TMS'),
(6, 'Tanimbar Utara', 'TMU'),
(6, 'Selaru', 'SLR'),
(6, 'Wer Tamrian', 'WTM'),
(6, 'Wer Maktian', 'WMK'),
(6, 'Nirunmas', 'NRM'),
(6, 'Molu Maru', 'MLM'),
(6, 'Wetar', 'WTR'),
(6, 'Wetar Barat', 'WTB'),
(6, 'Wetar Timur', 'WTT'),
(6, 'Wetar Utara', 'WTU'),
(6, 'Kisar Utara', 'KSU'),
(6, 'Pulau Leti', 'PLT'),
(6, 'Pulau Lakor', 'PLK'),
(6, 'Dawelor Dawera', 'DWD'),
(6, 'Pulau Moa', 'PMO'),
(6, 'Babar', 'BBR'),
(6, 'Babar Timur', 'BBT'),
(6, 'Dai', 'DAI'),
(6, 'Masela', 'MSL'),
(6, 'Batarkusu', 'BTK'),
(6, 'Romang', 'RMG'),
(6, 'Leti Moa Lakor', 'LML'),
(6, 'Mdona Hiera', 'MDH'),
(6, 'Tanimbar Utara', 'TNB');

-- Kabupaten Buru (regencies_id = 7)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(7, 'Namlea', 'NAM'),
(7, 'Air Buaya', 'ABY'),
(7, 'Waeapo', 'WAP'),
(7, 'Waplau', 'WPL'),
(7, 'Batabual', 'BTB'),
(7, 'Lolong Guba', 'LGB'),
(7, 'Teluk Kaiely', 'TKL'),
(7, 'Lilialy', 'LIL'),
(7, 'Waelata', 'WLT'),
(7, 'Fena Leisela', 'FLS'),
(7, 'Kepala Madan', 'KMD'),
(7, 'Waesama', 'WSM'),
(7, 'Dayamurni', 'DYM');

-- Kabupaten Buru Selatan (regencies_id = 8)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(8, 'Namrole', 'NMR'),
(8, 'Leksula', 'LKS'),
(8, 'Kepala Madan', 'KPM'),
(8, 'Waesama', 'WSA'),
(8, 'Ambalau', 'AMB'),
(8, 'Buru Selatan', 'BRS');

-- Kota Tual (regencies_id = 9)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(9, 'Dullah Utara', 'DLU'),
(9, 'Dullah Selatan', 'DLS'),
(9, 'Tayando', 'TYD'),
(9, 'Pulau Dullah Utara', 'PDU'),
(9, 'Kur Selatan', 'KRS');

-- Kabupaten Maluku Tenggara (regencies_id = 10)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(10, 'Kei Kecil', 'KKC'),
(10, 'Kei Besar', 'KBS'),
(10, 'Kei Kecil Timur', 'KKT'),
(10, 'Kei Besar Selatan', 'KBS'),
(10, 'Kei Kecil Barat', 'KKB'),
(10, 'Kei Besar Utara Timur', 'KBT'),
(10, 'Hoat Sorbay', 'HSB'),
(10, 'Manyeuw', 'MYW'),
(10, 'Kei Besar Utara', 'KBU'),
(10, 'Kei Kecil Timur Selatan', 'KTS'),
(10, 'Pulau Kur', 'PKR'),
(10, 'Kur Selatan', 'KUR');

-- Kabupaten Maluku Tenggara Barat (regencies_id = 11)
INSERT INTO districts (regencies_id, districts_name, districts_code) VALUES 
(11, 'Tanimbar Selatan', 'TNS'),
(11, 'Tanimbar Utara', 'TNU'),
(11, 'Selaru', 'SLU'),
(11, 'Wer Tamrian', 'WTN'),
(11, 'Wer Maktian', 'WMN'),
(11, 'Nirunmas', 'NRN'),
(11, 'Molu Maru', 'MLR'),
(11, 'Yaru', 'YRU'),
(11, 'Kormomolin', 'KML'),
(11, 'Fordata', 'FDT');

-- Insert Desa/Kelurahan untuk beberapa kecamatan utama
-- Kota Ambon - Kecamatan Sirimau
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(2, 'Mardika', 'MDK', 'kelurahan'),
(2, 'Batu Merah', 'BTM', 'kelurahan'),
(2, 'Benteng', 'BTG', 'kelurahan'),
(2, 'Wainitu', 'WNT', 'kelurahan'),
(2, 'Honipopu', 'HNP', 'kelurahan'),
(2, 'Rijali', 'RJL', 'kelurahan'),
(2, 'Karang Panjang', 'KPJ', 'kelurahan'),
(2, 'Ahusen', 'AHS', 'kelurahan'),
(2, 'Galala', 'GLL', 'kelurahan'),
(2, 'Soya', 'SOY', 'kelurahan');

-- Kota Ambon - Kecamatan Nusaniwe
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(1, 'Nusaniwe', 'NSW', 'kelurahan'),
(1, 'Waihaong', 'WHG', 'kelurahan'),
(1, 'Batu Gajah', 'BTG', 'kelurahan'),
(1, 'Kudamati', 'KDM', 'kelurahan'),
(1, 'Hatalae', 'HTL', 'kelurahan'),
(1, 'Waiheru', 'WHR', 'kelurahan'),
(1, 'Lateri', 'LTR', 'kelurahan'),
(1, 'Urimessing', 'URM', 'kelurahan'),
(1, 'Passo', 'PSO', 'kelurahan'),
(1, 'Hukurila', 'HKR', 'kelurahan');

-- Maluku Tengah - Kecamatan Salahutu
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(7, 'Tulehu', 'TLH', 'desa'),
(7, 'Liang', 'LNG', 'desa'),
(7, 'Waai', 'WAI', 'desa'),
(7, 'Tial', 'TIL', 'desa'),
(7, 'Rutong', 'RTG', 'desa'),
(7, 'Tengah-Tengah', 'TTG', 'desa'),
(7, 'Morella', 'MRL', 'desa'),
(7, 'Mamala', 'MML', 'desa'),
(7, 'Hitu', 'HTU', 'desa'),
(7, 'Hila', 'HLA', 'desa');

-- Maluku Tengah - Kecamatan Leihitu
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(8, 'Allang', 'ALG', 'desa'),
(8, 'Batu Merah', 'BTM', 'desa'),
(8, 'Hitu', 'HTU', 'desa'),
(8, 'Kaitetu', 'KTT', 'desa'),
(8, 'Mamala', 'MML', 'desa'),
(8, 'Seith', 'STH', 'desa'),
(8, 'Hitumessing', 'HTM', 'desa'),
(8, 'Hila', 'HLA', 'desa'),
(8, 'Kaitetu', 'KTT', 'desa'),
(8, 'Lilibooi', 'LLB', 'desa');

-- Seram Bagian Barat - Kecamatan Kairatu
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(21, 'Kairatu', 'KRT', 'desa'),
(21, 'Hatusua', 'HTS', 'desa'),
(21, 'Buria', 'BRI', 'desa'),
(21, 'Hatumeten', 'HTM', 'desa'),
(21, 'Lohia Sapalewa', 'LSP', 'desa'),
(21, 'Haturete', 'HTR', 'desa'),
(21, 'Rumahkay', 'RMK', 'desa'),
(21, 'Kamal', 'KML', 'desa'),
(21, 'Murnaten', 'MRN', 'desa'),
(21, 'Lohiatala', 'LHT', 'desa');

-- Buru - Kecamatan Namlea
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(47, 'Namlea', 'NML', 'kelurahan'),
(47, 'Wamlana', 'WML', 'desa'),
(47, 'Sawa', 'SWA', 'desa'),
(47, 'Waenetat', 'WNT', 'desa'),
(47, 'Wangongira', 'WGR', 'desa'),
(47, 'Waekasar', 'WKS', 'desa'),
(47, 'Jikumerasa', 'JKM', 'desa'),
(47, 'Waegeren', 'WGR', 'desa'),
(47, 'Waetawa', 'WTW', 'desa'),
(47, 'Kampung Baru', 'KBR', 'desa');

-- Kota Tual - Kecamatan Dullah Utara
INSERT INTO villages (districts_id, villages_name, villages_code, villages_type) VALUES 
(60, 'Langgur', 'LGR', 'kelurahan'),
(60, 'Ohoijang', 'OHJ', 'kelurahan'),
(60, 'Feer', 'FER', 'kelurahan'),
(60, 'Bombay', 'BMY', 'kelurahan'),
(60, 'Ruat', 'RUT', 'kelurahan'),
(60, 'Watdek', 'WTD', 'kelurahan'),
(60, 'Debut', 'DBT', 'kelurahan'),
(60, 'Dullah Laut', 'DLL', 'kelurahan'),
(60, 'Sathean', 'STH', 'kelurahan'),
(60, 'Revav', 'RVV', 'kelurahan');-- Update created_at for provinces
UPDATE provinces SET created_at = NOW();

-- Update created_at for regencies
UPDATE regencies SET created_at = NOW();

-- Update created_at for districts
UPDATE districts SET created_at = NOW();

-- Update created_at for villages
UPDATE villages SET created_at = NOW();
