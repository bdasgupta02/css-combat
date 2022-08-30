ALTER TABLE
  end_user
ADD
  CONSTRAINT username_unique UNIQUE (username);

ALTER TABLE
  end_user
ADD
  CONSTRAINT email_unique UNIQUE (email);

---- create above / drop below ----
ALTER TABLE
  end_user DROP CONSTRAINT username_unique;

ALTER TABLE
  end_user DROP CONSTRAINT email_unique;