CREATE TABLE IF NOT EXISTS raw_transactions (
  id SERIAL PRIMARY KEY, 
  user_id integer NOT NULL,
  account_id varchar(20),
  date_posted timestamp NOT NULL, 
  transaction_amount money NOT NULL, 
  fit_id integer NULL, 
  checknum varchar(11) NULL,
  memo varchar(50) NULL,
  description varchar(300) NULL,
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);


