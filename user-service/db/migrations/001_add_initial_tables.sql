CREATE TABLE end_user (
  id serial primary key,
  email varchar not null,
  username varchar not null,
  pass_hash varchar not null,
  pass_salt varchar,
  full_name varchar,
  contact varchar,
  resume_link varchar,
  portfolio_link varchar,
  is_blocked boolean,
  blocked_till timestamp,
  block_history jsonb,
  is_deactivated boolean,
  preferences jsonb
);

CREATE TABLE cosmetic_item (
  id serial primary key,
  cosmetic_type varchar not null,
  description varchar,
  price bigint,
  avatar_image varchar,
  banner_image varchar,
  editor_color_type varchar,
  editor_colors jsonb
);

CREATE TABLE cosmetic_inventory_item (
  id serial primary key,
  end_user_id integer references end_user not null,
  cosmetic_item_id integer references cosmetic_item not null,
  currency_used bigint not null,
  is_equipped boolean default false,
  time_purchased timestamp
);

---- create above / drop below ----
DROP TABLE cosmetic_inventory_item;

DROP TABLE end_user;

DROP TABLE cosmetic_item;