CREATE TABLE s_organization_models
(
  id uuid DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  zip_code VARCHAR(100) NOT NULL,
  address VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  owner_id uuid NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),
  banned_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_organ_user
  FOREIGN KEY (owner_id) REFERENCES s_user_models(id) ON DELETE CASCADE
);

CREATE TABLE s_edu_group_models
(
  id uuid DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL,
  id_organization uuid NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),
  banned_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_edu_organ
  FOREIGN KEY(id_organization) REFERENCES s_organization_models(id) ON DELETE CASCADE
);

CREATE TABLE s_class_models
(
  id uuid DEFAULT gen_random_uuid(),
  class_name VARCHAR(50) NOT NULL,
  id_edu uuid NOT NULL,

  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  deleted_at TIMESTAMP(3),

  PRIMARY KEY(id),

  CONSTRAINT fk_class_edu
  FOREIGN KEY(id_edu) REFERENCES s_edu_group_models(id) ON DELETE CASCADE
);