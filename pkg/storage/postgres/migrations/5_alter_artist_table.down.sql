ALTER TABLE artist
  DROP COLUMN img_src RESTRICT,
  DROP COLUMN external_url RESTRICT,
  DROP COLUMN followers RESTRICT,
  DROP COLUMN popularity RESTRICT;

DROP TABLE IF EXISTS genre CASCADE;
DROP TABLE IF EXISTS artist_genre CASCADE;
DROP TABLE IF EXISTS related_artist CASCADE;
DROP TABLE IF EXISTS top_track CASCADE;
DROP TABLE IF EXISTS album CASCADE;