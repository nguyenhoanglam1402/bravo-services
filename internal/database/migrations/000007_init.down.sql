BEGIN;

ALTER TABLE s_lesson_model DROP CONSTRAINT IF EXISTS fk_lesson_main_branch;
ALTER TABLE s_version DROP CONSTRAINT IF EXISTS s_version_branch_id_fkey;
ALTER TABLE s_branch DROP CONSTRAINT IF EXISTS  fk_branch_lesson;
ALTER TABLE s_branch DROP CONSTRAINT IF EXISTS  fk_branch_user;


ALTER TABLE s_lesson_model DROP COLUMN IF EXISTS main_branch_id;

ALTER TABLE s_lesson_model 
ADD COLUMN raw_data VARCHAR;

ALTER TABLE s_lesson_model 
ADD COLUMN comp_data VARCHAR;

DROP TABLE IF EXISTS s_version
CASCADE;
DROP TABLE IF EXISTS s_branch
CASCADE;

COMMIT;