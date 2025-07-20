-- Create users table (if not exists)
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `role` varchar(191) DEFAULT 'user',
  `is_active` tinyint(1) DEFAULT '1',
  `is_blocked` tinyint(1) DEFAULT '0',
  `is_admin` tinyint(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create user_profiles table
CREATE TABLE IF NOT EXISTS `user_profiles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `phone` longtext,
  `company` longtext,
  `avatar` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_profile` (`user_id`),
  CONSTRAINT `fk_users_profile` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create plans table
CREATE TABLE IF NOT EXISTS `plans` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `description` longtext,
  `price` double NOT NULL,
  `currency` varchar(191) DEFAULT 'USD',
  `interval` varchar(191) DEFAULT 'monthly',
  `features` longtext,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create subscriptions table (if not exists)
CREATE TABLE IF NOT EXISTS `subscriptions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `plan_id` bigint unsigned NOT NULL,
  `status` varchar(191) DEFAULT 'active',
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  `next_billing_date` datetime(3) DEFAULT NULL,
  `cancelled_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_subscriptions_user` (`user_id`),
  KEY `fk_subscriptions_plan` (`plan_id`),
  CONSTRAINT `fk_subscriptions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_subscriptions_plan` FOREIGN KEY (`plan_id`) REFERENCES `plans` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create subscription_events table
CREATE TABLE IF NOT EXISTS `subscription_events` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` bigint unsigned NOT NULL,
  `event_type` longtext NOT NULL,
  `description` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_subscriptions_events` (`subscription_id`),
  CONSTRAINT `fk_subscriptions_events` FOREIGN KEY (`subscription_id`) REFERENCES `subscriptions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create payments table (if not exists)
CREATE TABLE IF NOT EXISTS `payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `subscription_id` bigint unsigned DEFAULT NULL,
  `amount` double NOT NULL,
  `currency` varchar(191) DEFAULT 'USD',
  `status` varchar(191) DEFAULT 'pending',
  `payment_method` longtext,
  `transaction_id` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_payments_user` (`user_id`),
  KEY `fk_payments_subscription` (`subscription_id`),
  CONSTRAINT `fk_payments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_payments_subscription` FOREIGN KEY (`subscription_id`) REFERENCES `subscriptions` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create payment_methods table (if not exists)
CREATE TABLE IF NOT EXISTS `payment_methods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `type` longtext NOT NULL,
  `last4` longtext,
  `brand` longtext,
  `expiry_month` bigint DEFAULT NULL,
  `expiry_year` bigint DEFAULT NULL,
  `is_default` tinyint(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_payment_methods` (`user_id`),
  CONSTRAINT `fk_users_payment_methods` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create invoices table (if not exists)
CREATE TABLE IF NOT EXISTS `invoices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `subscription_id` bigint unsigned NOT NULL,
  `invoice_number` varchar(191) NOT NULL,
  `amount` double NOT NULL,
  `currency` varchar(191) DEFAULT 'USD',
  `status` varchar(191) DEFAULT 'pending',
  `due_date` datetime(3) DEFAULT NULL,
  `paid_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `invoice_number` (`invoice_number`),
  KEY `fk_invoices_user` (`user_id`),
  KEY `fk_invoices_subscription` (`subscription_id`),
  CONSTRAINT `fk_invoices_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_invoices_subscription` FOREIGN KEY (`subscription_id`) REFERENCES `subscriptions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create password_resets table (if not exists)
CREATE TABLE IF NOT EXISTS `password_resets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `token` varchar(191) NOT NULL,
  `expires_at` datetime(3) DEFAULT NULL,
  `used_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  KEY `fk_password_resets_user` (`user_id`),
  CONSTRAINT `fk_password_resets_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create cred_keys table (if not exists)
CREATE TABLE IF NOT EXISTS `cred_keys` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key_value` varchar(255) NOT NULL,
  `key_type` enum('activation','trial','partner') DEFAULT 'activation',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key_value` (`key_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create subscription_keys table (if not exists)
CREATE TABLE IF NOT EXISTS `subscription_keys` (
  `id` int NOT NULL AUTO_INCREMENT,
  `original_key_id` int NOT NULL,
  `dummy_key` varchar(255) NOT NULL,
  `assigned_to_user_id` bigint unsigned DEFAULT NULL,
  `is_used` tinyint(1) DEFAULT '0',
  `assigned_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `dummy_key` (`dummy_key`),
  KEY `original_key_id` (`original_key_id`),
  KEY `assigned_to_user_id` (`assigned_to_user_id`),
  CONSTRAINT `subscription_keys_ibfk_1` FOREIGN KEY (`original_key_id`) REFERENCES `cred_keys` (`id`) ON DELETE CASCADE,
  CONSTRAINT `subscription_keys_ibfk_2` FOREIGN KEY (`assigned_to_user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create support_tickets table
CREATE TABLE IF NOT EXISTS `support_tickets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `subject` varchar(191) NOT NULL,
  `status` varchar(191) DEFAULT 'open',
  `priority` varchar(191) DEFAULT 'medium',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_support_tickets_user` (`user_id`),
  CONSTRAINT `fk_support_tickets_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Create support_messages table
CREATE TABLE IF NOT EXISTS `support_messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ticket_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `message` longtext NOT NULL,
  `is_from_admin` tinyint(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_support_messages_ticket` (`ticket_id`),
  KEY `fk_support_messages_user` (`user_id`),
  CONSTRAINT `fk_support_messages_ticket` FOREIGN KEY (`ticket_id`) REFERENCES `support_tickets` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_support_messages_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
