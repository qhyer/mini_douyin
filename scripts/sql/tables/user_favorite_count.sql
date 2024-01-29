CREATE TABLE `user_favorite_count` (
    `user_id` BIGINT NOT NULL,
    `favorite_count` BIGINT NOT NULL,
    `favorited_count` BIGINT NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;