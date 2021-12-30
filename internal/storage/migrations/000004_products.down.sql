ALTER TABLE IF EXISTS merchants DROP CONSTRAINT merchants_users_id_fkey;
ALTER TABLE IF EXISTS products DROP CONSTRAINT products_merchant_id_fkey;

DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS merchants;