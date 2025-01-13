-- Rollback for users table
DROP INDEX IF EXISTS users_email_hash_idx;

-- Rollback for departments table
DROP INDEX IF EXISTS departments_manager_id_hash_idx;
DROP INDEX IF EXISTS departments_name_trgm_idx;
DROP EXTENSION IF EXISTS "pg_trgm";

-- Rollback for employees table
DROP INDEX IF EXISTS employees_user_id_hash_idx;
