CREATE TABLE subscription_keys (
    id INT AUTO_INCREMENT PRIMARY KEY,
    original_key_id INT NOT NULL,
    dummy_key VARCHAR(255) UNIQUE NOT NULL,
    assigned_to_user_id INT,
    is_used BOOLEAN DEFAULT FALSE,
    assigned_at DATETIME,
    FOREIGN KEY (original_key_id) REFERENCES cred_keys(id),
    FOREIGN KEY (assigned_to_user_id) REFERENCES users(id)
);
