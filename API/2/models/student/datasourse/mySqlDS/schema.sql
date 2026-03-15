CREATE TABLE IF NOT EXISTS `studentss` (
     `id` BIGINT NOT NULL AUTO_INCREMENT,
    `firstname` VARCHAR(128) NOT NULL,
    `lastname` VARCHAR(512) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL ,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_created_at` (`created_at`)
    );