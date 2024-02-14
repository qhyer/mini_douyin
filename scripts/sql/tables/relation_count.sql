CREATE TABLE `relation_count` (
                                  `user_id` BIGINT NOT NULL,
                                  `follow_count` BIGINT NOT NULL,
                                  `follower_count` BIGINT NOT NULL,
                                  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `deleted_at` DATETIME NULL DEFAULT NULL,
                                  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;