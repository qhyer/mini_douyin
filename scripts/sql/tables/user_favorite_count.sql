CREATE TABLE `user_favorite_count` (
    `user_id` INT(11) NOT NULL,
    `favorite_count` INT(11) NOT NULL,
    `favorited_count` INT(11) NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;