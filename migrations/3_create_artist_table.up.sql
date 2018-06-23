CREATE TABLE artist(
  id uuid UNIQUE,
  name varchar NOT NULL UNIQUE
);

CREATE TABLE festival_artist(
  id uuid UNIQUE,
  festival_id uuid REFERENCES festival(id),
  artist_id uuid REFERENCES artist(id)
);

