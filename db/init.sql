CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(225) NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    ttl_seconds INT,
    CONSTRAINT unique_short_code UNIQUE (short_code)
);


CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, 
    username VARCHAR(225) UNIQUE NOT NULL,
    password_hash text not null, 
    created_at TIMESTAMP DEFAULT NOW()
);