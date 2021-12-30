ALTER TABLE IF EXISTS orders DROP CONSTRAINT orders_user_id_fkey;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT order_items_order_id_fk;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT order_items_product_id_fk;

DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_items;