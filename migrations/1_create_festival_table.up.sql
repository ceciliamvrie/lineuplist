CREATE TABLE festival(
  id uuid UNIQUE,
  name varchar NOT NULL UNIQUE,
  date varchar,
  location varchar
);
