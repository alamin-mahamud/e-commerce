--PostgreSQL 9.6
--'\\' is a delimiter
CREATE OR REPLACE FUNCTION trigger_set_timestamp
()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW
();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users
(
  id serial PRIMARY KEY,
  email VARCHAR (355) UNIQUE NOT NULL,
  password VARCHAR (50) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp
();