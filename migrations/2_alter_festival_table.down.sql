ALTER TABLE festival
  ADD COLUMN date varchar,
  DROP COLUMN start_date RESTRICT,
  DROP COLUMN end_date RESTRICT,
  ADD COLUMN location varchar,
  DROP COLUMN country RESTRICT,
  DROP COLUMN state RESTRICT,
  DROP COLUMN city RESTRICT;
