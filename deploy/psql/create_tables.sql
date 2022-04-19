CREATE TABLE users(
  id SERIAL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  pwd_hash TEXT NOT NULL
);

CREATE TABLE user_sessions(
  user_id SERIAL UNIQUE REFERENCES users(id),
  token TEXT NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  UNIQUE (user_id, token)
);