CREATE TABLE accounts (
   id uuid DEFAULT public.gen_random_uuid() NOT NULL unique,
  owner varchar NOT NULL,
  balance  NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  currency varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
  id uuid DEFAULT public.gen_random_uuid() NOT NULL unique,
  account_id uuid NOT NULL unique,
  amount NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
  id uuid DEFAULT public.gen_random_uuid() NOT NULL unique,
  from_account_id uuid NOT NULL unique,
  to_account_id uuid NOT NULL unique,
  amount NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts (id);

ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);

ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON entries (account_id);

CREATE INDEX ON transfers (from_account_id);

CREATE INDEX ON transfers (to_account_id);
