ALTER TABLE s_lesson_model ALTER COLUMN author TYPE
uuid USING author::uuid;

ALTER TABLE s_lesson_model
ADD CONSTRAINT fk_less_user_model
FOREIGN KEY (author) REFERENCES s_user_models(id);

CREATE TABLE s_lesson_version
(
  id uuid DEFAULT gen_random_uuid(),
  lession_id uuid NOT NULL,

  author VARCHAR(255) NOT NULL,
  raw_data VARCHAR,
  comp_data VARCHAR,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_version_lesson_model
  FOREIGN KEY (lession_id) REFERENCES s_lesson_model(id) ON DELETE CASCADE
)