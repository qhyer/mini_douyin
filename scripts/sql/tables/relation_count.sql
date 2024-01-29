CREATE TABLE `relation_count` (
                                  `user_id` BIGINT NOT NULL,
                                  `follow_count` BIGINT NOT NULL,
                                  `follower_count` BIGINT NOT NULL,
                                  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;