CREATE TABLE IF NOT EXISTS users (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  name varchar(200) NOT NULL,
  email varchar(100) NOT NULL UNIQUE,
  passwordHash varchar(256) NOT NULL,
  salt varchar(25) NOT NULL
);
