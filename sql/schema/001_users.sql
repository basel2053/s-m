-- +goose Up 
CREATE TABLE users (
  id INT NOT NULL,
  username UNIQUE NOT NULL VARCHAR(40),
  password UNIQUE NOT NULL VARCHAR(40)
)

-- +goose Down
DROP TABLE users;
