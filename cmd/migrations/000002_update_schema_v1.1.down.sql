CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "email" varchar,
  "hashed_password" text,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  "is_delete" boolean NOT NULL DEFAULT false
);
ALTER TABLE "accounts" ADD "user_id" varchar NOT NULL;
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "awards" DROP COLUMN "icon";
ALTER TABLE "tech_tags" DROP COLUMN "icon";
ALTER TABLE "accounts" DROP COLUMN "email";