CREATE TABLE  IF NOT EXISTS "users" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL DEFAULT (public.gen_random_uuid()),
  "email" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);