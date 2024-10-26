-- +goose Up
-- +goose StatementBegin
CREATE TABLE seats (
   id BIGSERIAL PRIMARY KEY,
   seat_number VARCHAR(100) NOT NULL,
   venue_id INTEGER REFERENCES venues(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS seats
-- +goose StatementEnd
