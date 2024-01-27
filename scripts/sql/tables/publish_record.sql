CREATE TABLE `publish_record` (
                                  `id` INT(11) NOT NULL,
                                  `author_id` INT(11) NOT NULL,
                                  `title` VARCHAR(255) NOT NULL,
                                  `play_url` VARCHAR(255) NOT NULL,
                                  `cover_url` VARCHAR(255) NOT NULL,
                                  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `deleted_at` DATETIME NULL DEFAULT NULL,
                                  PRIMARY KEY (`id`),
                                  INDEX `author_id_idx` (`author_id`),
                                  INDEX `created_at_idx` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;