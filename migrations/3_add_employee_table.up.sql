CREATE TABLE employees (
    identity_number VARCHAR(33) NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    name VARCHAR(33) NOT NULL,
    employee_image_uri TEXT,
    department_id UUID NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (department_id) REFERENCES departments(department_id)
);

BEGIN;

CREATE TYPE enum_gender AS ENUM (
	'male',
	'female'
);
ALTER TABLE employees ADD COLUMN gender enum_gender;

COMMIT;