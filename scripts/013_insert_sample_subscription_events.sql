-- Insert sample subscription events
INSERT INTO subscription_events (subscription_id, event_type, description, created_at) VALUES
(1, 'created', 'Subscription created for Basic plan', '2024-02-01 14:30:00.000'),
(1, 'payment_succeeded', 'Monthly payment of $9.99 processed successfully', '2024-02-01 14:30:00.000'),
(1, 'payment_succeeded', 'Monthly payment of $9.99 processed successfully', '2024-03-01 14:30:00.000'),
(2, 'created', 'Subscription created for Pro plan', '2024-02-15 09:15:00.000'),
(2, 'payment_succeeded', 'Monthly payment of $29.99 processed successfully', '2024-02-15 09:15:00.000'),
(2, 'payment_succeeded', 'Monthly payment of $29.99 processed successfully', '2024-03-15 09:15:00.000'),
(3, 'created', 'Subscription created for Pro plan', '2024-03-01 16:45:00.000'),
(3, 'payment_succeeded', 'Monthly payment of $29.99 processed successfully', '2024-03-01 16:45:00.000'),
(3, 'cancelled', 'Subscription cancelled by user', '2024-05-15 10:00:00.000'),
(4, 'created', 'Subscription created for Enterprise plan', '2024-03-10 11:20:00.000'),
(4, 'payment_succeeded', 'Monthly payment of $99.99 processed successfully', '2024-03-10 11:20:00.000'),
(4, 'payment_succeeded', 'Monthly payment of $99.99 processed successfully', '2024-04-10 11:20:00.000'),
(5, 'created', 'Subscription created for Basic plan', '2024-04-01 08:30:00.000'),
(5, 'payment_succeeded', 'Monthly payment of $9.99 processed successfully', '2024-04-01 08:30:00.000'),
(6, 'created', 'Subscription created for Enterprise plan', '2024-04-15 15:45:00.000'),
(6, 'payment_succeeded', 'Monthly payment of $99.99 processed successfully', '2024-04-15 15:45:00.000');
