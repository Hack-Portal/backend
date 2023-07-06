CREATE TABLE "past_works" (
  "opus" serial PRIMARY KEY,
  "thumbnail_image" bytea NOT NULL,
  "explanatory_text" text NOT NULL
);

CREATE TABLE "hackathons_data" (
  "opus" int NOT NULL,
  "award_id" int NOT NULL,
  "hackathon_id" int NOT NULL
);

CREATE TABLE "awards" (
  "award_id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "past_work_tags" (
  "opus" int NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "account_tags" (
  "account_id" int NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "tech_tags" (
  "tech_tag_id" serial PRIMARY KEY,
  "tech_tag" varchar NOT NULL
);

CREATE TABLE "accounts" (
  "account_id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "username" int NOT NULL,
  "icon" bytea,
  "explanatory_text" text,
  "locate_id" int NOT NULL,
  "rate" int NOT NULL,
  "show_locate" boolean NOT NULL,
  "show_rate" boolena NOT NULL,
  "update_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "hashed_password" varchar,
  "email" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rate_entries" (
  "account_id" int NOT NULL,
  "rate" int NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "locates" (
  "locate_id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "hackathons" (
  "hackathon_id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" bytea,
  "description" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_Date" date NOT NULL,
  "term" int NOT NULL
);

CREATE TABLE "hackathon_status_tags" (
  "hackathon_id" int NOT NULL,
  "status_id" int NOT NULL
);

CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);

CREATE TABLE "bookmarks" (
  "hackathon_id" int NOT NULL,
  "account_id" int NOT NULL
);

CREATE TABLE "account_past_works" (
  "opus" int NOT NULL,
  "account_id" int NOT NULL
);

CREATE TABLE "follows" (
  "to_account_id" int NOT NULL,
  "from_account_id" int NOT NULL
);

CREATE TABLE "rooms" (
  "room_id" uuid PRIMARY KEY,
  "hackathon_id" int NOT NULL,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "limit" int NOT NULL,
  "is_status" boolean NOT NULL
);

CREATE TABLE "rooms_tech_tags" (
  "room_id" uuid NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "rooms_accounts" (
  "room_id" uuid NOT NULL,
  "account_id" int NOT NULL
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("locate_id") REFERENCES "locates" ("locate_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "past_work_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "past_work_tags" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "account_past_works" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "account_past_works" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "rate_entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("award_id") REFERENCES "awards" ("award_id");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "rooms_tech_tags" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "rooms_tech_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");
