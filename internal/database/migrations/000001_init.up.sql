CREATE TABLE authentications
(
  id uuid DEFAULT gen_random_uuid(),
  username VARCHAR(50) NOT NULL,
  password VARCHAR NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE users
(
  id uuid DEFAULT gen_random_uuid(),
  fullname VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  job_title VARCHAR(50) NOT NULL,
  country VARCHAR NOT NULL,
  auth_id uuid NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,

  PRIMARY KEY(id),

  CONSTRAINT fk_user_auth
  FOREIGN KEY(auth_id) REFERENCES authentications(id) ON DELETE CASCADE
);