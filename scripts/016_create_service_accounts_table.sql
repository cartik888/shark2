CREATE TABLE service_accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(32) NOT NULL,
    project_id VARCHAR(128) NOT NULL,
    private_key_id VARCHAR(128) NOT NULL,
    private_key TEXT NOT NULL,
    client_email VARCHAR(128) NOT NULL,
    client_id VARCHAR(64) NOT NULL,
    auth_uri VARCHAR(256) NOT NULL,
    token_uri VARCHAR(256) NOT NULL,
    auth_provider_x509_cert_url VARCHAR(256) NOT NULL,
    client_x509_cert_url VARCHAR(256) NOT NULL,
    universe_domain VARCHAR(64) NOT NULL,
    encrypted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
