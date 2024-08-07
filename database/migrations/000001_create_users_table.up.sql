CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY, 
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  name varchar(200) NOT NULL,
  email varchar(100) NOT NULL UNIQUE,
  password_hash varchar(256) NOT NULL,
  salt varchar(32) NOT NULL
);
