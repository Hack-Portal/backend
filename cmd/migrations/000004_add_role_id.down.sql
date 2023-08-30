ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_account_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_room_id_fkey";
ALTER TABLE "rooms_accounts_roles" DROP CONSTRAINT "rooms_accounts_roles_role_id_fkey";
ALTER TABLE "rooms_accounts_roles" DROP CONSTRAINT "rooms_accounts_roles_rooms_account_id_fkey";
DROP TABLE IF EXISTS "rooms_accounts";
DROP TABLE IF EXISTS "rooms_accounts_roles";
CREATE TABLE "rooms_accounts" (
  "account_id" varchar NOT NULL,
  "room_id" varchar NOT NULL,
  "role" int,
  "is_owner" boolean NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("role") REFERENCES "roles" ("role_id");
ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");
ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");


