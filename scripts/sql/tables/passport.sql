CREATE TABLE `passport` (
                            `id` BIGINT NOT NULL,
                            `name` VARCHAR(255) NOT NULL,
                            `password` VARCHAR(255) NOT NULL,
                            `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` DATETIME NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE INDEX `name_idx` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;