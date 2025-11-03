CREATE TABLE products_stores
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID         NOT NULL,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_products_stores_user_id ON products_stores (user_id);
