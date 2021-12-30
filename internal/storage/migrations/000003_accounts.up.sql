CREATE TABLE IF NOT EXISTS accounts (
   id uuid DEFAULT public.gen_random_uuid() NOT NULL unique,
  user_id uuid NOT NULL,
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

ALTER TABLE accounts ADD CONSTRAINT accounts_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE entries ADD CONSTRAINT entries_accounts_id_fkey FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD CONSTRAINT transfers_from_account_id_fkey FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD CONSTRAINT transfers_to_account_id_fkey FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON accounts (user_id);
CREATE INDEX ON entries (account_id);
CREATE INDEX ON transfers (from_account_id);
CREATE INDEX ON transfers (to_account_id);
