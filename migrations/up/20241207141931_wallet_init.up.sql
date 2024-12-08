CREATE TABLE wallets (
    id UUID PRIMARY KEY NOT NULL,
    amount BIGINT NOT NULL
);

INSERT INTO wallets (id, amount) VALUES
('103960a0-9a79-43ed-bff0-052a19eaa98e', 5000),
('f2bfc720-c67a-4f79-9bcb-47c146deb8e3', 10000),
('c6a51a4e-7c5e-4354-8576-d22c7649ad65', 15000),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 20000),
('b81d4fae-7dec-4c8e-8e0e-34d5dd4d88b0', 0),
('5a82d8af-1b3c-4f07-80f4-1c9be144c87e', 7500),
('29f360c5-8466-4dfb-8f94-ec4e004d632f', 500),
('8e4c1a17-4dbd-4910-a0d1-7e6eced20a5e', 100000),
('cad2e9e6-56cc-4aa9-b5c2-8722e0a1c7d7', 700),
('f4a6c8b1-34d3-4dc6-87b2-9b8a439dcc59', 1),
('77d7d5b5-54b7-41c3-9887-9a50b80cf6df', 92254775807);
