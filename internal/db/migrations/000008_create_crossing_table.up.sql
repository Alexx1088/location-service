CREATE TABLE crossings (
    id SERIAL PRIMARY KEY,
    crossroad_id INTEGER NOT NULL REFERENCES crossroads(id) ON DELETE CASCADE,
    event_time TIMESTAMP not null
    );