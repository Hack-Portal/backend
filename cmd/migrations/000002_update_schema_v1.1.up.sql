ALTER TABLE "accounts" DROP CONSTRAINT "accounts_user_id_fkey";
ALTER TABLE "users" DROP CONSTRAINT "users_pkey";
DROP TABLE IF EXISTS "users";
	
ALTER TABLE "accounts" DROP COLUMN "user_id";
ALTER TABLE "awards" ADD "icon" text NOT NULL;
ALTER TABLE "tech_tags" ADD "icon" text NOT NULL;
ALTER TABLE "frameworks" ADD "icon" text NOT NULL;
ALTER TABLE "accounts" ADD "email" varchar UNIQUE NOT NULL;