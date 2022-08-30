ALTER TABLE end_user ADD currency BIGINT;
---- create above / drop below ----
ALTER TABLE end_user DROP COLUMN currency;
