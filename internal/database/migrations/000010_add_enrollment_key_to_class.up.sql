-- Add enrollment_key column to s_class_models
ALTER TABLE s_class_models
ADD COLUMN enrollment_key VARCHAR (255) NOT NULL;
