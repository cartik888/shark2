-- Insert sample subscription keys (dummy keys assigned to users)
INSERT INTO subscription_keys (original_key_id, dummy_key, assigned_to_user_id, is_used, assigned_at) VALUES
(1, 'USER-a1b2c3d4e5f6', 2, 1, '2024-02-01 14:30:00'),
(2, 'USER-b2c3d4e5f6g7', 3, 1, '2024-02-15 09:15:00'),
(3, 'USER-c3d4e5f6g7h8', 4, 1, '2024-03-01 16:45:00'),
(4, 'USER-d4e5f6g7h8i9', 5, 1, '2024-03-10 11:20:00'),
(5, 'USER-e5f6g7h8i9j0', 7, 1, '2024-04-01 08:30:00'),
(6, 'USER-f6g7h8i9j0k1', 8, 1, '2024-04-15 15:45:00');
