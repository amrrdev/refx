CREATE TABLE urls (
    short_code VARCHAR(10) PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    expires_at TIMESTAMP NULL
);

CREATE INDEX idx_redirect ON urls (short_code) INCLUDE (long_url);
