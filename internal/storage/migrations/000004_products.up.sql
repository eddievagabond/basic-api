CREATE TABLE IF NOT EXISTS  "merchants" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (public.gen_random_uuid()),
  "name" varchar,
  "created at" varchar,
  "admin_user_id" uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS  "products" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (public.gen_random_uuid()),
  "name" varchar,
  "merchant_id" uuid NOT NULL,
  "price" NUMERIC(10,2),
  "quantity" int DEFAULT 0,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE merchants ADD CONSTRAINT merchants_users_id_fkey FOREIGN KEY (admin_user_id) REFERENCES users (id);
ALTER TABLE products ADD CONSTRAINT products_merchant_id_fkey FOREIGN KEY (merchant_id) REFERENCES merchants (id);

CREATE INDEX ON merchants (admin_user_id);
CREATE INDEX ON products (merchant_id);