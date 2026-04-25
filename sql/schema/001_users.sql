-- +goose Up 
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(254) UNIQUE NOT NULL ,
  password VARCHAR(32) NOT NULL 
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;
