CREATE TABLE s_authent_models
(
  id uuid DEFAULT gen_random_uuid(),
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id)
);

CREATE TABLE s_user_models
(
  id uuid DEFAULT gen_random_uuid(),
  fullname VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  job_title VARCHAR(50) NOT NULL,
  country VARCHAR NOT NULL,
  auth_id uuid NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_user_auth
  FOREIGN KEY(auth_id) REFERENCES s_authent_models(id) ON DELETE CASCADE
);