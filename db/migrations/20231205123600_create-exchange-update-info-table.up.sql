CREATE TABLE exchange_rate_update_info (
    time TIMESTAMP PRIMARY KEY,
    retry_count INTEGER,
    retry_time TIMESTAMP,
    success BOOLEAN
);
