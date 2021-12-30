ALTER TABLE  IF EXISTS accounts DROP CONSTRAINT accounts_users_id_fkey;
ALTER TABLE  IF EXISTS entries DROP CONSTRAINT entries_accounts_id_fkey;
ALTER TABLE  IF EXISTS  transfers DROP CONSTRAINT transfers_from_account_id_fkey;
ALTER TABLE  IF EXISTS  transfers DROP CONSTRAINT transfers_to_account_id_fkey;

DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;