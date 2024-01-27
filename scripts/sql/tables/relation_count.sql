CREATE TABLE `relation_count` (
                                  `user_id` INT(11) NOT NULL,
                                  `follow_count` INT(11) NOT NULL,
                                  `follower_count` INT(11) NOT NULL,
                                  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;