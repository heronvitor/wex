CREATE TABLE exchange_rate (
    record_date DATE NOt NULL,
    country VARCHAR(255),
    currency VARCHAR(50),
    exchange_rate DECIMAL(10,2),
    effective_date DATE NOt NULL,
    PRIMARY KEY(currency, record_date) 
);
