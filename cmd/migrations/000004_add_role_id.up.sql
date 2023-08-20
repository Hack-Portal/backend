ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_account_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_room_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_role_fkey";
DROP TABLE IF EXISTS "rooms_accounts";
CREATE TABLE "rooms_accounts" (
  "rooms_account_id" varchar PRIMARY KEY,
  "account_id" varchar NOT NULL,
  "room_id" varchar NOT NULL,
  "is_owner" boolean NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "rooms_accounts_roles" (
  "rooms_account_id" varchar NOT NULL,
  "role_id" int NOT NULL
);
ALTER TABLE "rooms_accounts_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");
ALTER TABLE "rooms_accounts_roles" ADD FOREIGN KEY ("rooms_account_id") REFERENCES "rooms_accounts" ("rooms_account_id");
ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");
ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");