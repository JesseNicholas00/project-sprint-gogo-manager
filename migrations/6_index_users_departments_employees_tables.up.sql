-- users table
CREATE INDEX IF NOT EXISTS users_email_hash_idx ON users USING HASH (email);

-- departments table
CREATE INDEX IF NOT EXISTS departments_manager_id_hash_idx ON departments USING HASH (manager_id);
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX IF NOT EXISTS departments_name_trgm_idx ON departments USING GIN (name gin_trgm_ops);

-- employees table
CREATE INDEX IF NOT EXISTS employees_user_id_hash_idx ON employees USING HASH (user_id);
