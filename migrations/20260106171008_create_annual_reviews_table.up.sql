CREATE TABLE annual_reviews (
    id BIGSERIAL PRIMARY KEY,
    employee_id BIGINT NOT NULL,
    review_date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE annual_reviews
ADD CONSTRAINT fk_annual_reviews_employee_id FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE;

CREATE INDEX idx_annual_reviews_employee_id ON annual_reviews (employee_id);