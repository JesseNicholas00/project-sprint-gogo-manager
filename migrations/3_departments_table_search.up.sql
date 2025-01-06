CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE INDEX IF NOT EXISTS idx_departments_name_trgm ON departments USING gin (name gin_trgm_ops);