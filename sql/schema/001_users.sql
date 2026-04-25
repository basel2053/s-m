-- +goose Up 
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(254) UNIQUE NOT NULL,
  password VARCHAR(32) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users;
