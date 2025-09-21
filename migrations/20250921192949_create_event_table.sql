-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  location_id INTEGER NOT NULL,
  type VARCHAR(50) NOT NULL,
  date TEXT NOT NULL,
  total_drivers INTEGER NOT NULL,
  car_type VARCHAR(50),
  FOREIGN KEY (location_id) REFERENCES location (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
-- +goose StatementEnd
