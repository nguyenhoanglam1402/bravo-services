CREATE TABLE s_category_model(
  id uuid DEFAULT gen_random_uuid(),
  name VARCHAR(255),

  PRIMARY KEY(id),

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3)
);

CREATE TABLE s_lesson_model(
  id uuid DEFAULT gen_random_uuid(),
  name VARCHAR(255),
  author VARCHAR(255) NOT NULL,
  raw_data VARCHAR,
  comp_data VARCHAR,

  category_id uuid NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_lesson_category_model
  FOREIGN KEY (category_id) REFERENCES s_category_model(id) ON DELETE CASCADE
);