DROP TABLE  IF EXISTS product_stock;
CREATE TABLE IF NOT EXISTS product_stock(
    id serial PRIMARY KEY,
    product_id int NOT NULL,
    warehouse_id int NOT NULL,
    stock INTEGER NOT NULL,
    updated_at timestamptz NOT NULL DEFAULT NOW()
);

INSERT INTO product_stock(product_id, warehouse_id, stock)
VALUES (1, 1, 1);