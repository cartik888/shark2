-- Insert sample plans
INSERT IGNORE INTO `plans` (`id`, `name`, `description`, `price`, `currency`, `interval`, `features`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 'Basic', 'Perfect for individuals getting started', 9.99, 'USD', 'monthly', '5 Projects, 10GB Storage, Email Support', 1, NOW(), NOW()),
(2, 'Pro', 'Great for growing businesses', 29.99, 'USD', 'monthly', 'Unlimited Projects, 100GB Storage, Priority Support, Advanced Analytics', 1, NOW(), NOW()),
(3, 'Enterprise', 'For large organizations', 99.99, 'USD', 'monthly', 'Everything in Pro, Custom Integrations, Dedicated Support, SLA', 1, NOW(), NOW());

-- Insert sample credential keys
INSERT IGNORE INTO `cred_keys` (`key_value`, `key_type`, `is_active`, `created_at`) VALUES
('ACTIVATION-KEY-001-ABCD1234', 'activation', 1, NOW()),
('ACTIVATION-KEY-002-EFGH5678', 'activation', 1, NOW()),
('ACTIVATION-KEY-003-IJKL9012', 'activation', 1, NOW()),
('ACTIVATION-KEY-004-MNOP3456', 'activation', 1, NOW()),
('ACTIVATION-KEY-005-QRST7890', 'activation', 1, NOW()),
('TRIAL-KEY-001-UVWX1234', 'trial', 1, NOW()),
('TRIAL-KEY-002-YZAB5678', 'trial', 1, NOW()),
('PARTNER-KEY-001-CDEF9012', 'partner', 1, NOW());

-- Insert a sample admin user (password: admin123)
INSERT IGNORE INTO `users` (`email`, `password`, `role`, `is_active`, `is_blocked`, `is_admin`, `created_at`, `updated_at`) VALUES
('admin@shark.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 1, 0, 1, NOW(), NOW());

-- Get the admin user ID and insert profile
SET @admin_user_id = (SELECT id FROM users WHERE email = 'admin@shark.com');
INSERT IGNORE INTO `user_profiles` (`user_id`, `first_name`, `last_name`, `company`, `created_at`, `updated_at`) VALUES
(@admin_user_id, 'Admin', 'User', 'Shark Company', NOW(), NOW());
