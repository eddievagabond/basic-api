CREATE TABLE IF NOT EXISTS "orders" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (public.gen_random_uuid()),
  "user_id" uuid NOT NULL,
  "status"  text not null check (status in ('processing', 'shipped', 'cancelled', 'returned')),
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "order_items" (
  "order_id" uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (public.gen_random_uuid()),
  "product_id" uuid NOT NULL,
  "quantity" int DEFAULT 1
);


ALTER TABLE orders ADD CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE order_items ADD CONSTRAINT order_items_order_id_fk FOREIGN KEY (order_id) REFERENCES orders (id);
ALTER TABLE order_items ADD CONSTRAINT order_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id);

CREATE INDEX ON orders (user_id);
CREATE INDEX ON order_items (order_id);
CREATE INDEX ON order_items (product_id);
