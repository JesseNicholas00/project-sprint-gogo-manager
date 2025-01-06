CREATE TABLE departments
(
    id         UUID        NOT NULL PRIMARY KEY,
    name       VARCHAR(33) NOT NULL,
    manager_id UUID        NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_departments_id_hash ON departments USING hash (id);
CREATE INDEX IF NOT EXISTS idx_departments_manager_id_hash ON departments USING hash (manager_id);