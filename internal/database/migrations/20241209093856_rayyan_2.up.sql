CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    total_amount DECIMAL(10,2),
    payment_method VARCHAR(50),
    customer_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);