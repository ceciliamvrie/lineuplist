ALTER TABLE artist
  ADD COLUMN img_src varchar,
  ADD COLUMN external_url varchar,
  ADD COLUMN followers integer,
  ADD COLUMN popularity integer;

CREATE TABLE genre(
  id uuid UNIQUE,
  name varchar NOT NULL UNIQUE
);

CREATE TABLE artist_genre(
  id uuid UNIQUE,
  artist_id uuid REFERENCES artist(id),
  genre_id uuid REFERENCES genre(id)
);

CREATE TABLE related_artist(
  id uuid UNIQUE,
  artist_id uuid REFERENCES artist(id),
  related_id uuid REFERENCES artist(id)
);

CREATE TABLE top_track(
  id uuid UNIQUE,
  artist_id uuid REFERENCES artist(id),
  name varchar NOT NULL,
  img_src varchar,
  external_url varchar
);

CREATE TABLE album(
  id uuid UNIQUE NOT NULL,
  artist_id uuid REFERENCES artist(id),
  name varchar,
  img_src varchar,
  externalURL varchar
);
