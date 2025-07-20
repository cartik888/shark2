-- Insert sample support messages
INSERT INTO support_messages (ticket_id, user_id, message, is_from_admin, created_at) VALUES
(1, 2, 'Hi, I am unable to access my dashboard after logging in. The page just shows a loading spinner and never loads.', 0, '2024-07-10 09:00:00.000'),
(1, 1, 'Hello John, thank you for contacting support. Can you please try clearing your browser cache and cookies, then attempt to log in again?', 1, '2024-07-10 10:30:00.000'),
(1, 2, 'I tried clearing cache but the issue persists. I am using Chrome browser on Windows 11.', 0, '2024-07-10 11:15:00.000'),
(2, 3, 'I upgraded to Pro plan but I cannot see the advanced analytics feature. Can you help me understand how to access it?', 0, '2024-07-08 14:30:00.000'),
(2, 1, 'Hi Jane, the advanced analytics feature is available in the Reports section of your dashboard. You should see a new "Advanced Analytics" tab there.', 1, '2024-07-09 10:15:00.000'),
(3, 4, 'I would like to cancel my subscription. Please process this request.', 0, '2024-05-15 16:45:00.000'),
(3, 1, 'Hi Mike, I have processed your cancellation request. Your subscription will remain active until the end of your current billing period (June 1st, 2024).', 1, '2024-05-16 09:30:00.000'),
(4, 5, 'I need documentation for API integration. Where can I find the API endpoints and authentication details?', 0, '2024-07-12 11:20:00.000'),
(5, 7, 'My payment failed yesterday and I cannot access my account features. Please help resolve this.', 0, '2024-07-11 08:30:00.000'),
(5, 1, 'Hi Lisa, I can see the payment failure in our system. I will send you a secure payment link to update your payment method.', 1, '2024-07-11 14:45:00.000'),
(6, 8, 'It would be great to have custom reporting features where we can create our own dashboard widgets.', 0, '2024-07-09 15:45:00.000');
