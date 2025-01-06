CREATE TABLE departments
(
    id   UUID        NOT NULL PRIMARY KEY,
    name VARCHAR(33) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_department_id_hash ON departments USING hash (id);