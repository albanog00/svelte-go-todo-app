-- | users |
CREATE TABLE `users` (
  `id` varchar(191) NOT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` longtext,
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_users_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- | tasks |
CREATE TABLE `tasks` (
  `id` varchar(191) NOT NULL,
  `description` longtext,
  `date` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_task` (`user_id`),
  CONSTRAINT `fk_users_task` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

