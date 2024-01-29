CREATE TABLE `video_favorited_count` (
                                         `video_id` BIGINT NOT NULL,
                                         `favorited_count` BIGINT NOT NULL,
                                         PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;