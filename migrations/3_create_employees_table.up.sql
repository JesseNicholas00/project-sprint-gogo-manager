CREATE TABLE employees
(
    identity_number VARCHAR(33) NOT NULL UNIQUE PRIMARY KEY,
    name VARCHAR(33) NOT NULL,
    employee_image_uri VARCHAR(255) NOT NULL,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')),
    department_id UUID REFERENCES departments(department_id) ON DELETE CASCADE,
    user_id TEXT REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_identity_number_hash ON employees USING hash (identity_number);