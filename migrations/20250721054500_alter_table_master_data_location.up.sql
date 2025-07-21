ALTER TABLE provinces
ADD COLUMN latitude DECIMAL(15, 8) AFTER province_code,
ADD COLUMN longitude DECIMAL(15, 8) AFTER latitude;

-- Add columns to regencies table
ALTER TABLE regencies
ADD COLUMN latitude DECIMAL(15, 8) AFTER regency_type,
ADD COLUMN longitude DECIMAL(15, 8) AFTER latitude;

-- Add columns to districts table
ALTER TABLE districts
ADD COLUMN latitude DECIMAL(15, 8) AFTER district_code,
ADD COLUMN longitude DECIMAL(15, 8) AFTER latitude;

-- Add columns to villages table
ALTER TABLE villages
ADD COLUMN latitude DECIMAL(15, 8) AFTER village_type,
ADD COLUMN longitude DECIMAL(15, 8) AFTER latitude;