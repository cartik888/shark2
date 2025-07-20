-- MySQL Database Setup Script for Shark SaaS Backend
-- Run this script as MySQL root user

-- Create database
CREATE DATABASE IF NOT EXISTS shark_saas CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create a dedicated user for the application
CREATE USER IF NOT EXISTS 'shark_user'@'localhost' IDENTIFIED BY 'shark_password_2024';

-- Grant all privileges on the shark_saas database to the user
GRANT ALL PRIVILEGES ON shark_saas.* TO 'shark_user'@'localhost';

-- Flush privileges to ensure they take effect
FLUSH PRIVILEGES;

-- Use the database
USE shark_saas;

-- Show confirmation
SELECT 'Database shark_saas created successfully!' as message;
SELECT 'User shark_user created with full access!' as message;
