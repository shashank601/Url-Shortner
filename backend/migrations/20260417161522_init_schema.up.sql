CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    short_code VARCHAR(20) UNIQUE NOT NULL,
    original_url VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    UNIQUE (customer_id, original_url)
);

CREATE TABLE click_events (
    id SERIAL PRIMARY KEY,
    url_id INTEGER NOT NULL,
    ip_address INET,
    user_agent TEXT,
    referrer TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
);


-- right now we are not storing weekly stats, we will do it later. for now just do 3 db round trips
-- CREATE TABLE url_stats (
--     id SERIAL PRIMARY KEY,
--     url_id INTEGER NOT NULL,
--     week_start DATE NOT NULL,
--     total_clicks INTEGER DEFAULT 0,
--     unique_clicks INTEGER DEFAULT 0,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
-- );

