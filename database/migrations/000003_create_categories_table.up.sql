CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY, 
  user_id integer NOT NULL,
  name varchar(200) NOT NULL,
  enabled boolean NOT NULL DEFAULT true
);

