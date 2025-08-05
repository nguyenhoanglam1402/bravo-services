BEGIN;

  DROP TABLE IF EXISTS s_lesson_version;

  CREATE TABLE s_branch
  (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lesson_id UUID NOT NULL,
    created_by UUID NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL,
    CONSTRAINT fk_branch_lesson
  FOREIGN KEY (lesson_id) REFERENCES s_lesson_model(id) ON DELETE CASCADE,
    CONSTRAINT fk_branch_user
  FOREIGN KEY (created_by) REFERENCES s_user_models(id) ON DELETE SET NULL
  );

  ALTER TABLE s_lesson_model 
ADD COLUMN main_branch_id UUID NULL;

ALTER TABLE s_lesson_model 
DROP COLUMN IF EXISTS raw_data;

ALTER TABLE s_lesson_model 
DROP COLUMN IF EXISTS comp_data;

ALTER TABLE s_lesson_model 
ADD CONSTRAINT fk_lesson_main_branch FOREIGN KEY (main_branch_id) REFERENCES s_branch(id) ON DELETE SET NULL;

CREATE TABLE s_version
(
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  branch_id UUID NOT NULL,
  author_id UUID NOT NULL,
  parent_version_id UUID NULL,
  raw_data TEXT NOT NULL,
  comp_data TEXT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT s_version_branch_id_fkey
  FOREIGN KEY (branch_id) REFERENCES s_branch(id) ON DELETE CASCADE,
  CONSTRAINT fk_version_auth_lesson
  FOREIGN KEY (author_id) REFERENCES s_user_models(id) ON DELETE CASCADE,
  CONSTRAINT fk_version_parent_version
  FOREIGN KEY (parent_version_id) REFERENCES s_version(id) ON DELETE SET NULL
);

COMMIT;