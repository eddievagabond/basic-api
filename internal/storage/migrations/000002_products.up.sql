
CREATE TABLE IF NOT EXISTS products
(
   id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);