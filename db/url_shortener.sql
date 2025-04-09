CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_code VARCHAR(225) NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    ttl_second INT,

    CONSTRAINT unique_short_code UNIQUE (short_code)


);