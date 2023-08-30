CREATE TABLE IF NOT EXISTS
    urls (
        id SERIAL PRIMARY KEY,
        url TEXT NOT NULL,
        code VARCHAR(6) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );