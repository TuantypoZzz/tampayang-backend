ALTER TABLE `employee` ADD `pangkat_id` INT(11) NULL DEFAULT NULL AFTER `bidang`;
ALTER TABLE `employee` ADD UNIQUE(`pangkat_id`);

ALTER TABLE `employee` ADD `golongan_id` INT(11) NULL DEFAULT NULL AFTER `pangkat_id`;
ALTER TABLE `employee` ADD UNIQUE(`golongan_id`);

ALTER TABLE `employee` ADD `tingkat` ENUM('E', 'F', '') DEFAULT '' AFTER `golongan_id`;

ALTER TABLE `employee` ADD `jabatan_id` INT(11) NULL DEFAULT NULL AFTER `seksi`;
ALTER TABLE `employee` ADD UNIQUE(`jabatan_id`);
