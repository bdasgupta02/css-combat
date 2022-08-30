ALTER TABLE
  end_user ALTER currency
SET
  DEFAULT 0;

---- create above / drop below ----
ALTER TABLE
  end_user ALTER currency DROP DEFAULT;