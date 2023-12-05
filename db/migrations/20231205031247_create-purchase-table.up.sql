CREATE TABLE purchase (
    id UUID PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    amount       DECIMAL(20,2) NOT NULL,
    transaction_date   DATE NOt NULL
);
