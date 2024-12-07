CREATE TABLE wallets (
    id UUID PRIMARY KEY NOT NULL,
    amount BIGINT NOT NULL,
    operation VARCHAR(255) NOT NULL
);
