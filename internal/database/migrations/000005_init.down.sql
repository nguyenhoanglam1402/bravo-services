BEGIN;
  ALTER TABLE s_lesson_model DROP CONSTRAINT IF EXISTS fk_less_user_model;
  ALTER TABLE s_lesson_model ALTER COLUMN author TYPE
  VARCHAR
  (255);
  DROP TABLE IF EXISTS s_lesson_version;
COMMIT;