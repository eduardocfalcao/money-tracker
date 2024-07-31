CREATE TABLE IF NOT EXISTS raw_transactions_categories (
  raw_transaction_id integer,
  category_id integer,
  PRIMARY KEY(raw_transaction_id, category_id),
  CONSTRAINT fk_raw_transaction FOREIGN KEY(raw_transaction_id) REFERENCES users(id),
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES users(id)
);

