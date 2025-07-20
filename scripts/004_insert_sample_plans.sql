-- Insert sample subscription plans
INSERT INTO plans (name, description, price, currency, interval, features, is_active, created_at, updated_at) VALUES
('Basic', 'Perfect for individuals getting started', 9.99, 'USD', 'monthly', 'Up to 5 projects, 10GB storage, Email support, Basic analytics', 1, '2024-01-01 00:00:00.000', '2024-01-01 00:00:00.000'),
('Pro', 'Great for growing businesses', 29.99, 'USD', 'monthly', 'Up to 50 projects, 100GB storage, Priority support, Advanced analytics, Team collaboration, API access', 1, '2024-01-01 00:00:00.000', '2024-01-01 00:00:00.000'),
('Enterprise', 'For large organizations', 99.99, 'USD', 'monthly', 'Unlimited projects, 1TB storage, 24/7 phone support, Custom analytics, Advanced team management, API access, Custom integrations, SLA guarantee', 1, '2024-01-01 00:00:00.000', '2024-01-01 00:00:00.000');
