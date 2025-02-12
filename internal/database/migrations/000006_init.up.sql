BEGIN;
ALTER TABLE s_lesson_model DROP CONSTRAINT IF EXISTS fk_less_user_model;
ALTER TABLE s_lesson_model RENAME COLUMN author TO author_id;
ALTER TABLE s_lesson_model 
ALTER COLUMN author_id TYPE UUID USING author_id::uuid;
ALTER TABLE s_lesson_model 
ADD CONSTRAINT fk_less_user_model 
FOREIGN KEY (author_id) REFERENCES s_user_models(id);

COMMIT;