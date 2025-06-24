CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    status TEXT,
    note TEXT,
    created_time TIMESTAMP DEFAULT now(),
    updated_time TIMESTAMP DEFAULT now(),
    priority TEXT,
    author_id INT
);
