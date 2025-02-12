BEGIN;
  ALTER TABLE s_lesson_model DROP CONSTRAINT IF EXISTS fk_less_user_model;
  ALTER TABLE s_lesson_model 

ALTER COLUMN author_id TYPE
  VARCHAR
  (255);
  ALTER TABLE s_lesson_model RENAME COLUMN author_id TO author;
  
COMMIT;