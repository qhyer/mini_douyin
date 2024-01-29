CREATE TABLE `comment_count` (
                                 `video_id` BIGINT NOT NULL,
                                 `comment_count` BIGINT NOT NULL,
                                 PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;