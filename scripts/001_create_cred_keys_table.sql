CREATE TABLE cred_keys (
    id INT AUTO_INCREMENT PRIMARY KEY,
    key_value VARCHAR(255) UNIQUE NOT NULL,
    key_type ENUM('activation', 'trial', 'partner') DEFAULT 'activation',
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
