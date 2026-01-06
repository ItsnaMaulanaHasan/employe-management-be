CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    hire_date DATE NOT NULL,
    termination_date DATE,
    salary INTEGER NOT NULL CHECK (salary > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_employees_last_name ON employees (last_name);