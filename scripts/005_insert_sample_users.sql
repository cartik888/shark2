-- Insert sample users (passwords are hashed for 'password123')
INSERT INTO users (email, password, role, is_active, is_blocked, is_admin, created_at, updated_at) VALUES
('admin@shark.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 1, 0, 1, '2024-01-15 10:00:00.000', '2024-01-15 10:00:00.000'),
('john.doe@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 0, 0, '2024-02-01 14:30:00.000', '2024-02-01 14:30:00.000'),
('jane.smith@example.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 0, 0, '2024-02-15 09:15:00.000', '2024-02-15 09:15:00.000'),
('mike.wilson@company.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 1, 0, '2024-03-01 16:45:00.000', '2024-03-01 16:45:00.000'),
('sarah.johnson@startup.io', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 0, 0, '2024-03-10 11:20:00.000', '2024-03-10 11:20:00.000'),
('david.brown@tech.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 0, 0, 0, '2024-03-20 13:10:00.000', '2024-03-20 13:10:00.000'),
('lisa.garcia@design.co', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 0, 0, '2024-04-01 08:30:00.000', '2024-04-01 08:30:00.000'),
('robert.taylor@enterprise.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', 1, 0, 0, '2024-04-15 15:45:00.000', '2024-04-15 15:45:00.000');
