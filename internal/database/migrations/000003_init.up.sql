CREATE TABLE s_role_models
(
  id uuid DEFAULT gen_random_uuid(),
  name VARCHAR(25),

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id)
);

ALTER TABLE s_user_models
ADD COLUMN role_id uuid NOT NULL;

ALTER TABLE s_user_models
ADD CONSTRAINT fk_user_model
FOREIGN KEY(role_id) REFERENCES s_role_models(id) ON DELETE CASCADE;