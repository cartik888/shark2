-- Insert sample invoices
INSERT INTO invoices (user_id, subscription_id, invoice_number, amount, currency, status, due_date, paid_at, created_at, updated_at) VALUES
(2, 1, 'INV-2024-001', 9.99, 'USD', 'paid', '2024-02-15 14:30:00.000', '2024-02-01 14:30:00.000', '2024-02-01 14:30:00.000', '2024-02-01 14:30:00.000'),
(3, 2, 'INV-2024-002', 29.99, 'USD', 'paid', '2024-03-01 09:15:00.000', '2024-02-15 09:15:00.000', '2024-02-15 09:15:00.000', '2024-02-15 09:15:00.000'),
(4, 3, 'INV-2024-003', 29.99, 'USD', 'paid', '2024-03-15 16:45:00.000', '2024-03-01 16:45:00.000', '2024-03-01 16:45:00.000', '2024-03-01 16:45:00.000'),
(5, 4, 'INV-2024-004', 99.99, 'USD', 'paid', '2024-03-24 11:20:00.000', '2024-03-10 11:20:00.000', '2024-03-10 11:20:00.000', '2024-03-10 11:20:00.000'),
(7, 5, 'INV-2024-005', 9.99, 'USD', 'paid', '2024-04-15 08:30:00.000', '2024-04-01 08:30:00.000', '2024-04-01 08:30:00.000', '2024-04-01 08:30:00.000'),
(8, 6, 'INV-2024-006', 99.99, 'USD', 'paid', '2024-04-29 15:45:00.000', '2024-04-15 15:45:00.000', '2024-04-15 15:45:00.000', '2024-04-15 15:45:00.000'),
(2, 1, 'INV-2024-007', 9.99, 'USD', 'pending', '2024-07-15 14:30:00.000', NULL, '2024-07-01 14:30:00.000', '2024-07-01 14:30:00.000'),
(3, 2, 'INV-2024-008', 29.99, 'USD', 'overdue', '2024-07-01 09:15:00.000', NULL, '2024-06-15 09:15:00.000', '2024-06-15 09:15:00.000');
