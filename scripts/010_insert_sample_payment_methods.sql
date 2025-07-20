-- Insert sample payment methods
INSERT INTO payment_methods (user_id, type, last4, brand, expiry_month, expiry_year, is_default, created_at, updated_at) VALUES
(2, 'card', '4242', 'Visa', 12, 2026, 1, '2024-02-01 14:30:00.000', '2024-02-01 14:30:00.000'),
(3, 'paypal', NULL, 'PayPal', NULL, NULL, 1, '2024-02-15 09:15:00.000', '2024-02-15 09:15:00.000'),
(4, 'card', '5555', 'Mastercard', 8, 2025, 1, '2024-03-01 16:45:00.000', '2024-03-01 16:45:00.000'),
(5, 'bank_transfer', NULL, 'Bank Transfer', NULL, NULL, 1, '2024-03-10 11:20:00.000', '2024-03-10 11:20:00.000'),
(7, 'card', '3782', 'American Express', 6, 2027, 1, '2024-04-01 08:30:00.000', '2024-04-01 08:30:00.000'),
(8, 'card', '6011', 'Discover', 10, 2025, 1, '2024-04-15 15:45:00.000', '2024-04-15 15:45:00.000');
