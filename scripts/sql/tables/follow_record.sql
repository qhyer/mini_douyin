CREATE TABLE `follow_record` (
                                 `id` BIGINT NOT NULL,
                                 `from_user_id` BIGINT NOT NULL,
                                 `to_user_id` BIGINT NOT NULL,
                                 `attr` INT(10) UNSIGNED NOT NULL,
                                 `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `deleted_at` DATETIME NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 INDEX `user_id_idx` (`from_user_id`,`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;