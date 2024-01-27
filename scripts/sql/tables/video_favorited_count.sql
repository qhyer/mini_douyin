CREATE TABLE `video_favorited_count` (
                                         `video_id` INT(11) NOT NULL,
                                         `favorited_count` INT(11) NOT NULL,
                                         PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;