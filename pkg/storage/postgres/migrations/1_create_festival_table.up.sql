CREATE TABLE festival(
  id UUID UNIQUE,
  name VARCHAR NOT NULL UNIQUE,
  date VARCHAR,
  location VARCHAR
);