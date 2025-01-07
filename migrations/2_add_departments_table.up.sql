CREATE TABLE departments
(
    department_id UUID        NOT NULL PRIMARY KEY,
    name          VARCHAR(33) NOT NULL,
    manager_id    UUID        NOT NULL
);