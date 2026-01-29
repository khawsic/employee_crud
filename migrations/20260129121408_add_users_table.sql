CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    role TEXT ,
    created_at TIMESTAMP DEFAULT now()
    updated_at TIMESTAMP DEFAULT now()
);
