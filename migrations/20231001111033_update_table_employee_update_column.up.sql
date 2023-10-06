ALTER TABLE `employee` CHANGE `bidang` `bidang_id` INT(11) NULL DEFAULT NULL;
ALTER TABLE `employee` ADD UNIQUE(`bidang_id`);

ALTER TABLE `employee` CHANGE `seksi` `seksi_id` INT(11) NULL DEFAULT NULL;
ALTER TABLE `employee` ADD UNIQUE(`seksi_id`);

ALTER TABLE `employee` CHANGE `unit_kerja` `uk_id` INT(11) NULL DEFAULT NULL;
ALTER TABLE `employee` ADD UNIQUE(`uk_id`);