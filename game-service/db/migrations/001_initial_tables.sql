CREATE TABLE achievement (
  id serial primary key,
  title varchar not null,
  details varchar,
  badge_img varchar
);

CREATE TABLE user_achievement (
  id serial primary key,
  end_user_id integer not null,
  achievement_id integer references achievement not null
);

CREATE TABLE player_rank (
  id serial primary key,
  title varchar not null,
  description varchar,
  mmr_lowest decimal,
  mmr_highest decimal,
  rank_img varchar
);

CREATE TABLE user_mmr (
  id serial primary key,
  end_user_id integer not null,
  player_rank_id integer references player_rank,
  mmr decimal not null,
  deviation decimal not null,
  volatility decimal not null,
  num_games integer,
  game_history varchar[]
);

CREATE TABLE match (
  id serial primary key,
  game_state integer,
  game_type varchar,
  target_img varchar,
  match_time timestamp,
  num_players integer not null,
  total_time varchar,
  time_taken varchar
);

CREATE TABLE match_performance (
  id serial primary key,
  end_user_id integer not null,
  match_id integer references match not null,
  position integer,
  final_img varchar,
  acc_score decimal,
  precision_vals jsonb,
  score_changes decimal[],
  is_penalized boolean,
  penalty_duration varchar,
  penalty_history varchar[],
  mmr_from decimal,
  mmr_to decimal
);

CREATE TABLE suggestion (
  id serial primary key,
  match_performance_id integer references match_performance not null,
  suggestion_type varchar,
  reason varchar,
  advice varchar
);

---- create above / drop below ----
DROP TABLE suggestion;

DROP TABLE match_performance;

DROP TABLE match;

DROP TABLE user_mmr;

DROP TABLE player_rank;

DROP TABLE user_achievement;

DROP TABLE achievement;