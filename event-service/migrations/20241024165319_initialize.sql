-- +goose Up
CREATE TABLE venues (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        description TEXT NOT NULL,
                        location VARCHAR(255) NOT NULL
);

CREATE TABLE performers (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL
);

CREATE TABLE events (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        description TEXT NOT NULL,
                        venue_id INTEGER REFERENCES venues(id) ON DELETE SET NULL,
                        performer_id INTEGER REFERENCES performers(id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS performers;
DROP TABLE IF EXISTS venues;