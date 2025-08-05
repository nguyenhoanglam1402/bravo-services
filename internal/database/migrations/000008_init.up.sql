BEGIN;

ALTER TABLE s_lesson_model RENAME TO s_lesson_models;
ALTER TABLE s_lesson_models ALTER COLUMN updated_at DROP NOT NULL;
ALTER TABLE s_branch RENAME TO s_branches;
ALTER TABLE s_branches ALTER COLUMN updated_at DROP NOT NULL;
ALTER TABLE s_branches ADD COLUMN deleted_at TIMESTAMP NULL;
ALTER TABLE s_version RENAME TO s_versions;
ALTER TABLE s_versions ADD COLUMN updated_at TIMESTAMP NULL;
ALTER TABLE s_versions ADD COLUMN deleted_at TIMESTAMP NULL;
ALTER TABLE s_category_model RENAME TO s_category_models;

COMMIT;