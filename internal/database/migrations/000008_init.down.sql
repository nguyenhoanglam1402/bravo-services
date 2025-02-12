BEGIN;

ALTER TABLE s_lesson_models RENAME TO s_lesson_model;
ALTER TABLE s_lesson_model ALTER COLUMN updated_at SET NOT NULL;
ALTER TABLE s_branches RENAME TO s_branch;
ALTER TABLE s_branch DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE s_branch ALTER COLUMN updated_at SET NOT NULL;
ALTER TABLE s_versions RENAME TO s_version;
ALTER TABLE s_version DROP COLUMN IF EXISTS updated_at;
ALTER TABLE s_version DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE s_category_models RENAME TO s_category_model;
ALTER TABLE s_category_model DROP COLUMN IF EXISTS updated_at;
ALTER TABLE s_category_model DROP COLUMN IF EXISTS deleted_at;

COMMIT;