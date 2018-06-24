ALTER TABLE festival
  DROP COLUMN date RESTRICT,
  ADD COLUMN start_date date,
  ADD COLUMN end_date date,
  DROP COLUMN location RESTRICT,
  ADD COLUMN country varchar,
  ADD COLUMN state varchar,
  ADD COLUMN city varchar;
