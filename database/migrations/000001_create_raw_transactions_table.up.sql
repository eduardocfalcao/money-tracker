CREATE TABLE IF NOT EXISTS raw_transactions (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
  account_id varchar(20),
  data_posted timestamp NOT NULL, 
  transaction_amount money NOT NULL, 
  fit_id integer NOT NULL, 
  CHECKNUM varchar(11) NOT NULL,
  MEMO varchar(50) NOT NULL
);
