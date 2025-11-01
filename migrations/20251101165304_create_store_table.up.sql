CREATE TABLE stores
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID         NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stores_user_id ON stores (user_id);
